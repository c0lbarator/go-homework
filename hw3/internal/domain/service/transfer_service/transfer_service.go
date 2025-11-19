package transfer_service

import (
	"context"
	entity "hw3/internal/domain/entity/account"
)

type TransferService struct {
}

func (service *TransferService) PerformTransfer(ctx context.Context, fromAccount *entity.Account, toAccount *entity.Account, amount float64) error {
	err := fromAccount.Withdraw(ctx, amount)
	if err != nil {
		return err
	}
	err = toAccount.Deposit(ctx, amount)
	if err != nil {
		return err
	}
	return nil
}
