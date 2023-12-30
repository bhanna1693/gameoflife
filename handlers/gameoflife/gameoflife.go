package gameoflife

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	gameoflifecomponents "github.com/bhanna1693/gameoflife/components/gameoflife"
	gameoflifemodel "github.com/bhanna1693/gameoflife/models/gameoflife"
	gameoflifeservice "github.com/bhanna1693/gameoflife/services/gameoflife"
	"github.com/bhanna1693/gameoflife/utils"
	"github.com/labstack/echo/v4"
)

func HandleGameOfLife(e echo.Context) error {
	return utils.Render(e, gameoflifecomponents.GameOfLife())
}

func HandleGameOfLifeStart(e echo.Context) error {
	gameDTO := gameoflifemodel.GameOfLifeDTO{
		Matrix: [][]int{
			// your matrix initialization here
			{randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize()},
			{randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize()},
			{randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize()},
			{randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize()},
			{randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize()},
			{randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize()},
			{randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize()},
			{randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize()},
			{randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize()},
			{randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize(), randomize()},
		},
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

	dto.Matrix = gameoflifeservice.GameOfLifeService{}.ProcessGameOfLife(matrix)

	log.Printf("Processed Game of Life Data: %+v", dto)

	return utils.Render(c, gameoflifecomponents.GameOfLifeStart(dto))
}
