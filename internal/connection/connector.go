package connection

import (
	"gorm.io/gorm"
)

type DBConnector interface {
	GetSqlDB() *gorm.DB
}

type dbConnector struct {
	db *gorm.DB
}

func NewDBConnector(db *gorm.DB) DBConnector {
	return &dbConnector{
		db: db,
	}
}

func (db *dbConnector) GetSqlDB() *gorm.DB {
	return db.db
}
