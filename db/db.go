package db

import (
	"TestTaskAPI/configs"
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	SQL  *sql.DB
	GORM *gorm.DB
}

func NewDb(conf *configs.Config) *Database {
	db, err := gorm.Open(postgres.Open(conf.Db.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	goose.SetBaseFS(nil)
	if err := goose.RunContext(context.Background(), "up", sqlDB, "migrations"); err != nil {
		panic(err)
	}

	return &Database{
		SQL:  sqlDB,
		GORM: db,
	}
}
