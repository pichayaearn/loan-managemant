package repo

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/pichayaearn/loan-management/pkg/customer/model"
	"github.com/shopspring/decimal"
	"github.com/uptrace/bun"
)

type paymentBun struct {
	bun.BaseModel  `bun:"table:customers.payment"`
	Id             int64 `bun:",pk,autoincrement,type:bigserial"`
	LoanId         int32
	MonthlyAmount  decimal.Decimal
	LoanAmount     decimal.Decimal
	InterestAmount decimal.Decimal
	LoanBalance    decimal.Decimal
	PayDate        time.Time
	NextPayDate    time.Time
	CreatedAt      time.Time
	CreatedBy      uuid.UUID
	UpdatedAt      time.Time
	UpdatedBy      uuid.UUID
	DeletedAt      time.Time
	DeletedBy      uuid.UUID
}
type PaymentRepo struct {
	db *bun.DB
}

func NewPaymentRepo(db *bun.DB) model.PaymentRepo {
	return &PaymentRepo{
		db: db,
	}
}

func (pr PaymentRepo) FindOnePaymentByLoanId(loainId int32, ctx context.Context) (*model.Payment, error) {
	payment := new(paymentBun)
	err := pr.db.NewSelect().Model(payment).Where("loan_id = ?", loainId).Limit(1).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.New("findPaymentByLoanIdisExists failed")
	}

	return payment.toModel()
}

func (pr PaymentRepo) Create(payment model.Payment) error {
	pb := toPaymentBun(payment)
	if _, err := pr.db.NewInsert().Model(&pb).Exec(context.Background()); err != nil {
		return errors.New("create payment failed")
	}
	return nil
}

func toPaymentBun(payment model.Payment) paymentBun {
	return paymentBun{
		Id:             payment.ID(),
		LoanId:         payment.LoanID(),
		MonthlyAmount:  payment.MonthlyAmount(),
		LoanAmount:     payment.LoanAmount(),
		InterestAmount: payment.InterestAmount(),
		LoanBalance:    payment.LoanBalance(),
		PayDate:        payment.PayDate(),
		NextPayDate:    payment.NextPayDate(),
		CreatedAt:      payment.CreatedAt(),
		CreatedBy:      payment.CreatedBy(),
		UpdatedAt:      payment.UpdatedAt(),
		UpdatedBy:      payment.UpdatedBy(),
		DeletedAt:      payment.DeletedAt(),
		DeletedBy:      payment.DeletedBy(),
	}
}

func (pb paymentBun) toModel() (*model.Payment, error) {
	return model.PaymentFactory(model.PaymentFactoryOpts{
		Id:             pb.Id,
		LoanId:         pb.LoanId,
		MonthlyAmount:  pb.MonthlyAmount,
		LoanAmount:     pb.LoanAmount,
		InterestAmount: pb.InterestAmount,
		LoanBalance:    pb.LoanBalance,
		PayDate:        pb.PayDate,
		NextPayDate:    pb.NextPayDate,
		CreatedA:       pb.CreatedAt,
		CreatedBy:      pb.CreatedBy,
		UpdatedAt:      pb.UpdatedAt,
		UpdatedBy:      pb.UpdatedBy,
		DeletedAt:      pb.DeletedAt,
		DeletedBy:      pb.DeletedBy,
	})
}
