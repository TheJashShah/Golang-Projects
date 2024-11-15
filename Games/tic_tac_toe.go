package main

import (
	"fmt"
	"strconv"
)

func displayBoard(board []string) {

	fmt.Printf(" %v | %v | %v \n", board[1], board[2], board[3])
	fmt.Print("---|---|--- \n")
	fmt.Printf(" %v | %v | %v \n", board[4], board[5], board[6])
	fmt.Print("---|---|--- \n")
	fmt.Printf(" %v | %v | %v \n", board[7], board[8], board[9])
}

func getInput(marker string, board []string) {

	input_sentence := fmt.Sprintf("Player with %v, enter a position between 1-9: ", marker)

	n := input(input_sentence)

	num, err := strconv.Atoi(n)

	if err != nil || (num < 1 || num > 9) {
		fmt.Println("Enter a number only between 1 and 9.")
		getInput(marker, board)

	} else if board[num] != " " {
		fmt.Println("This position is already occupied.")
		getInput(marker, board)

	} else {
		board[num] = marker
		displayBoard(board)
	}
}

func IsWinner(board []string, marker string) int {

	is_winner := 0

	if (board[1] == board[2] && board[2] == board[3] && board[1] == marker) ||
		(board[4] == board[5] && board[5] == board[6] && board[4] == marker) ||
		(board[7] == board[8] && board[8] == board[9] && board[9] == marker) ||
		(board[1] == board[4] && board[4] == board[7] && board[1] == marker) ||
		(board[2] == board[5] && board[5] == board[8] && board[2] == marker) ||
		(board[3] == board[6] && board[6] == board[9] && board[3] == marker) ||
		(board[1] == board[5] && board[5] == board[9] && board[1] == marker) ||
		(board[3] == board[5] && board[5] == board[7] && board[3] == marker) {

		is_winner = 1
	}

	return is_winner
}

func getFrequency(list []string, char string) int {

	freq := 0

	for _, V := range list {
		if V == char {
			freq++
		}
	}

	return freq
}

func isDraw(board []string, player_1 string, player_2 string) int {

	is_draw := 0

	if IsWinner(board, player_1) == 0 && IsWinner(board, player_2) == 0 && getFrequency(board, " ") <= 1 {
		is_draw = 1
	}

	return is_draw
}

func Tic_Tac_Toe() {

	board := []string{"", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "}

	fmt.Println("Welcome to Two-Player Tic-Tac-Toe!")
	fmt.Println("")

	player_one_choice := input("Player 1, enter your marker[X/O]: ")

	player_one_marker := ""
	player_two_marker := ""

	if player_one_choice == "X" || player_one_choice == "x" {

		player_one_marker = "X"
		player_two_marker = "O"

	} else if player_one_choice == "O" || player_one_choice == "o" {

		player_one_marker = "O"
		player_two_marker = "X"

	} else {

		player_one_marker = "&"
		player_two_marker = "#"

	}

	fmt.Printf("Player One is represented by: %v \n", player_one_marker)
	fmt.Printf("Player Two is represented by: %v \n", player_two_marker)
	fmt.Println("")

	for {

		getInput(player_one_marker, board)

		if IsWinner(board, player_one_marker) == 1 {
			fmt.Println("Player One Wins!")
			return
		}

		if isDraw(board, player_one_marker, player_two_marker) == 1 {
			fmt.Println("It's a Draw!")
			return
		}

		getInput(player_two_marker, board)

		if IsWinner(board, player_two_marker) == 1 {
			fmt.Println("Player Two Wins!")
			return
		}

		if isDraw(board, player_one_marker, player_two_marker) == 1 {
			fmt.Println("It's a Draw!")
			return
		}
	}

}
