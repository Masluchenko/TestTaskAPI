package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	SQL  *sql.DB
	GORM *gorm.DB
}

func NewDatabase() *Database {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		os.Getenv("USER"),
		os.Getenv("PASS"),
		os.Getenv("DBNAME"),
		os.Getenv("SSLMODE"),
	)

	// ОБЩЕЕ подключение
	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Проверка подключения
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// GORM использует уже существующий sql.DB
	gormDB, err := gorm.Open(
		postgres.New(postgres.Config{
			Conn: sqlDB,
		}),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatal(err)
	}

	return &Database{
		SQL:  sqlDB,
		GORM: gormDB,
	}
}
