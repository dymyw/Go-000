package data

import (
	"Week02/internal/biz"
	"database/sql"
)

var _ biz.AccountRepo = &accountRepo{}
//var _ biz.AccountRepo = new(accountRepo)
//var _ biz.AccountRepo = (*accountRepo)(nil)

// NewAccountRepo
func NewAccountRepo(db *sql.DB) biz.AccountRepo {
	return &accountRepo{db: db}
}

// AccountRepo
type accountRepo struct {
	db *sql.DB
}

// SaveAccount
func (ar *accountRepo) SaveAccount(a *biz.Account) {
	//
}
