package service

import (
	"http-type-one/models"
	"http-type-one/repository"
)

type Accounts interface {
	CreateAccounts(account models.Accounts) (int, error)
	GetAll() []models.Accounts
	GetAccount(id int) models.Accounts
	DeleteAccount(id int) (int, error)
	UpdateAccount(account models.Accounts) (int, error)
}

type AccountIntegration interface {
}

type Service struct {
	Accounts
	AccountIntegration
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Accounts:           NewAccountService(repos.Accounts),
		AccountIntegration: NewAccountIntegrationService(repos.AccountIntegration),
	}
}
