package main

import (
	"fmt"

	"github.com/josepiratilla/sudoku/models"
)

func main() {
	cell := [9][9]int{
		{8, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 3, 6, 0, 0, 0, 0, 0},
		{0, 7, 0, 0, 9, 0, 2, 0, 0},
		{0, 5, 0, 0, 0, 7, 0, 0, 0},
		{0, 0, 0, 0, 4, 5, 7, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 3, 0},
		{0, 0, 1, 0, 0, 0, 0, 6, 8},
		{0, 0, 8, 5, 0, 0, 0, 1, 0},
		{0, 9, 0, 0, 0, 0, 4, 0, 0},
	}
	s := models.CreateSudokuBoard(&cell)

	fmt.Println("Sudoku to solve:")
	fmt.Print(s.ToString())
	p, err := models.Solver(s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Solution:")
		fmt.Print(p.ToString())
		if p.Check() {
			fmt.Println("Checked: Solution is valid!")
		}
	}
}
