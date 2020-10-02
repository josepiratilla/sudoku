package generic

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//SudokuBoard represents a NxN sudoku game, solved or not. 0 means that the number has not been defined.
type SudokuBoard struct {
	//Size indicates the dimensions and # of options of a Sudoku Board
	//it needs to be a perfect square
	Size int
	//This is the size of houses dimensions. Size = smallSize^2
	smallSize int
	//Cell stores the real value of each cell
	Cell []int
	//Formamted Cell
	FormatedCell [][]int
	//Represents the N rows, columns and houses
	block [][]*int
}

func NewSudokuBoard(size int) *SudokuBoard {
	s := new(SudokuBoard)
	s.Size = size
	s.initialize()
	return s
}

func (s *SudokuBoard) initialize() {

	// if f := math.Sqrt(float64(s.Size)); f != math.Floor(f) {
	// 	panic("Size needs to be a perfect square")
	// }
	s.smallSize = int(math.Sqrt(float64(s.Size)))
	n := s.Size * s.Size

	//Create the Cell storage
	s.Cell = make([]int, n)
	//Create the slices with the board formated
	s.FormatedCell = make([][]int, s.Size)
	for i := 0; i < s.Size; i++ {
		s.FormatedCell[i] = s.Cell[i*s.Size : (i+1)*s.Size]
	}

	s.block = make([][]*int, s.Size*3)
	for i := 0; i < s.Size*3; i++ {
		s.block[i] = make([]*int, s.Size)
	}

	//Map the row pointers
	for i := 0; i < s.Size; i++ {
		for j := 0; j < s.Size; j++ {
			s.block[i][j] = &s.FormatedCell[i][j]
		}
	}
	//Map the column pointers
	for i := 0; i < s.Size; i++ {
		for j := 0; j < s.Size; j++ {
			s.block[i+s.Size][j] = &s.FormatedCell[j][i]
		}
	}
	//Map the house pointers
	for hi := 0; hi < s.smallSize; hi++ {
		for hj := 0; hj < s.smallSize; hj++ {
			for i := 0; i < s.smallSize; i++ {
				for j := 0; j < s.smallSize; j++ {
					s.block[hi*s.smallSize+hj+s.Size*2][i*s.smallSize+j] = &s.FormatedCell[hi*s.smallSize+i][hj*s.smallSize+j]
				}
			}
		}
	}
}

func (s *SudokuBoard) ToString() string {

	numCharacters := int(math.Floor(math.Log10(float64(s.Size)))) + 1
	numHyphen := numCharacters*s.smallSize + s.smallSize + 1
	hLine := strings.Repeat("-", numHyphen) + "+"
	hLine = "+" + strings.Repeat(hLine, s.smallSize) + "\n"
	blank := strings.Repeat(".", numCharacters)

	var b bytes.Buffer
	b.WriteString(hLine)
	for i := 0; i < s.Size; i++ {
		b.WriteString("| ")
		for j := 0; j < s.Size; j++ {
			if s.FormatedCell[i][j] == 0 {
				b.WriteString(blank)
			} else {
				b.WriteString(fmt.Sprintf("%"+strconv.Itoa(numCharacters)+"v", strconv.Itoa(s.FormatedCell[i][j])))
			}
			b.WriteString(" ")
			if j%s.smallSize == s.smallSize-1 {
				b.WriteString("| ")
			}
		}
		b.WriteString("\n")
		if i%s.smallSize == s.smallSize-1 {
			b.WriteString(hLine)
		}
	}

	return b.String()
}

func checkBlock(b []*int) bool {

	size := len(b)
	var count int
	for i := 1; i < size+1; i++ {
		count = 0
		for j := range b {
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

func (s *SudokuBoard) Check() bool {
	for _, block := range s.block {
		if !checkBlock(block) {
			return false
		}
	}
	return true
}

//Duplicate creates an independent SudokuBoard with the same contents.
func (s *SudokuBoard) Duplicate() *SudokuBoard {

	return CreateSudokuBoard(s.FormatedCell)
}

func CreateSudokuBoard(formatedCell [][]int) *SudokuBoard {
	size := len(formatedCell)
	c := NewSudokuBoard(size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			c.FormatedCell[i][j] = formatedCell[i][j]
		}
	}
	return c
}
