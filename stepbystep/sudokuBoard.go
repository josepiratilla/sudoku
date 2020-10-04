package stepbystep

import (
	"bytes"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

//Cell contains all info and links of a cell
type Cell struct {
	//Value stores the actual value of the cell. 0 is unknown
	Value int
	//Candidates is a list of all possible values of the cell. When the cell is defined, candidates is null
	Candidates map[int]int
	//Houses references the three houses the Cell belongs to (row, column and box)
	Houses [][]*Cell
	//HousesPos indicates the inner position of the cell in the house
	HousePos []int
	//Row indicates the Row number the Cell belongs to
	Row int
	//Column indicates the Column number the Cell belongs to
	Column int
	//Index indicates the position in the linear organization of the cell
	Index int
}

//SudokuBoard is the base object of a sudoku
type SudokuBoard struct {
	//ListCells contains all cells in a linear slice.
	ListCells []Cell
	//FormatedCells points to the same values as List Cells, but ordering them in 2 dimensional slice
	FormatedCells [][]Cell
	//Houses contains all possible blocks defined in the sudoku board. They are ordered as Rows, Columns, Boxes.
	Houses [][]*Cell
	//Rows is the part of Houses that contains the Rows
	Rows [][]*Cell
	//Columns is the part of Houses that contains the Columns
	Columns [][]*Cell
	//Boxes is the part of Houses that contains the Boxes
	Boxes [][]*Cell
	//Size is the complete size of the sudoku board
	Size int
	//SmallSize is the inner size of a house.
	SmallSize int
}

//NewSudokuBoard creates a zeroed SudokuBoard.
func NewSudokuBoard(smallSize int) *SudokuBoard {

	s := new(SudokuBoard)
	s.SmallSize = smallSize
	s.Size = smallSize * smallSize
	sudokuLen := s.Size * s.Size

	//Create the main List
	s.ListCells = make([]Cell, sudokuLen)

	//Link the Formated Cells
	s.FormatedCells = make([][]Cell, s.Size)
	for i := 0; i < s.Size; i++ {
		s.FormatedCells[i] = s.ListCells[i*s.Size : (i+1)*s.Size]
	}

	//Link all Houses
	s.Houses = make([][]*Cell, s.Size*3)
	s.Rows = s.Houses[0:s.Size]
	s.Columns = s.Houses[s.Size : s.Size*2]
	s.Boxes = s.Houses[s.Size*2 : s.Size*3]

	for i := 0; i < s.Size; i++ {
		s.Rows[i] = make([]*Cell, s.Size)
		for j := range s.Rows[i] {
			s.Rows[i][j] = &s.FormatedCells[i][j]
		}
		s.Columns[i] = make([]*Cell, s.Size)
		for j := range s.Columns[i] {
			s.Columns[i][j] = &s.FormatedCells[j][i]
		}
		s.Boxes[i] = make([]*Cell, s.Size)
		for j := range s.Boxes[i] {
			s.Boxes[i][j] = &s.ListCells[boxPosToIndex(i, j, s.SmallSize)]
		}
	}

	// Set initial values for each cell
	for r := range s.FormatedCells {
		for c := range s.FormatedCells[r] {
			boxNum, boxIndex := posToBoxPos(r, c, s.SmallSize)

			s.FormatedCells[r][c].Houses = make([][]*Cell, 3)
			s.FormatedCells[r][c].HousePos = make([]int, 3)

			s.FormatedCells[r][c].Houses[0] = s.Rows[r][:]
			s.FormatedCells[r][c].HousePos[0] = c

			s.FormatedCells[r][c].Houses[1] = s.Columns[c][:]
			s.FormatedCells[r][c].HousePos[1] = r

			s.FormatedCells[r][c].Houses[2] = s.Boxes[boxNum][:]
			s.FormatedCells[r][c].HousePos[2] = boxIndex

			s.FormatedCells[r][c].Candidates = make(map[int]int)
			for i := 1; i <= s.Size; i++ {
				s.FormatedCells[r][c].Candidates[i] = i
			}

			s.FormatedCells[r][c].Index = r*s.Size + c
		}
	}

	return s
}

func boxPosToIndex(boxNum int, boxIndex int, smallSize int) int {
	size := smallSize * 2
	innerRPos := boxIndex / smallSize
	innerCPos := boxIndex % smallSize
	boxRPos := boxNum / smallSize
	boxCPos := boxNum % smallSize
	rPos := boxRPos*smallSize + innerRPos
	cPos := boxCPos*smallSize + innerCPos
	index := rPos*size + cPos
	return index
}

func posToBoxPos(rPos int, cPos int, smallSize int) (boxNum int, boxIndex int) {
	boxRPos := rPos / smallSize
	innnerRPos := rPos % smallSize
	boxCPos := cPos / smallSize
	innerCPos := cPos % smallSize
	boxNum = boxRPos*smallSize + boxCPos
	boxIndex = innnerRPos*smallSize + innerCPos
	return
}

//ToString creates an string with the contents of the cell, showing candidates if it's not defined
func (c *Cell) ToString() string {
	if c.Value == 0 {
		var b bytes.Buffer
		b.WriteString("(")
		a := make([]int, len(c.Candidates))
		first := true
		i := 0
		for k := range c.Candidates {
			a[i] = k
			i++
		}
		sort.Ints(a)
		for _, cand := range a {
			if first {
				first = false
			} else {
				b.WriteString(",")
			}
			b.WriteString(strconv.Itoa(cand))
		}
		b.WriteString(")")
		return b.String()
	}
	return strconv.Itoa(c.Value)
}

//ToString creates a graphical representation of the SudokuBoard
func (s *SudokuBoard) ToString() string {

	text := make([]string, s.Size*s.Size)
	maxText := 0

	for i := range text {
		text[i] = s.ListCells[i].ToString()
		if l := len(text[i]); l > maxText {
			maxText = l
		}
	}
	for i := range text {
		text[i] = centerText(text[i], maxText)
	}

	numHyphen := maxText*s.SmallSize + s.SmallSize + 1
	hLine := strings.Repeat("-", numHyphen) + "+"
	hLine = "+" + strings.Repeat(hLine, s.SmallSize) + "\n"

	var b bytes.Buffer
	b.WriteString(hLine)
	for i := 0; i < s.Size; i++ {
		b.WriteString("| ")
		for j := 0; j < s.Size; j++ {
			b.WriteString(text[i*s.Size+j])
			b.WriteString(" ")
			if j%s.SmallSize == s.SmallSize-1 {
				b.WriteString("| ")
			}
		}
		b.WriteString("\n")
		if i%s.SmallSize == s.SmallSize-1 {
			b.WriteString(hLine)
		}
	}

	return b.String()

}

//ToStringWithoutCandidates draws a simple version of the sudoku board without the candidates
func (s *SudokuBoard) ToStringWithoutCandidates() string {

	numCharacters := int(math.Floor(math.Log10(float64(s.Size)))) + 1
	numHyphen := numCharacters*s.SmallSize + s.SmallSize + 1
	hLine := strings.Repeat("-", numHyphen) + "+"
	hLine = "+" + strings.Repeat(hLine, s.SmallSize) + "\n"
	blank := strings.Repeat(".", numCharacters)

	var b bytes.Buffer
	b.WriteString(hLine)
	for i := 0; i < s.Size; i++ {
		b.WriteString("| ")
		for j := 0; j < s.Size; j++ {
			if s.FormatedCells[i][j].Value == 0 {
				b.WriteString(blank)
			} else {
				b.WriteString(fmt.Sprintf("%"+strconv.Itoa(numCharacters)+"v", strconv.Itoa(s.FormatedCells[i][j].Value)))
			}
			b.WriteString(" ")
			if j%s.SmallSize == s.SmallSize-1 {
				b.WriteString("| ")
			}
		}
		b.WriteString("\n")
		if i%s.SmallSize == s.SmallSize-1 {
			b.WriteString(hLine)
		}
	}

	return b.String()
}

func centerText(str string, max int) string {
	l := len(str)
	pad := max - l
	lPad := pad / 2
	rPad := pad - lPad
	return strings.Repeat(" ", lPad) + str + strings.Repeat(" ", rPad)
}

// func (s *SudokuBoard) Test() {
// 	//fmt.Println(s.FormatedCells[0][0])
// 	fmt.Println(formatCellToString(&s.FormatedCells[0][0]))
// 	s.FormatedCells[0][1].Set(1)
// 	fmt.Println(formatCellToString(&s.FormatedCells[0][1]))
// 	delete(s.FormatedCells[0][2].Candidates, 2)
// 	fmt.Println(formatCellToString(&s.FormatedCells[0][2]))
// }

//Set a value for a cell. Removes all it's candiates and removes the value from all houses the candidate belongs to.
func (c *Cell) Set(value int) {

	c.Value = value
	c.Candidates = nil

	for h := range c.Houses {
		for i := range c.Houses[h] {
			if i != c.HousePos[h] {
				delete(c.Houses[h][i].Candidates, value)
			}
		}
	}

}

//CreateSudokuBoardFromMatrix creates and initialices a SudokuBoard from a matrix of integers
func CreateSudokuBoardFromMatrix(matrix [][]int) *SudokuBoard {
	size := len(matrix)
	smallSize := int(math.Sqrt(float64(size)))
	s := NewSudokuBoard(smallSize)
	//fmt.Println(s.ToString())
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] != 0 {
				s.FormatedCells[i][j].Set(matrix[i][j])
				//fmt.Println(s.ToString())
			}
		}
	}
	return s
}
