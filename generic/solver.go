package generic

import (
	"errors"
)

func houseRelative(row int, column int, smallSize int) (house int, innerPos int) {
	house = (row/smallSize)*smallSize + column/smallSize
	innerPos = (row%smallSize)*smallSize + column%smallSize
	return
}

func guess(sudoku *SudokuBoard, row int, column int, value int) bool {
	house, innerPos := houseRelative(row, column, sudoku.smallSize)
	for i := 0; i < sudoku.Size; i++ {
		//review the row
		if i != column && *sudoku.block[row][i] == value {
			return false
		}
		//review the column
		if i != row && *sudoku.block[column+sudoku.Size][i] == value {
			return false
		}
		//review the house
		if i != innerPos && *sudoku.block[house+sudoku.Size*2][i] == value {
			return false
		}
	}
	//review the column
	return true
}

// findNext will try to find the next possible number for the position defined in row/column.
// Seed is the number for the ring of all the integers that marks the init.
// If initial is 0, it will start from the Seed and check all numbers, until one is found.
// If initial is not 0, it will start from this number, and stop when arriving to seed again.
// It will return the next valid number, or 0 if no number was found.
func findNext(sudoku *SudokuBoard, row int, column int, seed int, initial int) int {
	current := initial
	for {
		if current == 0 {
			current = seed
		} else {
			current = current%sudoku.Size + 1
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
	fixedMap := make([]bool, s.Size*s.Size)
	for i := range s.Cell {
		fixedMap[i] = !(s.Cell[i] == 0)
	}
	return fixedMap
}

func solver(sudoku *SudokuBoard, fixedMap []bool, i int) (*SudokuBoard, error) {

	var row, column, next int
	// a := 0
	for {
		// a++
		// if a == 1000000 {
		// 	fmt.Print(sudoku.ToString())
		// 	a = 0
		// }
		if i >= sudoku.Size*sudoku.Size {
			return sudoku, nil
		}
		if fixedMap[i] {
			i++
			continue
		}
		row = i / sudoku.Size
		column = i % sudoku.Size
		next = findNext(sudoku, row, column, 1, sudoku.Cell[i])
		if next == 0 {
			//Going back process
			sudoku.Cell[i] = 0
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
			sudoku.Cell[i] = next
			i++
		}
	}
}

func Solver(sudoku *SudokuBoard) (*SudokuBoard, error) {
	workingSudoku := sudoku.Duplicate()
	//a := workingSudoku.Duplicate()
	//solverStepOnlyValue(a)
	//fmt.Print(a.ToString())
	solverStepOnlyValue(workingSudoku)
	fixedMap := createFixedMap(workingSudoku)
	return solver(workingSudoku, fixedMap[:], 0)
}

func solverStepOnlyValue(s *SudokuBoard) {
	//var previousFoundCandidate, foundCandidate, row, column, house int
	var candidates []int
	found := true
	for found {
		found = false
		for cell := range s.Cell {
			if s.Cell[cell] == 0 {
				// row = cell / s.Size
				// column = cell % s.Size
				// house, _ = houseRelative(row, column, s.smallSize)
				// column += s.Size
				// house += s.Size + s.Size
				// previousFoundCandidate = 0
				// for candidate := 1; candidate < s.Size; candidate++ {
				// 	foundCandidate = 0
				// 	for i := 1; i < s.Size; i++ {
				// 		if *s.block[row][i] == candidate || *s.block[column][i] == candidate || *s.block[house][i] == candidate {
				// 			foundCandidate = candidate
				// 			break
				// 		}
				// 	}
				// 	if foundCandidate != 0 {
				// 		if previousFoundCandidate != 0 {
				// 			break
				// 		} else {
				// 			previousFoundCandidate = foundCandidate
				// 		}
				// 	}
				// }
				// if previousFoundCandidate == foundCandidate && foundCandidate != 0 {
				// 	s.Cell[cell] = foundCandidate
				// 	found = true
				// }
				candidates = possibleCandidates(s, cell)
				if len(candidates) == 1 {
					s.Cell[cell] = candidates[0]
					found = true
				}
			}
		}
	}
}

func possibleCandidates(s *SudokuBoard, cell int) []int {
	var candidates []int
	var blocks []int
	found := false

	if s.Cell[cell] != 0 {
		candidates = append(candidates, s.Cell[cell])
	} else {
		blocks = blocksPerCell(cell, s.smallSize)
		for c := 1; c < s.Size+1; c++ {
			found = false
			for _, b := range blocks {
				for _, v := range s.block[b] {
					if *v == c {
						found = true
						break
					}
				}
				if found {
					break
				}
			}
			if !found {
				candidates = append(candidates, c)
			}
		}
	}
	return candidates
}

func blocksPerCell(cell int, smallSize int) []int {
	size := smallSize * smallSize
	row := cell / size
	column := cell % size
	house, _ := houseRelative(row, column, smallSize)
	a := [3]int{row, column + size, house + size + size}
	return a[:]
}
