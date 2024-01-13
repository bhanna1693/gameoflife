package gameoflife

type GameOfLifeDTO struct {
	Success  bool `form:"success"`
	Finished bool `form:"finished"`
	Rows     int  `form:"rows" query:"rows"`
	Columns  int  `form:"columns" query:"columns"`
	Cycles   int
	Matrix   [][]int `form:"matrix"`
}
