package postgres

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/delanri/commonutil/logs"
	"github.com/delanri/commonutil/persistent"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func New(uri string, option *persistent.Option, logger logs.Logger) (persistent.ORM, error) {
	db, err := gorm.Open("postgres", uri)

	if err != nil {
		return nil, errors.Wrap(err, "failed to open postgres connection!")
	}

	db.SetLogger(logger)
	db.LogMode(option.LogMode)

	db.DB().SetMaxIdleConns(option.MaxIdleConnection)
	db.DB().SetMaxOpenConns(option.MaxOpenConnection)
	db.DB().SetConnMaxLifetime(option.ConnMaxLifetime)

	return &persistent.Impl{Database: db, Logger: logger}, nil
}
