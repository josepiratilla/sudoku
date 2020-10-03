package main

import (
	"fmt"
	"time"

	"github.com/josepiratilla/sudoku/generic"
	"github.com/josepiratilla/sudoku/stepbystep"
)

var sudoku16 = [][]int{
	[]int{0, 6, 0, 0, 14, 0, 0, 0, 0, 0, 2, 0, 15, 0, 0, 0},
	[]int{0, 0, 11, 0, 0, 0, 12, 3, 0, 0, 1, 0, 8, 0, 16, 0},
	[]int{0, 0, 0, 0, 0, 5, 0, 9, 0, 0, 0, 0, 0, 2, 0, 10},
	[]int{15, 3, 12, 0, 7, 11, 0, 16, 0, 8, 0, 5, 0, 0, 6, 0},
	[]int{6, 11, 0, 0, 8, 10, 0, 0, 0, 0, 5, 14, 0, 3, 2, 0},
	[]int{9, 10, 0, 0, 0, 0, 3, 0, 0, 15, 7, 0, 0, 14, 5, 11},
	[]int{14, 16, 0, 0, 1, 0, 0, 7, 0, 0, 10, 4, 0, 0, 0, 12},
	[]int{0, 0, 5, 0, 0, 2, 0, 0, 12, 0, 0, 8, 0, 10, 0, 13},
	[]int{11, 9, 10, 0, 0, 8, 0, 14, 0, 0, 0, 0, 0, 0, 0, 0},
	[]int{0, 0, 16, 8, 10, 0, 0, 0, 6, 0, 0, 2, 0, 0, 3, 1},
	[]int{5, 2, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 12, 14, 0},
	[]int{7, 15, 0, 0, 0, 0, 6, 0, 3, 0, 0, 0, 0, 0, 0, 0},
	[]int{0, 0, 0, 13, 16, 7, 0, 11, 9, 12, 0, 10, 4, 0, 0, 0},
	[]int{16, 0, 2, 0, 0, 14, 1, 0, 0, 0, 11, 13, 0, 0, 0, 0},
	[]int{8, 0, 0, 11, 5, 0, 4, 12, 0, 0, 0, 1, 13, 0, 0, 0},
	[]int{0, 5, 9, 1, 0, 3, 15, 13, 0, 0, 0, 16, 0, 0, 12, 14},
}

var sudoku9 = [][]int{
	[]int{8, 0, 0, 0, 0, 0, 0, 0, 0},
	[]int{0, 0, 3, 6, 0, 0, 0, 0, 0},
	[]int{0, 7, 0, 0, 9, 0, 2, 0, 0},
	[]int{0, 5, 0, 0, 0, 7, 0, 0, 0},
	[]int{0, 0, 0, 0, 4, 5, 7, 0, 0},
	[]int{0, 0, 0, 1, 0, 0, 0, 3, 0},
	[]int{0, 0, 1, 0, 0, 0, 0, 6, 8},
	[]int{0, 0, 8, 5, 0, 0, 0, 1, 0},
	[]int{0, 9, 0, 0, 0, 0, 4, 0, 0},
}

var sudoku9easy = [][]int{
	[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	[]int{0, 3, 0, 0, 0, 0, 1, 6, 0},
	[]int{0, 6, 7, 0, 3, 5, 0, 0, 4},
	[]int{6, 0, 8, 1, 2, 0, 9, 0, 0},
	[]int{0, 9, 0, 0, 8, 0, 0, 3, 0},
	[]int{0, 0, 2, 0, 7, 9, 8, 0, 6},
	[]int{8, 0, 0, 6, 9, 0, 3, 5, 0},
	[]int{0, 2, 6, 0, 0, 0, 0, 9, 0},
	[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
}

func main() {
	// cell := [9][9]int{
	// 	{8, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 3, 6, 0, 0, 0, 0, 0},
	// 	{0, 7, 0, 0, 9, 0, 2, 0, 0},
	// 	{0, 5, 0, 0, 0, 7, 0, 0, 0},
	// 	{0, 0, 0, 0, 4, 5, 7, 0, 0},
	// 	{0, 0, 0, 1, 0, 0, 0, 3, 0},
	// 	{0, 0, 1, 0, 0, 0, 0, 6, 8},
	// 	{0, 0, 8, 5, 0, 0, 0, 1, 0},
	// 	{0, 9, 0, 0, 0, 0, 4, 0, 0},
	// }
	// s := models.CreateSudokuBoard(&cell)

	// fmt.Println("Sudoku to solve:")
	// fmt.Print(s.ToString())
	// p, err := models.Solver(s)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Solution:")
	// 	fmt.Print(p.ToString())
	// 	if p.Check() {
	// 		fmt.Println("Checked: Solution is valid!")
	// 	}
	// }

	s := stepbystep.CreateSudokuBoard(3)
	println(s)
}

func solveAndPrint(matrix [][]int) {
	startTime := time.Now()
	s := generic.CreateSudokuBoard(matrix)
	fmt.Print(s.ToString())
	p, err := generic.Solver(s)
	if err == nil {
		fmt.Println("Solution:")
		fmt.Print(p.ToString())
	} else {
		fmt.Println("There is no solution for the Sudoku")
	}
	elapsedTime := time.Since(startTime)
	fmt.Printf("Time for solution %s\n", elapsedTime.String())
}
