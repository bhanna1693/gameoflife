package gameoflife

import "database/sql"

func CreateGameOfLifeTable(db *sql.DB) error {
	// Create a table
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS gameoflife (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			rows INTEGER NOT NULL DEFAULT 0,
			columns INTEGER NOT NULL DEFAULT 0,
			cycles INTEGER NOT NULL DEFAULT 0,
			matrix TEXT NOT NULL DEFAULT '',
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	return err
}
