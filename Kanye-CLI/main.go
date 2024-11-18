package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/fatih/color"
)

func generateRandom() int {

	return rand.Intn(256)
}

func main() {

	res, err := http.Get("https://api.kanye.rest")

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Not Available Currently.")
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	var data map[string]string
	err = json.Unmarshal(body, &data)

	if err != nil {
		panic(err)
	}

	final_quote := data["quote"]
	message := fmt.Sprintf("Kanye says: %v", final_quote)

	color.RGB(generateRandom(), generateRandom(), generateRandom()).Println(message)
}
