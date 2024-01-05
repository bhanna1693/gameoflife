package database

import (
	"database/sql"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func CreateGameOfLifeTable(db *sql.DB) error {
	// Create a table
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS gameoflife (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			result TEXT NOT NULL
		)
	`)
	return err
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// LoadDBConfig loads the database configuration from environment variables.
func LoadDBConfig() *DBConfig {
	godotenv.Load(".env")
	return &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     parseInt(os.Getenv("DB_PORT")),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		// Handle error, e.g., log it or set a default value
		return 0
	}
	return i
}