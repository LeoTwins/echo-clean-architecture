package service

import "gorm.io/gorm"

type RepositoryFactory func(*gorm.DB) interface{}

type ITransactionManager interface {
	ExecuteTransaction(fn func() error) error
}

type transactionManager struct {
	db *gorm.DB
}

func NewTransactionManager(db *gorm.DB) ITransactionManager {
	return &transactionManager{db}
}

func (tm *transactionManager) ExecuteTransaction(fn func() error) error {
	return tm.db.Transaction(func(tx *gorm.DB) error {
		return fn()
	})
}
