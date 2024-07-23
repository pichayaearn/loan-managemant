package repo

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/pichayaearn/loan-management/pkg/customer/model"
	"github.com/shopspring/decimal"
	"github.com/uptrace/bun"
)

type loanBun struct {
	bun.BaseModel `bun:"table:customers.loan"`
	Id            int32 `bun:",pk,autoincrement,type:serial"`
	CustomerID    uuid.UUID
	Amount        decimal.Decimal
	Interest      decimal.Decimal
	StartDate     time.Time
	EndDate       time.Time
	DebtPayDate   int
	CreatedAt     time.Time
	CreatedBy     uuid.UUID
	UpdatedAt     time.Time
	UpdatedBy     uuid.UUID
	DeletedAt     time.Time
	DeletedBy     uuid.UUID
}

type LoanRepo struct {
	db *bun.DB
}

func NewLoanRepo(db *bun.DB) model.LoanRepo {
	return &LoanRepo{
		db: db,
	}
}

func (lr LoanRepo) Create(loan model.Loan) error {
	lb := toLoanBun(loan)
	log.Printf("create loan")
	if _, err := lr.db.NewInsert().Model(&lb).Exec(context.Background()); err != nil {
		return errors.New("create loan failed")
	}
	return nil
}

func (lr LoanRepo) GetByID(id int32, ctx context.Context) (*model.Loan, error) {
	loanBun := new(loanBun)
	if err := lr.db.NewSelect().Model(loanBun).Where("id = ?", id).Scan(ctx); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("get loan by id not found")
		}
		return nil, errors.New("get loan by id")
	}

	return loanBun.toModel()

}

func toLoanBun(loan model.Loan) loanBun {
	return loanBun{
		Id:         loan.ID(),
		CustomerID: loan.CustomerID(),
		Amount:     loan.Amount(),
		Interest:   loan.Interest(),
		StartDate:  loan.StartDate(),
		EndDate:    loan.EndDate(),
		CreatedAt:  loan.CreatedAt(),
		CreatedBy:  loan.CreatedBy(),
		UpdatedAt:  loan.UpdatedAt(),
		UpdatedBy:  loan.UpdatedBy(),
		DeletedAt:  loan.DeletedAt(),
		DeletedBy:  loan.DeletedBy(),
	}
}

func (lb loanBun) toModel() (*model.Loan, error) {
	return model.LoanFactory(model.LoanFactoryOpts{
		Id:          lb.Id,
		CustomerID:  lb.CustomerID,
		Amount:      lb.Amount,
		Interest:    lb.Interest,
		StartDate:   lb.StartDate,
		EndDate:     lb.EndDate,
		DebtPayDate: lb.DebtPayDate,
		CreatedAt:   lb.CreatedAt,
		CreatedBy:   lb.CreatedBy,
		UpdatedAt:   lb.UpdatedAt,
		UpdatedBy:   lb.UpdatedBy,
		DeletedAt:   lb.DeletedAt,
		DeletedBy:   lb.DeletedBy,
	})
}
