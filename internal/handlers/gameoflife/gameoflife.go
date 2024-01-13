package gameoflife

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	gameoflifemodel "github.com/bhanna1693/gameoflife/internal/models/gameoflife"
	gameoflifeservice "github.com/bhanna1693/gameoflife/internal/services/gameoflife"
	"github.com/bhanna1693/gameoflife/internal/utils"
	gameoflifecomponents "github.com/bhanna1693/gameoflife/web/templates/gameoflife"
	"github.com/labstack/echo/v4"
)

func HandleGameOfLife(e echo.Context) error {
	var dto gameoflifemodel.GameOfLifeDTO

	err := e.Bind(&dto)
	if err != nil {
		return err
	}
	log.Printf("Game of Life DTO: %+v\n", dto)
	if dto.Rows == 0 {
		dto.Rows = 10
	}
	if dto.Columns == 0 {
		dto.Columns = 10
	}

	for row := 0; row < dto.Rows; row++ {
		var rowValues []int
		for col := 0; col < dto.Columns; col++ {
			rowValues = append(rowValues, randomize())
		}
		dto.Matrix = append(dto.Matrix, rowValues)
	}
	log.Printf("Matrix: %+v\n", dto.Matrix)

	e.Response().Header().Set("HX-Replace-Url", "/gameoflife?rows="+strconv.Itoa(dto.Rows)+"&columns="+strconv.Itoa(dto.Columns))
	return utils.Render(e, gameoflifecomponents.GameOfLife(dto))
}

func randomize() int {
	// Create a new source with a specific seed
	src := rand.NewSource(time.Now().UnixNano())

	// Use the source to create a new local random generator
	r := rand.New(src)

	// uss the random generator to create a random number
	return r.Intn(2)
}

func HandleGameOfLifeBoard(c echo.Context) error {
	var dto gameoflifemodel.GameOfLifeDTO

	err := c.Bind(&dto)
	if err != nil {
		return err
	}
	var matrix [][]int

	for row := 0; row < dto.Rows; row++ {
		var rowValues []int
		for col := 0; col < dto.Columns; col++ {
			key := fmt.Sprintf("matrix[%d][%d]", row, col)
			value := c.FormValue(key)
			num, err := strconv.Atoi(value)
			if err != nil {
				return err
			}
			rowValues = append(rowValues, num)
		}
		matrix = append(matrix, rowValues)
	}
	log.Printf("Matrix: %+v\n", matrix)

	dto.Matrix = gameoflifeservice.GameOfLifeService{}.ProcessGameOfLife(cloneMatrix(matrix))

	if areMatricesEqual(matrix, dto.Matrix) {
		log.Printf("Matrices are equal. Returning 286 status code.")
		dto.Finished = true
		dto.Success = isSuccess(dto.Matrix)
		c.Response().WriteHeader(286)
		return utils.Render(c, gameoflifecomponents.GameOfLife(dto))
	}

	log.Printf("Processed Game of Life Data: %+v", dto)

	return utils.Render(c, gameoflifecomponents.GameOfLife(dto))
}

func cloneMatrix(matrix [][]int) [][]int {
	// Ensure the original matrix is not nil or empty
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	// Get the dimensions of the matrix
	rows := len(matrix)
	cols := len(matrix[0])

	// Create a new matrix with the same dimensions
	clonedMatrix := make([][]int, rows)
	for i := range clonedMatrix {
		clonedMatrix[i] = make([]int, cols)
		copy(clonedMatrix[i], matrix[i])
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

func isSuccess(matrix [][]int) bool {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] == 0 {
				return false
			}
		}
	}
	return true
}
