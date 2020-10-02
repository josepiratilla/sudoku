package models

import (
	"errors"
)

func houseRelative(row int, column int) (house int, innerPos int) {
	house = (row/3)*3 + column/3
	innerPos = (row%3)*3 + column%3
	return
}

func guess(sudoku *SudokuBoard, row int, column int, value int) bool {
	house, innerPos := houseRelative(row, column)
	for i := 0; i < 9; i++ {
		//review the row
		if i != column && *sudoku.block[row][i] == value {
			return false
		}
		//review the column
		if i != row && *sudoku.block[column+9][i] == value {
			return false
		}
		//review the house
		if i != innerPos && *sudoku.block[house+18][i] == value {
			return false
		}
	}
	//review the column
	return true
}

// findNext will try to find the next possible number for the position defined in row/column.
// Seed is the number for the ring of 9 integers that marks the init.
// If initial is 0, it will start from the Seed and check all 9 numbers, until one is found.
// If initial is not 0, it will start from this number, and stop when arriving to seed again.
// It will return the next valid number, or 0 if no number was found.
func findNext(sudoku *SudokuBoard, row int, column int, seed int, initial int) int {
	current := initial
	for {
		if current == 0 {
			current = seed
		} else {
			current = current%9 + 1
			if current == seed {
				return 0
			}
		}
		if guess(sudoku, row, column, current) {
			return current
		}
	}

}

func createFixedMap(s *SudokuBoard) []bool {
	var fixedMap [81]bool
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fixedMap[i*9+j] = !(s.Cell[i][j] == 0)
		}
	}
	return fixedMap[:]
}

func solver(sudoku *SudokuBoard, fixedMap []bool, i int) (*SudokuBoard, error) {

	var row, column, next int

	for {
		//fmt.Println(sudoku.ToString())
		if i > 80 {
			return sudoku, nil
		}
		if fixedMap[i] {
			i++
			continue
		}
		row = i / 9
		column = i % 9
		next = findNext(sudoku, row, column, 1, sudoku.Cell[row][column])
		if next == 0 {
			//Going back process
			sudoku.Cell[row][column] = 0
			for {
				i--
				if i < 0 {
					return nil, errors.New("No solution found")
				}
				if !fixedMap[i] {
					break
				}
			}
		} else {
			sudoku.Cell[row][column] = next
			i++
		}
	}
}

func Solver(sudoku *SudokuBoard) (*SudokuBoard, error) {
	workingSudoku := sudoku.Duplicate()
	fixedMap := createFixedMap(sudoku)
	return solver(workingSudoku, fixedMap[:], 0)
}
