package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Rock_Paper_Scissors(total_rounds int) {

	fmt.Println("Welcome to Rock, Paper, Scissors!")

	list := []string{"Rock", "Paper", "Scissors"}

	player_wins := 0
	rounds := 0
	player_option := ""
	comp_option := ""

	for {
		fmt.Println("")
		user_choice := input("Enter your choice([R/r]ock/ [P/p]aper/ [S/s]cissors): ")

		rounds++

		comp_option = list[generateRandom(3)]

		if user_choice == "R" || user_choice == "r" {
			player_option = "Rock"

			fmt.Println("User chose:", player_option)
			fmt.Println("Computer chose:", comp_option)

			output := DetermineWin(player_option, comp_option)

			if output == 1 {
				fmt.Println("The User Wins.")
				player_wins++
			} else if output == 0 {
				fmt.Println("It's a Draw.")
			} else {
				fmt.Println("The Computer Wins.")
			}

		} else if user_choice == "P" || user_choice == "p" {
			player_option = "Paper"

			fmt.Println("User chose:", player_option)
			fmt.Println("Computer chose:", comp_option)

			output := DetermineWin(player_option, comp_option)

			if output == 1 {
				fmt.Println("The User Wins.")
				player_wins++
			} else if output == 0 {
				fmt.Println("It's a Draw.")
			} else {
				fmt.Println("The Computer Wins.")
			}

		} else if user_choice == "S" || user_choice == "s" {
			player_option = "Scissors"

			fmt.Println("User chose:", player_option)
			fmt.Println("Computer chose:", comp_option)

			output := DetermineWin(player_option, comp_option)

			if output == 1 {
				fmt.Println("The User Wins.")
				player_wins++
			} else if output == 0 {
				fmt.Println("It's a Draw.")
			} else {
				fmt.Println("The Computer Wins.")
			}

		} else {
			fmt.Println("")
			fmt.Println("Enter one of the three options.")
			rounds--
		}

		if rounds == total_rounds {
			fmt.Println("")
			fmt.Printf("The User won %v out of %v games. \n", player_wins, total_rounds)
			return
		}
	}
}

func Guess_the_number() {

	tries := 0

	fmt.Println("Welcome to Guess The Number! ")

	b := input("Enter the maximum number possible[1 - number]: ")

	bound, err := strconv.Atoi(b)

	if err != nil {
		fmt.Println("Enter an integer.")
		return
	}

	number := generateRandom(bound) + 1

	for {
		fmt.Println("")
		u := input("Enter your guess: ")
		user_guess, err := strconv.Atoi(u)

		tries++

		if err != nil {
			fmt.Println("Enter an integer.")
			tries--

		} else if user_guess < number {
			fmt.Println("Your guess is Lower!")

		} else if user_guess > number {
			fmt.Println("Your guess is Higher!")

		} else if user_guess == number {
			fmt.Printf("You have correctly guessed the number in %v tries!", tries)
			return
		}
	}
}

func Hangman() {

	words := []string{"APPLE", "BANANA", "ORANGE", "MANGO", "PINEAPPLE", "CUSTARD", "WATERMELON"}
	WORD := words[generateRandom(len(words))]
	user_word := user_word_list(len(WORD))
	lives := 6

	indices := []int{}

	fmt.Println("Welcome to Hangman!")
	fmt.Printf("You have a total of %v lives. \n", lives)

	fmt.Printf("Your Word is: %v \n", user_word)

	for {

		guess := input("Enter a letter: ")

		indices = []int{}

		if len(guess) > 1 {
			fmt.Println("Enter only a single character.")
		} else {

			if strings.Contains(WORD, strings.ToUpper(guess)) {

				for i := 0; i < len(WORD); i++ {
					if string(WORD[i]) == strings.ToUpper(guess) {
						indices = append(indices, i)
					}
				}

				for _, v := range indices {
					user_word[v] = strings.ToUpper(guess)
				}

				fmt.Printf("The Word currently is: %v \n", user_word)

			} else {

				fmt.Println("This character doesn't exist in the Word.")
				lives--
				fmt.Printf("You have %v lives currently remaining. \n", lives)
			}
		}

		if ((strings.Join(user_word, "")) == WORD) && lives > 0 {
			fmt.Println("You have correctly guessed the Word!")
			return
		} else if lives <= 0 {
			fmt.Println("You Failed to correctly guess the Word.")
			fmt.Printf("The Word was: %v \n", WORD)
			return
		}
	}
}

func main() {

	fmt.Println("Welcome to Games!")

	for {

		fmt.Println("")
		fmt.Println("1. Rock, Paper, Scissors.")
		fmt.Println("2. Guess the Number.")
		fmt.Println("3. Hangman.")
		fmt.Println("4. Tic-Tac-Toe.")
		fmt.Println("5. Sudoku.")
		fmt.Println("6. Exit.")

		c := input("Enter your choice: ")

		choice, err := strconv.Atoi(c)

		if err != nil {
			fmt.Println("Enter a number.")
			return
		}

		switch choice {

		case 1:

			r := input("How many rounds: ")
			round, err := strconv.Atoi(r)

			if err != nil || (round < 1) {
				fmt.Println("Enter properly.")
				break
			}

			Rock_Paper_Scissors(round)

		case 2:

			Guess_the_number()

		case 3:

			Hangman()

		case 4:

			Tic_Tac_Toe()

		case 5:

			Sudoku()

		case 6:

			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Enter a number between 1 and 5.")
		}
	}
}
