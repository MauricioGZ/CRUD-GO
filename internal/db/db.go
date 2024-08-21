package db

import (
	"database/sql"
	"fmt"

	"github.com/MauricioGZ/CRUD-GO/settings"
	_ "github.com/go-sql-driver/mysql"
)

func New(s settings.Settings) (*sql.DB, error) {
	dbConnTCP := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True",
		s.DB.User,
		s.DB.Password,
		s.DB.Host,
		s.DB.Port,
		s.DB.Name)
	db, err := sql.Open("mysql", dbConnTCP)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(5)

	return db, nil
}
