package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DisplayBoard(board [9][9]int) {

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != 0 {
				fmt.Printf(" | %v ", board[i][j])
			} else {
				fmt.Printf(" | %v ", " ")
			}
		}
		fmt.Printf("\n---------------------------------------------\n")
	}
}

func LoadBoard(board *[9][9]int, rows *[]int, cols *[]int) {

	file, err := os.Open("sudoku.txt")

	if err != nil {
		fmt.Println("Error in Opening File.")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) < 3 {
			fmt.Println("Error in Line.")
			continue
		}

		num, err := strconv.Atoi(fields[0])

		if err != nil {
			fmt.Println("Error in converting number.")
			continue
		}

		row, err := strconv.Atoi(fields[1])

		if err != nil {
			fmt.Println("Error in converting number.")
			continue
		}

		col, err := strconv.Atoi(fields[2])

		if err != nil {
			fmt.Println("Error in converting number.")
			continue
		}

		board[row][col] = num

		*rows = append(*rows, row)
		*cols = append(*cols, col)
	}
}

func getBounds(pos int, lower *int, upper *int) {

	mod := pos % 3

	switch mod {

	case 0:
		*lower = pos
		*upper = (pos + 1)

	case 1:
		*lower = (pos - 1)
		*upper = (pos + 1)

	case 2:
		*lower = (pos - 2)
		*upper = pos
	}
}

func isValidnum(num int, row int, col int, board [9][9]int) bool {

	for i := 0; i < 9; i++ {
		if board[row-1][i] == num {
			return false
		}
	}
	// ROW

	for i := 0; i < 9; i++ {
		if board[i][col-1] == num {
			return false
		}
	}
	// COL

	lower_row := 0
	upper_row := 0
	lower_col := 0
	upper_col := 0

	getBounds(row-1, &lower_row, &upper_row)
	getBounds(col-1, &lower_col, &upper_col)

	for j := lower_row; j <= upper_row; j++ {
		for k := lower_col; k <= upper_col; k++ {
			if board[j][k] == num {
				return false
			}
		}
	}
	// 3x3 GRID

	return true
}

func insertNum(board *[9][9]int, rows []int, cols []int) {

	fmt.Println("")
	r := input("Enter Row Number[1 - 9]: ")
	c := input("Enter Col Number[1 - 9]: ")
	n := input("Enter Number[1 - 9]: ")

	row, err_1 := strconv.Atoi(r)
	col, err_2 := strconv.Atoi(c)
	num, err_3 := strconv.Atoi(n)

	flag := 0

	if err_1 != nil || err_2 != nil || err_3 != nil {

		fmt.Println("There is an error in your inputs.")
		insertNum(board, rows, cols)

	} else if (row > 9 || row < 1) || (col > 9 || col < 1) || (num > 9 || num < 1) {

		fmt.Println("Each of these numbers should be between 1 and 9.")
		insertNum(board, rows, cols)

	} else {

		for i := 0; i < len(rows); i++ {
			if rows[i] == row-1 && cols[i] == col-1 {
				flag = 1
				break
			}
		}

		if flag == 1 {

			fmt.Println("You cannot change an initial number.")
			insertNum(board, rows, cols)

		} else {

			if isValidnum(num, row, col, *board) {
				board[row-1][col-1] = num

				DisplayBoard(*board)

			} else {
				fmt.Println("Wrong Move.")
				insertNum(board, rows, cols)
			}
		}
	}
}

func getCount(board [9][9]int) int {

	count := 0

	for _, v := range board {
		for _, k := range v {
			if k == 0 {
				count++
			}
		}
	}

	return count
}

func SolveBoard(board *[9][9]int, row int, col int) int {

	if row == 8 && col == 9 {
		return 1
	}

	if col == 9 {
		row++
		col = 0
	}

	if board[row][col] != 0 {
		return SolveBoard(board, row, col+1)
	}

	for i := 1; i <= 9; i++ {
		if isValidnum(i, row+1, col+1, *board) {

			board[row][col] = i

			if SolveBoard(board, row, col+1) == 1 {
				return 1
			}

			board[row][col] = 0

			DisplayBoard(*board)
			fmt.Println("")
		}
	}

	return 0
}

func Sudoku() {

	board := [9][9]int{}

	rows := []int{}
	cols := []int{}

	LoadBoard(&board, &rows, &cols)

	fmt.Println("Welcome to Sudoku!")

	DisplayBoard(board)
	fmt.Println("  ")

	fmt.Println("")
	fmt.Println("1. Solve the Board.")
	fmt.Println("2. Make the Computer Solve.")

	c := input("Enter your choice: ")

	choice, err := strconv.Atoi(c)

	if err != nil {
		fmt.Println("Enter a number.")

	} else {

		switch choice {

		case 1:

			for {

				insertNum(&board, rows, cols)

				if getCount(board) == 0 {
					fmt.Println("You have solved the Board.")
					return
				}
			}

		case 2:

			if SolveBoard(&board, 0, 0) == 1 {

				DisplayBoard(board)
				fmt.Println("The Board is Solved.")
			}

		default:
			fmt.Println("Enter a number between 1 and 2.")
		}
	}
}
