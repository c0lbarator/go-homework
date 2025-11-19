package repository

import (
	"context"
	entity "hw3/internal/domain/entity/account"
)

type AccountRepository interface {
	Find(ctx context.Context, id int) (entity.Account, error)
	Update(ctx context.Context, account entity.Account) error
	CreateRepository() error
	Register(ctx context.Context) (int, error)
}
