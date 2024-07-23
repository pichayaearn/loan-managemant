package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/pichayaearn/loan-management/cmd/api/config"
	admin_repo "github.com/pichayaearn/loan-management/pkg/admin/repo"
	admin_service "github.com/pichayaearn/loan-management/pkg/admin/service"
	auth_service "github.com/pichayaearn/loan-management/pkg/auth/service"

	auth_route "github.com/pichayaearn/loan-management/pkg/auth/route"
	"github.com/pichayaearn/loan-management/pkg/customer/repo"
	"github.com/pichayaearn/loan-management/pkg/customer/route"

	"github.com/pichayaearn/loan-management/pkg/customer/service"

	"github.com/pichayaearn/loan-management/pkg/middleware"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun/extra/bundebug"
)

func newServer(cfg *config.Config) *echo.Echo {
	log.Printf("env Db %+v", cfg.DB)
	db := cfg.DB.MustNewDB()
	logger := logrus.New()
	logger.Info("new server")
	if cfg.Environment == "development" {
		db.AddQueryHook(bundebug.NewQueryHook())
		logger.SetFormatter(&logrus.JSONFormatter{})
	}
	e := echo.New()

	mw := middleware.Authenticate{
		Secret: cfg.SecretKey,
	}
	customerRepo := repo.NewCustomerRepo(db)
	adminRepo := admin_repo.NewUserRepo(db)
	loanRepo := repo.NewLoanRepo(db)
	paymentRepo := repo.NewPaymentRepo(db)

	customerService := service.NewCustomerService(service.NewCustomerServiceCfg{
		CustomerRepo: customerRepo,
	})
	adminService := admin_service.NewUserService(admin_service.NewUserServiceCfgs{
		UserRepo: adminRepo,
	})
	authService := auth_service.NewAuthService(auth_service.NewAuthServiceCfgs{
		UserService: adminService,
		SecretKey:   cfg.SecretKey,
	})
	loanService := service.NewLoanService(service.NewLoanServiceCfg{
		Repo: loanRepo,
	})
	paymentService := service.NewPaymentService(service.NewPaymentServiceCfgs{
		LoanRepo:    loanRepo,
		PaymentRepo: paymentRepo,
	})

	e.POST("/sign-up", auth_route.CreateUser(auth_route.CreateUserCfg{
		UserService: adminService,
	}))

	e.POST("/login", auth_route.Login(auth_route.LoginCfg{
		AuthService: authService,
	}))

	g := e.Group("/customer")
	g.POST("", route.CreateCustomer(route.CreateCustomerCfg{
		CustomerService: customerService,
	}), mw.Authenticate)
	g.POST("/loan", route.CreateLoan(route.CreateLoanCfg{
		LoanService: loanService,
	}), mw.Authenticate)
	g.POST("/payment", route.CreatePayment(route.CreatePaymentCfg{
		PaymentService: paymentService,
	}), mw.Authenticate)
	return e

}
