package models

import (
	"math/rand"
	"time"
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

// Find next will try to find the next possible number for the position defined in row/column.
// Seed is the number for the ring of 9 integers that marks the init. 
// If current is 0, it will start from the Seed and check all 9 numbers, until one is found.
// If current is not 0, it will start from this number, and stop when arriving to seed again.
// It will return the next valid number, or 0 if no number was found.
func FindNext(sudoku *SudokuBoard, row int, column int, seed int, current int) value int {
	return 0
}

