package usecase

import (
	"context"
	"hw3/internal/application/dto/transfer"
	repository "hw3/internal/domain/repository/account"
	"hw3/internal/domain/service/transfer_service"
)

type BalanceUsecase struct {
	AccountRepo repository.AccountRepository
}

func (usecase *BalanceUsecase) Deposit(ctx context.Context, userId int, amount float64) error {
	account, err := usecase.AccountRepo.Find(ctx, userId)
	if err != nil {
		return err
	}
	err = account.Deposit(ctx, amount)
	if err != nil {
		return err
	}
	usecase.AccountRepo.Update(ctx, account)
	return nil
}

func (usecase *BalanceUsecase) Withdraw(ctx context.Context, userId int, amount float64) error {
	account, err := usecase.AccountRepo.Find(ctx, userId)
	if err != nil {
		return err
	}
	err = account.Withdraw(ctx, amount)
	if err != nil {
		return err
	}
	usecase.AccountRepo.Update(ctx, account)
	return nil
}

func (usecase *BalanceUsecase) Transfer(ctx context.Context, transferRequest transfer.TransferRequestDTO) error {
	fromAccount, err := usecase.AccountRepo.Find(ctx, transferRequest.FromId)
	if err != nil {
		return err
	}
	toAccount, err := usecase.AccountRepo.Find(ctx, transferRequest.ToId)
	if err != nil {
		return err
	}
	transferService := transfer_service.TransferService{}
	transferService.PerformTransfer(ctx, &fromAccount, &toAccount, transferRequest.Amount)
	usecase.AccountRepo.Update(ctx, fromAccount)
	usecase.AccountRepo.Update(ctx, toAccount)
	return nil
}

func (usecase *BalanceUsecase) GetBalance(ctx context.Context, userId int) (float64, error) {
	account, err := usecase.AccountRepo.Find(ctx, userId)
	if err != nil {
		return 0, err
	}
	return account.GetBalance(ctx), nil
}
