package main

import (
	"fmt"

	"github.com/josepiratilla/sudoku/models"
)

func main() {
	s := models.NewSudokuBoard()
	s.Cell[1][3] = 1
	s.Cell[2][4] = 1
	s.Cell[0][0] = 1
	fmt.Println(s.House(1))
	fmt.Print(s.ToString())
	fmt.Println(s.Check())

}
