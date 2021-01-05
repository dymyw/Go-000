package data

import (
	"Week02/internal/biz"
	"database/sql"
)

// NewAccountRepo
func NewAccountRepo(db *sql.DB) biz.AccountRepo {
	return new(AccountRepo)
}

// AccountRepo
type AccountRepo struct {}

// SaveAccount
func (ar *AccountRepo) SaveAccount(a *biz.Account) {
	//
}
