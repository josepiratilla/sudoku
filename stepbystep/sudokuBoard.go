package stepbystep

//Cell contains all info and links of a cell
type Cell struct {
	//Value stores the actual value of the cell. 0 is unknown
	Value int
	//Candidates is a list of all possible values of the cell. When the cell is defined, candidates is null
	Candidates []int
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

	for i := 0; i < s.Size; {
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

	return s
}

func boxPosToIndex(boxNum int, boxIndex int, smallSize int) int {
	innerRPos := boxIndex / smallSize
	innerCPos := boxIndex % smallSize
	boxRPos := boxNum / smallSize
	boxCPos := boxNum % smallSize
	RPos := boxRPos*smallSize + innerRPos
	CPos := boxCPos*smallSize + innerCPos
	index := RPos*smallSize + CPos
	return index
}
