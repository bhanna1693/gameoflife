package gameoflife

import (
	"database/sql"
	"log"

	"github.com/bhanna1693/gameoflife/internal/models/gameoflife"
)

func CreateGameOfLifeTable(db *sql.DB) error {
	// Create a table
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS gameoflife (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			rows INTEGER NOT NULL DEFAULT 0,
			columns INTEGER NOT NULL DEFAULT 0,
			cycles INTEGER NOT NULL DEFAULT 0,
			matrix TEXT NOT NULL DEFAULT '',
			success BOOLEAN NOT NULL DEFAULT FALSE,
			finished BOOLEAN NOT NULL DEFAULT FALSE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	return err
}

func SelectGameOfLifeBoardById(db *sql.DB, id int) (*gameoflife.GameOfLifeModel, error) {
	log.Println("Getting Game of Life Board")
	var dto gameoflife.GameOfLifeModel
	row := db.QueryRow(`SELECT id, rows, columns, cycles, matrix, success, finished FROM gameoflife WHERE id = ?`, id)
	err := row.Scan(&dto.Id, &dto.Rows, &dto.Columns, &dto.Cycles, &dto.Matrix, &dto.Success, &dto.Finished)
	if err != nil {
		log.Println("Error scanning row")
		return nil, err
	}
	err = dto.UnmarshalMatrix()
	if err != nil {
		log.Println("Error unmarshalling matrix")
		return nil, err
	}
	return &dto, nil
}

func InsertGameOfLifeBoard(db *sql.DB, dto *gameoflife.GameOfLifeModel) error {
	str, err := dto.MarshalMatrix()
	if err != nil {
		log.Println("Error marshalling matrix")
		return err
	}
	dto.Matrix = string(str)
	log.Println("Inserting Game of Life Board")
	insertSQL := `INSERT INTO gameoflife (rows, columns, cycles, matrix, success, finished) VALUES (?, ?, ?, ?, ?, ?)`
	statement, err := db.Prepare(insertSQL)
	if err != nil {
		log.Println("Error preparing insert statement")
		return err
	}
	result, err := statement.Exec(dto.Rows, dto.Columns, dto.Cycles, dto.Matrix, dto.Success, dto.Finished)
	if err != nil {
		log.Println("Error executing insert statement")
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting last insert id")
		return err
	}
	dto.Id = int(id)
	return nil
}

func UpdateGameOfLifeBoard(db *sql.DB, dto *gameoflife.GameOfLifeModel) error {
	str, err := dto.MarshalMatrix()
	if err != nil {
		log.Println("Error marshalling matrix")
		return err
	}
	dto.Matrix = string(str)
	log.Println("Updating Game of Life Board")
	updateSQL := `UPDATE gameoflife SET rows = ?, columns = ?, cycles = ?, matrix = ?, success = ?, finished = ? WHERE id = ?`
	statement, err := db.Prepare(updateSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(dto.Rows, dto.Columns, dto.Cycles, dto.Matrix, dto.Success, dto.Finished, dto.Id)
	if err != nil {
		return err
	}
	return nil
}

func SelectAllGameOfLifeResults(db *sql.DB) ([]*gameoflife.GameOfLifeModel, error) {
	log.Println("Getting all Game of Life Results")
	var dtos []*gameoflife.GameOfLifeModel
	rows, err := db.Query(`SELECT id, rows, columns, cycles, matrix, success, finished FROM gameoflife`)
	if err != nil {
		log.Println("Error querying rows")
		return nil, err
	}
	for rows.Next() {
		var dto gameoflife.GameOfLifeModel
		err = rows.Scan(&dto.Id, &dto.Rows, &dto.Columns, &dto.Cycles, &dto.Matrix, &dto.Success, &dto.Finished)
		if err != nil {
			log.Println("Error scanning row")
			return nil, err
		}
		err = dto.UnmarshalMatrix()
		if err != nil {
			log.Println("Error unmarshalling matrix")
			return nil, err
		}
		dtos = append(dtos, &dto)
	}
	return dtos, nil
}

func DeleteGameOfLifeBoard(db *sql.DB, id int) error {
	log.Println("Deleting Game of Life Board")
	deleteSQL := `DELETE FROM gameoflife WHERE id = ?`
	statement, err := db.Prepare(deleteSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
