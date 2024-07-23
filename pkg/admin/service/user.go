package service

import (
	"context"
	"errors"

	"github.com/pichayaearn/loan-management/pkg/admin/model"
)

type UserService struct {
	userRepo model.UserRepo
}

type NewUserServiceCfgs struct {
	UserRepo model.UserRepo
}

func NewUserService(cfg NewUserServiceCfgs) model.UserService {
	return &UserService{
		userRepo: cfg.UserRepo,
	}
}

func (uService UserService) CreateUser(opts model.CreateUser) error {
	//check email exist
	userExist, err := uService.userRepo.Get(model.GetUserOpts{
		Email: opts.Email,
	}, context.Background())
	if err != nil {
		return err
	}

	if userExist != nil {
		//email already used
		return errors.New("email is already used")
	}

	//create user
	newUser, err := model.NewUser(opts.Email, opts.Password)
	if err != nil {
		return err
	}
	if err := uService.userRepo.Create(*newUser); err != nil {
		return err
	}

	return nil
}

func (uSvc UserService) GetUser(opts model.GetUserOpts, ctx context.Context) (*model.User, error) {
	return uSvc.userRepo.Get(opts, ctx)
}
