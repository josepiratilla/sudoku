package stepbystep

import (
	"bytes"
)

var steps []func(*SudokuBoard) (string, bool)

type stepDescription struct {
	sudoku      string
	description string
}

//Solver solves the sudoku and returns a description on how was it solved.
func Solver(s *SudokuBoard) string {
	steps := []func(*SudokuBoard) (string, bool){
		stepOnly1Candidate,
		stepOnly1OptionInHouse,
	}
	maxSteps := len(steps)
	k := 0
	var description string
	var result bool
	var currentStepLog stepDescription
	logs := make([]stepDescription, 0)
	currentStepLog.sudoku = s.ToString()
	currentStepLog.description = "Initial status\n"
	logs = append(logs, currentStepLog)
	for k < maxSteps {
		description, result = steps[k](s)
		if result {
			currentStepLog.sudoku = s.ToString()
			currentStepLog.description = description
			logs = append(logs, currentStepLog)
			k = 0
		} else {
			k++
		}
	}
	var b bytes.Buffer
	for k := range logs {
		b.WriteString(logs[k].description)
		b.WriteString(logs[k].sudoku)
		b.WriteString("\n")
	}
	return b.String()
}
