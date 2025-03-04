package repo

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/pichayaearn/loan-management/pkg/admin/model"
	"github.com/uptrace/bun"
)

type userBun struct {
	bun.BaseModel `bun:"table:admin.users"`
	UserID        uuid.UUID `bun:",pk,type:uuid,default:uuid_generate_v4()"`
	Email         string
	Password      string
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

type UserRepo struct {
	db *bun.DB
}

func NewUserRepo(db *bun.DB) model.UserRepo {
	return &UserRepo{db: db}
}

func (ur UserRepo) Get(opts model.GetUserOpts, ctx context.Context) (*model.User, error) {
	user := userBun{}
	q := ur.db.NewSelect().Model(&user)
	if err := q.ApplyQueryBuilder(addUserFilter(opts)).Limit(1).Scan(ctx); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.New("get user error")
	}
	return user.toUserModel()

}

func (ur UserRepo) Create(user model.User) error {
	ub := toUserBun(user)
	if _, err := ur.db.NewInsert().Model(&ub).Exec(context.Background()); err != nil {
		return errors.New("create user failed")
	}
	return nil
}

func addUserFilter(opts model.GetUserOpts) func(q bun.QueryBuilder) bun.QueryBuilder {
	return func(q bun.QueryBuilder) bun.QueryBuilder {
		if opts.UserID != uuid.Nil {
			q.Where("user_id = ?", opts.UserID)
		}
		if opts.Email != "" {
			q.Where("email = ?", opts.Email)
		}
		if opts.Status != "" {
			q.Where("status = ?", opts.Status)
		}
		return q

	}

}

func (ub userBun) toUserModel() (*model.User, error) {
	return model.UserFactory(model.UserFactoryOpts{
		UserID:    ub.UserID,
		Email:     ub.Email,
		Password:  ub.Password,
		Status:    ub.Status,
		CreatedAt: ub.CreatedAt,
		UpdatedAt: ub.UpdatedAt,
		DeletedAt: ub.DeletedAt,
	})
}

func toUserBun(user model.User) userBun {
	return userBun{
		UserID:    user.UserID(),
		Email:     user.Email(),
		Password:  user.Password(),
		Status:    string(user.Status()),
		CreatedAt: user.CreatedAt(),
		UpdatedAt: user.UpdatedAt(),
		DeletedAt: user.DeletedAt(),
	}
}
