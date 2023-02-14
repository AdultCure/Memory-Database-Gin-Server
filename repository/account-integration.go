package repository

import "http-type-one/db"

type AccountIntegrationDB struct {
	db *db.MemoryDB
}

func NewAccountIntegration(db *db.MemoryDB) *AccountsDB {
	return &AccountsDB{db: db}
}
