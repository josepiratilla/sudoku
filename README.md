# Sudoku
Sodoku solver in Go. I've done this project just to learn go and put in practice some concepts.

## Packages
The project contains three packages:
* **models**: A basic sudoku solver by backtracking. It can only accept 9x9 sudokus. Pointer and matrix use.
* **generic**: A generic sudoku solver by backtracking. It accepts any square sudoku size. Since everything is dynamic, I needed to make extensive use of slices.
* **stepbystep**: A verbose sudoku solver. It can apply different strategies, and when no more are available it starts to guess (which is backtracking at the end). Guessing is done with multiple candidates concurrently. Also, the steps can be added easily making use of arrays of anonymous functions. 

## Step by Step features

The package contains a much more sophisticated sudoku board structure. It contains multiple slices pointing to the same cells, so any cell can be obtained as an ordered list, 2 dimensional matrix, per rows, per columns, per boxes, or inderctly from a cell, as the neightbourgs of that cell (cells in the same row, column or box)
Also each cell, not only contains the value (when 0 is unknown) but also the possible candidates that have not been discarded.
### Solvers included

#### Naked Candidate
When there is only one possible candidate in a cell, it is converted into a found value.

#### Hidden Candidate
When a missing value in a house has a candidate only in one cell, that cell should have that value.
In the next example in the central row, the value 8 has only one candidate at the 4th column. So it must be 8 for that cell.
```
+-------------------------+-------------------------+-------------------------+
|    8       1       2    |    7       5       3    |    6       4       9    | 
|    9       4       3    |    6       8       2    |    1       7       5    | 
|    6       7       5    |    4       9       1    |    2       8       3    | 
+-------------------------+-------------------------+-------------------------+
|    1       5     (4,9)  | (2,3,9) (2,3,6)    7    |    8     (2,9)   (4,6)  | 
|  (2,3)   (3,6)   (6,9)  | (2,8,9)    4       5    |    7     (2,9)     1    | 
|  (2,7)     8    (4,7,9) |    1     (2,6)   (6,9)  |    5       3     (4,6)  | 
+-------------------------+-------------------------+-------------------------+
|    5     (2,3)     1    | (2,3,9)    7       4    |  (3,9)     6       8    | 
|    4    (2,3,6)    8    |    5    (2,3,6)  (6,9)  |  (3,9)     1       7    | 
|  (3,7)     9     (6,7)  |  (3,8)     1     (6,8)  |    4       5       2    | 
+-------------------------+-------------------------+-------------------------+

Hidden Single at row 5: [r5c4] -> 8
+-------------------------+-------------------------+-------------------------+
|    8       1       2    |    7       5       3    |    6       4       9    | 
|    9       4       3    |    6       8       2    |    1       7       5    | 
|    6       7       5    |    4       9       1    |    2       8       3    | 
+-------------------------+-------------------------+-------------------------+
|    1       5     (4,9)  | (2,3,9) (2,3,6)    7    |    8     (2,9)   (4,6)  | 
|  (2,3)   (3,6)   (6,9)  |    8       4       5    |    7     (2,9)     1    | 
|  (2,7)     8    (4,7,9) |    1     (2,6)   (6,9)  |    5       3     (4,6)  | 
+-------------------------+-------------------------+-------------------------+
|    5     (2,3)     1    | (2,3,9)    7       4    |  (3,9)     6       8    | 
|    4    (2,3,6)    8    |    5    (2,3,6)  (6,9)  |  (3,9)     1       7    | 
|  (3,7)     9     (6,7)  |   (3)      1     (6,8)  |    4       5       2    | 
+-------------------------+-------------------------+-------------------------+
```

#### Locked Pairs
In an intersection between a line(row or column) and a box if a candidate is present in the intersection and only the line or the box, all options outside the intersection are removed.
In the next example we consider the intersection between the 4th row and the central box. The number 3 only appears at the intersection of both and in the box. So the candidates for 3 in the box, out of the intersection can be removed.
```
+-------------------------------+-------------------------------+-------------------------------+
|     8         1         2     |     7         5         3     |     6         4         9     | 
|     9         4         3     |     6         8         2     |     1         7         5     | 
|     6         7         5     |     4         9         1     |     2         8         3     | 
+-------------------------------+-------------------------------+-------------------------------+
|     1         5      (4,6,9)  |  (2,3,9)   (2,3,6)      7     |     8       (2,9)     (4,6)   | 
|   (2,3)    (3,6,8)    (6,9)   | (2,3,8,9)     4         5     |     7       (2,9)       1     | 
|   (2,7)     (6,8)   (4,6,7,9) |     1       (2,6)    (6,8,9)  |     5         3       (4,6)   | 
+-------------------------------+-------------------------------+-------------------------------+
|     5       (2,3)       1     |  (2,3,9)      7         4     |   (3,9)       6         8     | 
|     4      (2,3,6)      8     |     5      (2,3,6)    (6,9)   |   (3,9)       1         7     | 
|   (3,7)       9       (6,7)   |   (3,8)       1       (6,8)   |     4         5         2     | 
+-------------------------------+-------------------------------+-------------------------------+

Locked Candidates: for number 3 and intersection [r4c4][r4c5][r4c6]. Removed from [r5c4]
+-------------------------------+-------------------------------+-------------------------------+
|     8         1         2     |     7         5         3     |     6         4         9     | 
|     9         4         3     |     6         8         2     |     1         7         5     | 
|     6         7         5     |     4         9         1     |     2         8         3     | 
+-------------------------------+-------------------------------+-------------------------------+
|     1         5      (4,6,9)  |  (2,3,9)   (2,3,6)      7     |     8       (2,9)     (4,6)   | 
|   (2,3)    (3,6,8)    (6,9)   |  (2,8,9)      4         5     |     7       (2,9)       1     | 
|   (2,7)     (6,8)   (4,6,7,9) |     1       (2,6)    (6,8,9)  |     5         3       (4,6)   | 
+-------------------------------+-------------------------------+-------------------------------+
|     5       (2,3)       1     |  (2,3,9)      7         4     |   (3,9)       6         8     | 
|     4      (2,3,6)      8     |     5      (2,3,6)    (6,9)   |   (3,9)       1         7     | 
|   (3,7)       9       (6,7)   |   (3,8)       1       (6,8)   |     4         5         2     | 
+-------------------------------+-------------------------------+-------------------------------+
```

### Creating new solving steps
New steps can be developed by creating an step function. Steps functions are placed at `steps.go` file. The fuctions has a sudokuBoard as entry and returns a strings, explaining the modification done and a bool, that will be true if the step has found a reduction. The sudoku board is modified by the function with the reduction.
Example:
```go
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
```
Steps must then be added to the array indicating the solving steps at `solver.go` in the function `solver`:
```go
func solver(s *SudokuBoard) ([]stepDescription, error) {
	steps := []func(*SudokuBoard) (string, bool){
		stepNakedSingle,
		stepHiddenSingle,
		stepLockedCandidates,
    }
    ...
```
The upper in the list, the more priority for using this step.