package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/nanopoker/minisns/config"
	"github.com/pkg/errors"
)

var Db *gorm.DB

func Init() error {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_DATABASE))
	if err != nil {
		return errors.WithMessage(err, "Open connection to mysql failed")
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(10000)
	db.DB().SetConnMaxLifetime(7200)
	err = db.DB().Ping()
	if err != nil {
		err = errors.WithMessage(err, "Ping mysql failed")
		return err
	}
	Db = db
	return nil
}

// CloseDB closes current connections
func CloseDB() error {
	err := Db.Close()
	if err != nil {
		err = errors.WithMessage(err, "Close")
		return err
	}
	return nil
}
