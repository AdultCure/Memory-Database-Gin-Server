package repository

import (
	"errors"
	"fmt"
	"http-type-one/db"
	"http-type-one/models"
)

type AccountsDB struct {
	db *db.MemoryDB
}

func NewAccounts(db *db.MemoryDB) *AccountsDB {
	return &AccountsDB{db: db}
}

func (r *AccountsDB) CreateAccounts(account models.Accounts) (int, error) {
	account.ID = r.db.Accounts[len(r.db.Accounts)-1].ID + 1

	r.db.Accounts = append(r.db.Accounts, account)
	fmt.Println(r.db.Accounts)

	return account.ID, nil
}

func (r *AccountsDB) GetAll() []models.Accounts {
	var result []models.Accounts

	for i := 0; i < len(r.db.Accounts); i++ {
		result = append(result, r.db.Accounts[i])
	}

	return result
}

func (r *AccountsDB) GetAccount(id int) models.Accounts {
	var result models.Accounts
	for i := 0; i < len(r.db.Accounts); i++ {
		if r.db.Accounts[i].ID == id {
			result = r.db.Accounts[i]
		}
	}
	return result
}

func (r *AccountsDB) DeleteAccount(id int) (int, error) {
	for i := 0; i < len(r.db.Accounts); i++ {
		if r.db.Accounts[i].ID == id {
			r.db.Accounts = append(r.db.Accounts[:i], r.db.Accounts[i+1])
		} else {
			return 0, errors.New("not found")
		}
	}

	return id, nil
}

func (r *AccountsDB) UpdateAccount(account models.Accounts) (int, error) {
	for i := 0; i < len(r.db.Accounts); i++ {
		if r.db.Accounts[i].ID == account.ID {
			if account.AccessToken != "" {
				r.db.Accounts[i].AccessToken = account.AccessToken
			} else if account.RefreshToken != "" {
				r.db.Accounts[i].RefreshToken = account.RefreshToken
			} else if account.Expires != 0 {
				r.db.Accounts[i].Expires = account.Expires
			}
		}
	}
	return account.ID, nil
}
