package db

import (
	"demo/api/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(conf *configs.Config) *Db {
	db, err := gorm.Open(postgres.Open(conf.Db.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &Db{db}
}

// import (
// 	"TestTaskAPI/configs"
// 	"context"

// 	"github.com/pressly/goose/v3"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// func NewDb(conf *configs.Config) *Db {
// 	db, err := gorm.Open(postgres.Open(conf.Db.Dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}

// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		panic(err)
// 	}

// 	goose.SetBaseFS(nil)
// 	if err := goose.RunContext(context.Background(), "up", sqlDB, "migrations"); err != nil {
// 		panic(err)
// 	}

// 	return &Db{
// 		SQL:  sqlDB,
// 		GORM: db,
// 	}
// }
