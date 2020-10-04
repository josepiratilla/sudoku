package stepbystep

import (
	"bytes"
	"fmt"
)

//stepOnly1Candidate looks for one cell containing only one candidate available.
func stepNakedSingle(s *SudokuBoard) (string, bool) {

	for i := range s.ListCells {
		if s.ListCells[i].Value == 0 {
			if len(s.ListCells[i].Candidates) == 1 {
				for k := range s.ListCells[i].Candidates {
					s.ListCells[i].Set(k)
					str := fmt.Sprintf("Naked Single: %s -> %d\n", s.ListCells[i].TextPosition(), k)
					return str, true
				}
			}
		}
	}

	return "", false
}

func stepHiddenSingle(s *SudokuBoard) (string, bool) {

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

				str := fmt.Sprintf("Hidden Single at %s: %s -> %d\n",
					houseTextDescription(s, house), s.Houses[house][lastPos].TextPosition(),
					candidate)
				return str, true
			}
		}
	}

	return "", false
}

func stepLockedCandidates(s *SudokuBoard) (string, bool) {

	var intersection, remainingLine, remainingBox []*Cell
	var line, box int
	var str string
	var ok bool

	//Look for intersections Row-Box
	for i := 0; i < s.Size*s.Size; i += s.SmallSize {

		intersection = make([]*Cell, 0)
		for j := i; j < i+s.SmallSize; j++ {
			intersection = append(intersection, &s.ListCells[j])
		}
		remainingLine = make([]*Cell, 0)
		remainingBox = make([]*Cell, 0)
		line = s.ListCells[i].Row
		box = s.ListCells[i].Box
		for j := range s.Rows[line] {
			if s.Rows[line][j].Box != box {
				remainingLine = append(remainingLine, s.Rows[line][j])
			}
			if s.Boxes[box][j].Row != line {
				remainingBox = append(remainingBox, s.Boxes[box][j])
			}
		}
		str, ok = stepLockedCandidatesOneGroup(s, intersection, remainingLine, remainingBox)
		if ok {
			return str, ok
		}
	}

	//Look for intersections Column-Box
	for i := 0; i < s.Size*s.Size; {

		intersection = make([]*Cell, 0)
		for j := i; j < i+s.Size*s.SmallSize; j += s.Size {
			intersection = append(intersection, &s.ListCells[j])
		}
		remainingLine = make([]*Cell, 0)
		remainingBox = make([]*Cell, 0)
		line = s.ListCells[i].Column
		box = s.ListCells[i].Box
		for j := range s.Columns[line] {
			if s.Columns[line][j].Box != box {
				remainingLine = append(remainingLine, s.Columns[line][j])
			}
			if s.Boxes[box][j].Column != line {
				remainingBox = append(remainingBox, s.Boxes[box][j])
			}
		}
		str, ok = stepLockedCandidatesOneGroup(s, intersection, remainingLine, remainingBox)
		if ok {
			return str, ok
		}
		i++
		if i%s.Size == 0 {
			i += s.Size * (s.SmallSize - 1)
		}
	}

	return "", false
}

func stepLockedCandidatesOneGroup(s *SudokuBoard, intersection, remainingLine, remainingBox []*Cell) (string, bool) {

	for candidate := 1; candidate <= s.Size; candidate++ {
		if candidatePresentInCells(candidate, intersection) {
			inLine := candidatePresentInCells(candidate, remainingLine)
			inBox := candidatePresentInCells(candidate, remainingBox)
			if inLine != inBox {
				//We have found a situation of LockedCandidates!
				var b bytes.Buffer
				b.WriteString(fmt.Sprintf("Locked Candidates: for number %d and intersection ", candidate))
				for i := range intersection {
					b.WriteString(intersection[i].TextPosition())
				}
				b.WriteString(fmt.Sprint(". Removed from "))
				for i := range remainingLine {
					if _, exists := remainingLine[i].Candidates[candidate]; exists {
						delete(remainingLine[i].Candidates, candidate)
						b.WriteString(remainingLine[i].TextPosition())
					}
				}
				for i := range remainingBox {
					if _, exists := remainingBox[i].Candidates[candidate]; exists {
						delete(remainingBox[i].Candidates, candidate)
						b.WriteString(remainingBox[i].TextPosition())
					}
				}
				b.WriteString("\n")
				return b.String(), true
			}
		}
	}
	return "", false
}

func candidatePresentInCells(candidate int, array []*Cell) bool {

	for i := range array {
		_, o := array[i].Candidates[candidate]
		if o {
			return true
		}
	}
	return false
}
