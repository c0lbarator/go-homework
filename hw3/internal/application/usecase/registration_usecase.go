package usecase

import (
	"context"
	repository "hw3/internal/domain/repository/account"
)

type RegisterUsecase struct {
	AccountRepo repository.AccountRepository
}

func (usecase *RegisterUsecase) Register(ctx context.Context) (int, error) {
	userId, err := usecase.AccountRepo.Register(ctx)
	if err != nil {
		return 0, err
	}
	return userId, nil
}
