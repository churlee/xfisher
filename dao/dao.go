package dao

import (
	"database/sql"
	"lilith/config"
)

type Dao struct {
	db *sql.DB
}

func NewDB() *Dao {
	dao := &Dao{
		db: config.GetDB(),
	}
	return dao
}
