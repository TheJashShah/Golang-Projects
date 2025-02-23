package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/fatih/color"
)

func generateRandom() int {

	return rand.Intn(256)
}

func LoadTextFile(Hashmap map[string]string) {

	file, err := os.Open("C:/Users/Jash/OneDrive/Desktop/Golang-Projects/Console-Projects/Morse-CLI/morse_to_text.txt")

	if err != nil {
		log.Fatal("Error in Opening File.")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()
		fields := strings.Fields(line)

		Hashmap[fields[0]] = fields[1]
	}
}

func getKeyfromVal(Hashmap map[string]string, Val string) (string, int) {

	for k, v := range Hashmap {
		if v == Val {
			return k, 0
		}
	}

	return "", -1
}

func TextToMorse(Hashmap map[string]string, String string) string {

	result := ""

	for _, character := range String {

		translation, ok := Hashmap[string(character)]

		translation += " "

		if ok {
			result += translation

		} else {
			if character == ' ' {
				result += "/ "
			}
		}
	}

	return result
}

func MorseToText(Hashmap map[string]string, String string) string {

	result := ""

	morseArray := strings.Split(String, " ")

	for _, char := range morseArray {

		if char == "/" {
			result += " "

		} else {

			res, err := getKeyfromVal(Hashmap, char)

			if err != -1 {
				result += res
			}
		}
	}

	return result
}

func main() {

	commands := []string{"--help", "--morse", "--text"}

	var String string

	TextMorse := make(map[string]string)

	LoadTextFile(TextMorse)

	fmt.Println()
	fmt.Println()

	if len(os.Args) == 1 {

		fmt.Println("Here are the commands for the Morse-CLI: ")

		for _, command := range commands {

			color := color.RGB(generateRandom(), generateRandom(), generateRandom())

			color.Println(command)
		}

	} else if len(os.Args) == 2 {

		if os.Args[1] == commands[0] {
			color := color.RGB(generateRandom(), generateRandom(), generateRandom())

			color.Println("Welcome to Morse-CLI!")
			color.Println("--morse translates from Text to Morse.")
			color.Println("--text translates from Morse to Text.")
			color.Println("Enter --morse or --text along with the characters in a single string to get the decoding/encoding.")

		} else if os.Args[1] == commands[1] || os.Args[1] == commands[2] {

			fmt.Println("Enter the characters to translate.")

		} else {
			fmt.Println("Enter your command correctly.")
		}

	} else if len(os.Args) == 3 {

		if os.Args[1] == commands[1] {
			String = os.Args[2]

			fmt.Println("Here is your translation: ")
			result := TextToMorse(TextMorse, strings.ToUpper(String))
			fmt.Println(result)

		} else if os.Args[1] == commands[2] {
			String = os.Args[2]

			fmt.Println("Here is your translation: ")
			result := MorseToText(TextMorse, String)
			fmt.Println(result)

		} else {
			fmt.Println("Enter your command correctly.")
		}

	} else {
		fmt.Print("You have entered an incorrect amount of arguments.")
	}
}
