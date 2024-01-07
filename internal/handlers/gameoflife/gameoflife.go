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
	return utils.Render(e, gameoflifecomponents.GameOfLife())
}

func HandleGameOfLifeStart(e echo.Context) error {
	x := 10
	y := 10
	matrix := make([][]int, x)
	for i := range matrix {
		matrix[i] = make([]int, y)
		for j := range matrix[i] {
			matrix[i][j] = randomize()
		}
	}

	gameDTO := gameoflifemodel.GameOfLifeDTO{
		Matrix: matrix,
	}
	return utils.Render(e, gameoflifecomponents.GameOfLifeStart(gameDTO))
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
	numberOfRows := 10
	numberOfColumns := 10
	// Assuming a 3x3 matrix for illustration purposes
	var matrix [][]int

	for row := 0; row < numberOfRows; row++ {
		var rowValues []int
		for col := 0; col < numberOfColumns; col++ {
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

	var dto gameoflifemodel.GameOfLifeDTO

	dto.Matrix = gameoflifeservice.GameOfLifeService{}.ProcessGameOfLife(cloneMatrix(matrix))

	if areMatricesEqual(matrix, dto.Matrix) {
		log.Printf("Matrices are equal. Returning 286 status code.")
		c.Response().WriteHeader(286)
		return utils.Render(c, gameoflifecomponents.GameOfLifeStart(dto))
	}

	log.Printf("Processed Game of Life Data: %+v", dto)

	return utils.Render(c, gameoflifecomponents.GameOfLifeStart(dto))
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
