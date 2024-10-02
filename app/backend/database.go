package backend

import "gorm.io/gorm"

type DBWrapper interface {
	DB() *gorm.DB
	Transaction(func(dbw DBWrapper) error) error
}

type dbWrapperImpl struct {
	db *gorm.DB
}

func (dbw *dbWrapperImpl) DB() *gorm.DB {
	return dbw.db
}

func (dbw *dbWrapperImpl) Transaction(fn func(txw DBWrapper) error) error {
	return dbw.db.Transaction(func(tx *gorm.DB) error {
		return fn(NewDBWrapper(tx))
	})
}

func NewDBWrapper(db *gorm.DB) DBWrapper {
	return &dbWrapperImpl{
		db: db,
	}
}
