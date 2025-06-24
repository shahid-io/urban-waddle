package config

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

/**
DATABASE CONFIGURATION
*/

type MySQLStore struct {
	DB *sql.DB
}

func NewMySQLStore(cfg mysql.Config) (*MySQLStore, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &MySQLStore{DB: db}, nil
}
