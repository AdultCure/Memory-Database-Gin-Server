package db

import (
	"http-type-one/models"
)

type MemoryDB struct {
	Accounts           []models.Accounts
	AccountIntegration []models.AccountIntegration
}

func NewDB() *MemoryDB {
	return &MemoryDB{
		Accounts:           []models.Accounts{{1, "123", "456", 35}, {2, "123456", "456789678", 151}},
		AccountIntegration: []models.AccountIntegration{},
	}
}
