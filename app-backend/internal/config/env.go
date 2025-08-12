package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ApiBaseUrl string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

func Load(logger *log.Logger) *Config {
	if err := godotenv.Load(); err != nil {
		logger.Println("[INFO] No .env file found, using environment variables")
	}

	apiBaseUrl, exists := os.LookupEnv("API_BASE_URL")
	if !exists {
		logger.Println("[ERROR] API_BASE_URL not set in environment")
		os.Exit(1)
	}

	dbUser, exists := os.LookupEnv("DB_USER")
	if !exists {
		logger.Println("[ERROR] DB_USER not set in environment")
		os.Exit(1)
	}

	dbPassword, exists := os.LookupEnv("DB_PASSWORD")
	if !exists {
		logger.Println("[ERROR] DB_PASSWORD not set in environment")
		os.Exit(1)
	}

	dbHost, exists := os.LookupEnv("DB_HOST")
	if !exists {
		logger.Println("[ERROR] DB_HOST not set in environment")
		os.Exit(1)
	}

	dbPort, exists := os.LookupEnv("DB_PORT")
	if !exists {
		logger.Println("[ERROR] DB_PORT not set in environment")
		os.Exit(1)
	}

	dbName, exists := os.LookupEnv("DB_NAME")
	if !exists {
		logger.Println("[ERROR] DB_NAME not set in environment")
		os.Exit(1)
	}

	return &Config{
		ApiBaseUrl: apiBaseUrl,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBHost:     dbHost,
		DBPort:     dbPort,
		DBName:     dbName,
	}
}

func (c *Config) DatabaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
}
