package stepbystep

import (
	"fmt"
)

//stepOnly1Candidate looks for one cell containing only one candidate available.
func stepOnly1Candidate(s *SudokuBoard) (string, bool) {

	for i := range s.ListCells {
		if s.ListCells[i].Value == 0 {
			if len(s.ListCells[i].Candidates) == 1 {
				for k := range s.ListCells[i].Candidates {
					s.ListCells[i].Set(k)
					str := fmt.Sprintf("%s -> %d: Only one candidate available.\n", s.ListCells[i].TextPosition(), k)
					return str, true
				}
			}
		}
	}

	return "", false
}

func stepOnly1OptionInHouse(s *SudokuBoard) (string, bool) {

	for house := range s.Houses {
		for candidate := 1; candidate <= s.Size; candidate++ {
			count := 0
			lastPos := 0
			for i := range s.Houses[house] {
				if _, exists := s.Houses[house][i].Candidates[candidate]; exists {
					count++
					lastPos = i
				}
			}
			if count == 1 {
				s.Houses[house][lastPos].Set(candidate)

				str := fmt.Sprintf("%s -> %d: There was only one possible option at %s for candidate %d\n",
					s.Houses[house][lastPos].TextPosition(),
					candidate, houseTextDescription(s, house), candidate)
				return str, true
			}
		}
	}

	return "", false
}
