package main

import (
	"os"

	"github.com/01-edu/z01"
)

func IsValid(board [][]rune, row, col int, c rune) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] == c {
			return false
		}
		if board[row][i] == c {
			return false
		}
		if board[3*(row/3)+(i/3)][3*(col/3)+(i%3)] == c {
			return false
		}
	}
	return true
}

func Insert(args []string) [][]rune {
	var res [][]rune
	for i := 0; i < len(args); i++ {
		res = append(res, []rune(args[i]))
	}
	if !CheckInsert(res) {
		return res
	}
	return res
}

func CheckInsert(resOfInsert [][]rune) bool {
	for row := 0; row < len(resOfInsert); row++ {
		for col := 0; col < len(resOfInsert[row]); col++ {
			if resOfInsert[row][col] != '.' {
				for z := 0; z < 9; z++ {
					if z != row {
						if resOfInsert[z][col] == resOfInsert[row][col] {
							return false
						}
					}
					if z != col {
						if resOfInsert[row][z] == resOfInsert[row][col] {
							return false
						}
					}
					if 3*(row/3)+(z/3) != row && 3*(col/3)+(z%3) != col {
						if resOfInsert[3*(row/3)+(z/3)][3*(col/3)+(z%3)] == resOfInsert[row][col] {
							return false
						}
					}
				}
			}
		}
	}
	return true
}

func PrintBoard(board [][]rune) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			z01.PrintRune(board[i][j])
			if j == 8 {
				break
			}
			z01.PrintRune(' ')
		}
		z01.PrintRune('\n')
	}
}

func SudokuSolver(board [][]rune) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for k := '1'; k <= '9'; k++ {
					if IsValid(board, i, j, k) {
						board[i][j] = k
						if SudokuSolver(board) {
							return true
						} else {
							board[i][j] = '.'
						}
					}
				}
				return false
			}
		}
	}
	return true
}

func PrintError() {
	err := "Error"
	for _, r := range err {
		z01.PrintRune(rune(r))
	}
	z01.PrintRune('\n')
}

func CorrectSudoku(osArg []string) bool {
	if len(osArg) != 10 {
		return false
	}
	nums := 0
	for i := 1; i < len(osArg); i++ {
		if len(osArg[i]) != 9 {
			return false
		}
		for j := 0; j < len(osArg[i]); j++ {
			if rune(osArg[i][j]) <= '9' && rune(osArg[i][j]) >= '1' {
				nums++
			} else if rune(osArg[i][j]) == '.' {
			} else {
				return false
			}
		}
	}
	if nums <= 16 {
		return true
	}
	return true
}

func main() {
	if !CorrectSudoku(os.Args) {
		PrintError()
		return
	}
	args := os.Args[1:]
	board := Insert(args)
	if board[0][0] == '$' {
		PrintError()
		return
	}
	SudokuSolver(board)
	if !CheckInsert(board) {
		PrintError()
		return
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				PrintError()
				return
			}
		}
	}
	PrintBoard(board)
}