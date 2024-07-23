package repo

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/pichayaearn/loan-management/pkg/customer/model"
	"github.com/uptrace/bun"
)

type customerBun struct {
	bun.BaseModel `bun:"table:customers.customer"`
	CustomerID    uuid.UUID `bun:",pk,type:uuid,default:uuid_generate_v4()"`
	Firstname     string
	Lastname      string
	Mobile        string
	Email         string
	Status        string
	CreatedAt     time.Time
	CreatedBy     uuid.UUID
	UpdatedAt     time.Time
	UpdatedBy     uuid.UUID
	DeletedAt     time.Time
	DeletedBy     uuid.UUID
}

type CustomerRepo struct {
	db *bun.DB
}

func NewCustomerRepo(db *bun.DB) model.CustomerRepo {
	return &CustomerRepo{
		db: db,
	}
}

func (cr CustomerRepo) Create(customer model.Customer) error {
	cb := toCustomerBun(customer)
	if _, err := cr.db.NewInsert().Model(&cb).Exec(context.Background()); err != nil {
		return errors.New("create customer failed")
	}
	return nil
}

func (cr CustomerRepo) FindByCustomerId(customerID uuid.UUID, ctx context.Context) (*model.Customer, error) {
	cb := new(customerBun)
	if err := cr.db.NewSelect().Model(cb).Where("customer_id = ?", customerID).Scan(ctx); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("get customer by customer_id not found")
		}
		return nil, errors.New("get customer by customer_id ")
	}
	return cb.toModel()
}

func (cr CustomerRepo) Update(custmer model.Customer) error {
	cb := toCustomerBun(custmer)
	if _, err := cr.db.NewUpdate().
		Model(&cb).
		WherePK().
		Exec(context.Background()); err != nil {
		return err
	}
	return nil
}

func toCustomerBun(customer model.Customer) customerBun {
	return customerBun{
		CustomerID: customer.ID(),
		Firstname:  customer.Firstname(),
		Lastname:   customer.Lastname(),
		Mobile:     customer.Mobile(),
		Email:      customer.Email(),
		Status:     customer.Status(),
		CreatedAt:  customer.CreatedAt(),
		CreatedBy:  customer.CreatedBy(),
		UpdatedAt:  customer.UpdatedAt(),
		UpdatedBy:  customer.UpdatedBy(),
		DeletedAt:  customer.DeletedAt(),
		DeletedBy:  customer.DeletedBy(),
	}
}

func (cb customerBun) toModel() (*model.Customer, error) {
	return model.CustomerFactory(model.CustomerFactoryOpts{
		Id:        cb.CustomerID,
		Firtname:  cb.Firstname,
		Lastname:  cb.Lastname,
		Mobile:    cb.Mobile,
		Email:     cb.Email,
		Status:    cb.Status,
		CreatedAt: cb.CreatedAt,
		CreatedBy: cb.CreatedBy,
		UpdatedAt: cb.UpdatedAt,
		UpdatedBy: cb.UpdatedBy,
		DeletedAt: cb.DeletedAt,
		DeletedBy: cb.DeletedBy,
	})
}
