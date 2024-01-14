package gameoflife

import (
	"encoding/json"
	"fmt"
	"log"
)

type GameOfLifeModel struct {
	Id         int  `param:"id"`
	Success    bool `form:"success"`
	Finished   bool `form:"finished"`
	Cycles     int  `form:"cycles"`
	Rows       int  `form:"rows"`
	Columns    int  `form:"columns"`
	Matrix     string
	MatrixData [][]int
}

func (dto *GameOfLifeModel) EvaluateMatrix() error {
	log.Println("Evaluating matrix")

	// Check if MatrixData is initialized and the size matches Rows and Columns
	if dto.MatrixData == nil || len(dto.MatrixData) != dto.Rows || len(dto.MatrixData[0]) != dto.Columns {
		return fmt.Errorf("matrix data is not properly initialized or does not match the specified dimensions")
	}

	for row := 0; row < dto.Rows; row++ {
		for col := 0; col < dto.Columns; col++ {
			if dto.MatrixData[row][col] == 1 {
				dto.Success = true
				return nil
			}
		}
	}

	dto.Success = false
	return nil
}

// transforms string representation of matrix into memory representation
func (dto *GameOfLifeModel) UnmarshalMatrix() error {
	return json.Unmarshal([]byte(dto.Matrix), &dto.MatrixData)
}

// transforms memory representation of matrix into string representation
func (dto *GameOfLifeModel) MarshalMatrix() ([]byte, error) {
	return json.Marshal(dto.MatrixData)
}
