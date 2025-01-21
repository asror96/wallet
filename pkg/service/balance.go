package service

import (
	"wallet/pkg/repository"
)

type BalanceService struct {
	repo repository.Wallet
}

func NewBalanceService(repo repository.Wallet) *BalanceService {
	return &BalanceService{repo: repo}
}

func (s *BalanceService) Balance(uid string) (float64, error) {

	return s.repo.Balance(uid)
}
