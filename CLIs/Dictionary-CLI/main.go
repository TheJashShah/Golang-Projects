package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"

	"github.com/fatih/color"
)

func generateRandom() int {

	return rand.Intn(256)
}

type Word struct {
	Word     string `json:"word"`
	Meanings []struct {
		PartOfSpeech string `json:"partOfSpeech"`
		Definitions  []struct {
			Definition string   `json:"definition"`
			Synonyms   []string `json:"synonyms,omitempty"`
			Antonyms   []string `json:"antonyms,omitempty"`
		} `json:"definitions"`
		Synonyms []string `json:"synonyms,omitempty"`
		Antonyms []string `json:"antonyms,omitempty"`
	} `json:"meanings"`
}

func main() {

	word := "Hello"

	if len(os.Args) >= 2 {
		word = os.Args[1]
	}

	res, err := http.Get("https://api.dictionaryapi.dev/api/v2/entries/en/" + word)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Not Avaiable Currently.")
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	var words []Word
	err = json.Unmarshal(body, &words)

	if err != nil {
		panic(err)
	}

	color := color.RGB(generateRandom(), generateRandom(), generateRandom())

	for _, word := range words {
		color.Printf("Word: %v \n", word.Word)

		for _, meaning := range word.Meanings {
			color.Printf("   Part Of Speech: %v \n", meaning.PartOfSpeech)

			for _, definition := range meaning.Definitions {
				color.Printf("      Definition: %v \n", definition.Definition)

				if len(definition.Synonyms) > 0 {
					color.Printf("         Synonyms: \n")
					for _, syn := range definition.Synonyms {
						color.Printf("         %v \n", syn)
					}
				}

				if len(definition.Antonyms) > 0 {
					color.Printf("         Antonyms: \n")
					for _, ant := range definition.Antonyms {
						color.Printf("         %v \n", ant)
					}
				}
			}

			if len(meaning.Synonyms) > 0 {
				color.Printf("   Synonyms: \n")

				for _, syn := range meaning.Synonyms {
					color.Printf("   %v \n", syn)
				}
			}

			if len(meaning.Antonyms) > 0 {
				color.Printf("   Antonyms: \n")

				for _, ant := range meaning.Antonyms {
					color.Printf("   %v \n", ant)
				}
			}

			fmt.Println("")
		}
	}

}
