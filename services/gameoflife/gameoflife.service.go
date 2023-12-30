package gameoflife

type GameOfLifeService struct {
}

// ProcessGameOfLife processes the game of life based on the provided gameoflife.GameOfLifeDTO.
// It implements the rules of the game of life and returns the updated gameoflife.GameOfLifeDTO.
func (s GameOfLifeService) ProcessGameOfLife(matrix [][]int) [][]int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			liveNeighbors := 0
			if i > 0 {
				if matrix[i-1][j] == 1 {
					liveNeighbors++
				}
			}
			if i < len(matrix)-1 {
				if matrix[i+1][j] == 1 {
					liveNeighbors++
				}
			}
			if j > 0 {
				if matrix[i][j-1] == 1 {
					liveNeighbors++
				}
			}
			if j < len(matrix[i])-1 {
				if matrix[i][j+1] == 1 {
					liveNeighbors++
				}
			}
			if i > 0 && j > 0 {
				if matrix[i-1][j-1] == 1 {
					liveNeighbors++
				}
			}
			if i > 0 && j < len(matrix[i])-1 {
				if matrix[i-1][j+1] == 1 {
					liveNeighbors++
				}
			}
			if i < len(matrix)-1 && j > 0 {
				if matrix[i+1][j-1] == 1 {
					liveNeighbors++
				}
			}
			if i < len(matrix)-1 && j < len(matrix[i])-1 {
				if matrix[i+1][j+1] == 1 {
					liveNeighbors++
				}
			}
			if matrix[i][j] == 1 {
				if liveNeighbors < 2 {
					matrix[i][j] = 0
				} else if liveNeighbors > 3 {
					matrix[i][j] = 0
				}
			} else {
				if liveNeighbors == 3 {
					matrix[i][j] = 1
				}
			}
		}
	}
	return matrix
}
