package repository

import (
	"http-type-one/db"
	"http-type-one/models"
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

type Repository struct {
	Accounts
	AccountIntegration
}

func NewRepository(db *db.MemoryDB) *Repository {
	return &Repository{
		Accounts:           NewAccounts(db),
		AccountIntegration: NewAccountIntegration(db),
	}
}
