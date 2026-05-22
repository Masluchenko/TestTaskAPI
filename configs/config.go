package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	ServerPort string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return &Config{
		DBHost:     os.Getenv("HOST"),
		DBPort:     os.Getenv("PORT"),
		DBUser:     os.Getenv("USER"),
		DBPassword: os.Getenv("PASSWORD"),
		DBName:     os.Getenv("DBNAME"),
		DBSSLMode:  os.Getenv("SSLMODE"),
	}
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSSLMode)
}

// func getEnv(key, defaultValue string) string {
//     if value := os.Getenv(key); value != "" {
//         return value
//     }
//     return defaultValue
// }

// import (
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// )

// type Config struct {
// 	Db DbConfig
// }

// type DbConfig struct {
// 	Dsn string
// }

// func LoadConfig() *Config {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Println("Error loading .env file, using default config")
// 	}
// 	return &Config{
// 		Db: DbConfig{
// 			Dsn: os.Getenv("DSN"),
// 		},
// 	}
// }
