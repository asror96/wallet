package service

import (
	"fmt"
	"wallet"
	"wallet/pkg/repository"
)

type SendService struct {
	repo repository.Transaction
}

func NewSendService(repo repository.Transaction) *SendService {
	return &SendService{repo: repo}
}

func (s *SendService) Send(transaction wallet.Transaction) (int, error) {
	walletFrom, err := s.repo.GetWallet(transaction.From)

	if err != nil {
		return 0, err
	}

	if walletFrom.Balance < transaction.Sum {
		return 0, fmt.Errorf("insufficient balance")
	}

	newBalanceFrom := walletFrom.Balance - transaction.Sum
	err = s.repo.UpdateWalletBalance(transaction.From, newBalanceFrom)

	if err != nil {
		return 0, err
	}

	walletTo, err := s.repo.GetWallet(transaction.To)

	if err != nil {
		return 0, err
	}

	newBalanceTo := walletTo.Balance + transaction.Sum
	err = s.repo.UpdateWalletBalance(transaction.To, newBalanceTo)

	if err != nil {
		return 0, err
	}

	return s.repo.Send(transaction)
}
