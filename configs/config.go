package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigData struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	ServerPort string
}

type Config struct {
	Db DbConfig
}

type DbConfig struct {
	Dsn string
}

func LoadConfig() *ConfigData {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return &ConfigData{
		DBHost:     os.Getenv("HOST"),
		DBPort:     os.Getenv("PORT"),
		DBUser:     os.Getenv("USER"),
		DBPassword: os.Getenv("PASSWORD"),
		DBName:     os.Getenv("DBNAME"),
		DBSSLMode:  os.Getenv("SSLMODE"),
	}
}

func (c *ConfigData) GetDSN() *Config {
	transformDsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSSLMode)
	return &Config{
		Db: DbConfig{
			Dsn: transformDsn,
		},
	}
}
