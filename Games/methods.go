package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func generateRandom(bound int) int {

	return rand.Intn(bound)
}

func input(prompt string) string {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print(prompt)

	i, _ := reader.ReadString('\n')

	input := strings.TrimSpace(i)

	return input
}

func DetermineWin(user string, comp string) int {

	output := 0
	flag := false

	win_map := map[string]string{
		"Rock":     "Paper",
		"Paper":    "Scissors",
		"Scissors": "Rock",
	}

	if user == comp {
		output = 0
	}

	for k, v := range win_map {
		if k == user && v == comp {
			flag = true
			output = -1
		}
	}

	if !flag && (user != comp) {
		output = 1
	}

	return output
}

func user_word_list(len int) []string {

	list := []string{}

	for i := 0; i < len; i++ {
		list = append(list, "_")
	}

	return list
}
