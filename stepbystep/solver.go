package stepbystep

import (
	"bytes"
	"errors"
	"fmt"
	"sync"
	"time"
)

var steps []func(*SudokuBoard) (string, bool)

type stepDescription struct {
	sudoku      string
	description string
}

const complexityLogged = 0
const showBoardLog = true

var errNoSolution error = errors.New("This soduku has no solution")
var errMultipleSolutions error = errors.New("This soduku has muliple solutions")
var errAlreadySolved error = errors.New("The sudoku is solved")

//Solver solves the sudoku and returns a description on how was it solved.
func Solver(s *SudokuBoard) string {

	initLogs := make([]stepDescription, 1)
	initLogs[0].sudoku = s.ToString()
	initLogs[0].description = "Initial status\n"
	startTime := time.Now()
	logs, _ := solver(s)
	elapsedTime := time.Since(startTime)
	logs = append(initLogs, logs...)
	var b bytes.Buffer
	for k := range logs {
		b.WriteString(logs[k].description)
		if showBoardLog {
			b.WriteString(logs[k].sudoku)
		}
		b.WriteString("\n")
	}
	b.WriteString(fmt.Sprintf("Time for solution %s\n", elapsedTime.String()))
	return b.String()

}

func solver(s *SudokuBoard) ([]stepDescription, error) {
	steps := []func(*SudokuBoard) (string, bool){
		stepNakedSingle,
		stepHiddenSingle,
		stepLockedCandidates,
	}
	maxSteps := len(steps)
	k := 0
	var description string
	var result bool
	var currentStepLog stepDescription
	logs := make([]stepDescription, 0)
	// currentStepLog.sudoku = s.ToString()
	// currentStepLog.description = "Initial status\n"
	// logs = append(logs, currentStepLog)
	for k < maxSteps {
		description, result = steps[k](s)
		if result {
			currentStepLog.sudoku = s.ToString()
			currentStepLog.description = description
			if k >= complexityLogged {
				logs = append(logs, currentStepLog)
			}
			k = 0
		} else {
			k++
		}
		if !s.Check() {
			currentStepLog.sudoku = s.ToString()
			currentStepLog.description = "At this point the Sudoku is incongruent. It has no solution.\n"
			logs = append(logs, currentStepLog)
			return logs, errNoSolution
		}
	}
	logsGuess, errGuess := guess(s)
	logs = append(logs, logsGuess...)
	return logs, errGuess
}

func guess(s *SudokuBoard) ([]stepDescription, error) {

	//fmt.Print(s.ToString())
	var wg sync.WaitGroup

	for i := range s.ListCells {
		if s.ListCells[i].Value == 0 {
			numCandidates := len(s.ListCells[i].Candidates)
			candidates := make([]int, numCandidates)
			logs := make([][]stepDescription, numCandidates)
			errs := make([]error, numCandidates)
			j := 0
			wg.Add(numCandidates)
			for candidate := range s.ListCells[i].Candidates {

				// fmt.Printf("**** Starting to guess candidate %d for position %s\n", candidate, s.ListCells[i].TextPosition())
				// if s.ListCells[1].Value == 1 && s.ListCells[2].Value == 2 && s.ListCells[3].Value == 7 && s.ListCells[4].Value == 5 && s.ListCells[5].Value == 3 && s.ListCells[6].Value == 6 && s.ListCells[7].Value == 4 && s.ListCells[13].Value == 8 && s.ListCells[15].Value == 1 && i == 27 && candidate == 1 {
				// 	fmt.Println(s.ToString())
				// 	fmt.Println("****")
				// }
				candidates[j] = candidate
				logs[j] = make([]stepDescription, 1)
				//This block could be concurrent.
				go func(t *SudokuBoard, candidate int, position int, logs *[]stepDescription, err *error) {
					defer wg.Done()

					t.ListCells[position].Set(candidate)
					(*logs)[0].description = fmt.Sprintf("Guess: %s -> %d\n", t.ListCells[position].TextPosition(), candidate)
					(*logs)[0].sudoku = t.ToString()
					var logsSolver []stepDescription
					logsSolver, *err = solver(t)
					(*logs) = append(*logs, logsSolver...)

				}(s.Duplicate(), candidate, i, &logs[j], &errs[j])
				j++
			}
			// fmt.Printf("**** Results for position %s\n", s.ListCells[i].TextPosition())
			wg.Wait()
			candidatesSolved := make([]int, 0)
			for candidate := range candidates {
				//Check if any has multiple solutions

				if errs[candidate] == errMultipleSolutions {
					return logs[candidate], errs[candidate]
				}
				//Check if more than one has multiple solutions
				if errs[candidate] == errAlreadySolved {
					candidatesSolved = append(candidatesSolved, candidate)
				}
			}
			// fmt.Println("**** Candidates ", candidatesSolved)
			if len(candidatesSolved) == 0 {
				// fmt.Println("**** No solution")
				return logs[0], errNoSolution

			}
			if len(candidatesSolved) == 1 {
				// fmt.Println("**** Solution found")
				return logs[candidatesSolved[0]], errAlreadySolved
			}
			// fmt.Printf("**** Multiple solutions found")
			return logs[candidatesSolved[0]], errMultipleSolutions
		}
	}
	return nil, errAlreadySolved
}
