package db

import (
	"context"
	"errors"
	entity "hw3/internal/domain/entity/account"
)

type InmemoryStorage struct {
	accounts map[int]entity.Account
}

func (storage *InmemoryStorage) Find(ctx context.Context, id int) (entity.Account, error) {
	select {
	case <-ctx.Done():
		return entity.Account{}, ctx.Err()
	default:
	}
	account, ok := storage.accounts[id]
	if !ok {
		return entity.Account{}, errors.New("user with this ID not found")
	}
	return account, nil
}
func (storage *InmemoryStorage) Update(ctx context.Context, account entity.Account) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	_, ok := storage.accounts[account.Id]
	if !ok {
		return errors.New("user with this ID not found")
	}
	storage.accounts[account.Id] = account
	return nil
}

func (storage *InmemoryStorage) CreateRepository() error {
	storage.accounts = make(map[int]entity.Account)
	return nil
}

func (storage *InmemoryStorage) Register(ctx context.Context) (int, error) {
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
	}
	newId := len(storage.accounts) + 1
	newAccount := entity.Account{
		Id:      newId,
		Balance: 0,
	}
	storage.accounts[newId] = newAccount
	return newId, nil
}
