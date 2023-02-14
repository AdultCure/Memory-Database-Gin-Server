package service

import (
	"http-type-one/models"
	"http-type-one/repository"
)

type AccountService struct {
	repo repository.Accounts
}

func NewAccountService(repo repository.Accounts) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) CreateAccounts(account models.Accounts) (int, error) {
	return s.repo.CreateAccounts(account)
}

func (s *AccountService) GetAll() []models.Accounts {
	return s.repo.GetAll()
}

func (s *AccountService) GetAccount(id int) models.Accounts {
	return s.repo.GetAccount(id)
}

func (s *AccountService) DeleteAccount(id int) (int, error) {
	return s.repo.DeleteAccount(id)
}

func (s *AccountService) UpdateAccount(account models.Accounts) (int, error) {
	return s.repo.UpdateAccount(account)
}
