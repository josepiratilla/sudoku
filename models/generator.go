package models

func randomInvalidSudoku() *SudokuBoard {
	rand.Seed(time.Now().UTC().UnixNano())
	s := NewSudokuBoard()
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			s.Cell[i][j] = rand.Intn(9) + 1
		}
	}
	return s
}
