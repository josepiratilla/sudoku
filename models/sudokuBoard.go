package models

import (
	"bytes"
	"strconv"
)

//SudokuBoard represents a board of sudoku that can be solved or not. A 0 means that number has not been defined.
type SudokuBoard struct {
	Cell  [9][9]int8
	block [27][9]*int8
}

func NewSudokuBoard() *SudokuBoard {
	s := new(SudokuBoard)
	//Map the row pointers
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			s.block[i][j] = &s.Cell[i][j]
		}
	}
	//Map the column pointers
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			s.block[i+9][j] = &s.Cell[j][i]
		}
	}
	//Map the house pointers
	for hi := 0; hi < 3; hi++ {
		for hj := 0; hj < 3; hj++ {
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					s.block[hi*3+hj+18][i*3+j] = &s.Cell[hi*3+i][hj*3+j]
				}
			}
		}
	}
	return s
}

//Row returns the nth row (starting from 0)
func (s *SudokuBoard) Row(rowNum int) [9]int8 {
	return s.Cell[rowNum]
}

//Column returns the nth column (starting from 0)
func (s *SudokuBoard) Column(columnNum int) [9]int8 {
	var c [9]int8
	for a := 0; a < 9; a++ {
		c[a] = s.Cell[a][columnNum]
	}
	return c
}

//House returns the elements of the nth house. A house is a 3x3 cell block.
//House are numbered as folows:
//    012
//    345
//    678
func (s *SudokuBoard) House(houseNum int) [9]int8 {
	var h [9]int8

	for i := 0; i < 9; i++ {
		h[i] = *(s.block[houseNum+18][i])
	}

	return h
}

func (s *SudokuBoard) ToString() string {
	var b bytes.Buffer
	const hLine = "+-------+-------+-------+\n"
	b.WriteString(hLine)
	for i := 0; i < 9; i++ {
		b.WriteString("| ")
		for j := 0; j < 9; j++ {
			if s.Cell[i][j] == 0 {
				b.WriteString(".")
			} else {
				b.WriteString(strconv.Itoa(int(s.Cell[i][j])))
			}
			b.WriteString(" ")
			if j%3 == 2 {
				b.WriteString("| ")
			}
		}
		b.WriteString("\n")
		if i%3 == 2 {
			b.WriteString(hLine)
		}
	}
	return b.String()
}

func checkBlock(b *[9]*int8) bool {

	var count int
	for i := int8(1); i < 10; i++ {
		count = 0
		for j := 0; j < 9; j++ {
			if *b[j] == i {
				count++
				if count > 1 {
					return false
				}
			}
		}
	}
	return true
}

//Check if the board is in a valid status (returns true) or if there are repeated values in the same block (returns false)
func (s *SudokuBoard) Check() bool {
	for i := 0; i < 27; i++ {
		if !checkBlock(&s.block[i]) {
			return false
		}
	}
	return true
}
