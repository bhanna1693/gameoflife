package gameoflife

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	gameoflifemodel "github.com/bhanna1693/gameoflife/internal/models/gameoflife"
	gameoflifeservice "github.com/bhanna1693/gameoflife/internal/services/gameoflife"
	"github.com/bhanna1693/gameoflife/internal/utils"
	gameoflifecomponents "github.com/bhanna1693/gameoflife/web/views/gameoflife"

	gameoflifeDatabase "github.com/bhanna1693/gameoflife/internal/database/gameoflife"
	"github.com/labstack/echo/v4"
)

const (
	defaultRows    = 10
	defaultColumns = 10
)

var (
	randSource = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func HandleGameOfLife(e echo.Context, db *sql.DB) error {
	var dto gameoflifemodel.GameOfLifeModel
	if err := e.Bind(&dto); err != nil {
		return fmt.Errorf("error binding GameOfLifeModel: %w", err)
	}

	setDefaultDimensions(&dto)
	log.Printf("Game of Life DTO: %+v\n", dto)

	dto.MatrixData = initializeMatrix(dto.Rows, dto.Columns)
	log.Printf("Matrix: %+v\n", dto.MatrixData)

	if err := gameoflifeDatabase.InsertGameOfLifeBoard(db, &dto); err != nil {
		return fmt.Errorf("error inserting game of life board: %w", err)
	}

	e.Response().Header().Set("HX-Replace-Url", "/gameoflife?rows="+strconv.Itoa(dto.Rows)+"&columns="+strconv.Itoa(dto.Columns))
	return utils.Render(e, gameoflifecomponents.GameOfLife(dto))
}

func HandleGameOfLifeBoard(c echo.Context, db *sql.DB) error {
	var dto gameoflifemodel.GameOfLifeModel

	err := c.Bind(&dto)
	if err != nil {
		return err
	}
	log.Printf("Game of Life Board DTO %v\n", dto)
	matrix, err := buildMatrixFromContext(c, dto.Rows, dto.Columns)
	if err != nil {
		return fmt.Errorf("error building matrix: %w", err)
	}

	dto.MatrixData = gameoflifeservice.GameOfLifeService{}.ProcessGameOfLife(cloneMatrix(matrix))
	dto.Cycles++

	if areMatricesEqual(matrix, dto.MatrixData) {
		log.Printf("Matrices are equal. Returning 286 status code.")
		dto.Finished = true
		err = dto.EvaluateMatrix()
		if err != nil {
			return err
		}

		err = gameoflifeDatabase.UpdateGameOfLifeBoard(db, &dto)
		if err != nil {
			return err
		}
		c.Response().WriteHeader(286)
		return utils.Render(c, gameoflifecomponents.GameOfLife(dto))
	}

	log.Printf("Processed Game of Life Data: %+v", dto)

	err = gameoflifeDatabase.UpdateGameOfLifeBoard(db, &dto)
	if err != nil {
		return err
	}

	return utils.Render(c, gameoflifecomponents.GameOfLife(dto))
}

func setDefaultDimensions(dto *gameoflifemodel.GameOfLifeModel) {
	if dto.Rows == 0 {
		dto.Rows = defaultRows
	}
	if dto.Columns == 0 {
		dto.Columns = defaultColumns
	}
}

func buildMatrixFromContext(c echo.Context, rows, columns int) ([][]int, error) {
	var matrix [][]int
	for row := 0; row < rows; row++ {
		var rowValues []int
		for col := 0; col < columns; col++ {
			key := fmt.Sprintf("matrix[%d][%d]", row, col)
			value, err := strconv.Atoi(c.FormValue(key))
			if err != nil {
				return nil, fmt.Errorf("invalid matrix value at [%d][%d]: %w", row, col, err)
			}
			rowValues = append(rowValues, value)
		}
		matrix = append(matrix, rowValues)
	}
	return matrix, nil
}

func initializeMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
		for j := range matrix[i] {
			matrix[i][j] = randSource.Intn(2)
		}
	}
	return matrix
}

func cloneMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	clonedMatrix := make([][]int, len(matrix))
	for i, row := range matrix {
		clonedMatrix[i] = make([]int, len(row))
		copy(clonedMatrix[i], row)
	}
	return clonedMatrix
}

func areMatricesEqual(matrix1, matrix2 [][]int) bool {
	// Check if matrices have the same dimensions
	if len(matrix1) != len(matrix2) || len(matrix1[0]) != len(matrix2[0]) {
		return false
	}

	// Iterate over each element and compare values
	for i := 0; i < len(matrix1); i++ {
		for j := 0; j < len(matrix1[0]); j++ {
			if matrix1[i][j] != matrix2[i][j] {
				return false
			}
		}
	}

	return true
}

func HandleGameOfLifeResults(e echo.Context, db *sql.DB) error {
	dtos, err := gameoflifeDatabase.SelectAllGameOfLifeResults(db)
	if err != nil {
		log.Println("Error getting all game of life results")
		return err
	}

	return utils.Render(e, gameoflifecomponents.GameOfLifeResults(dtos))
}

func HandleGameOfLifeDelete(e echo.Context, db *sql.DB) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		log.Println("Error converting id to int")
		return err
	}
	log.Printf("Deleting Game of Life Board with id: %v", id)
	err = gameoflifeDatabase.DeleteGameOfLifeBoard(db, id)
	if err != nil {
		log.Println("Error deleting game of life board")
		return err
	}
	return nil
}
