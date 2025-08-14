package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: no word provided")
	}
	words := os.Args[1:]

	for _, word := range words {
		resp, err := http.Get("https://api.dictionaryapi.dev/api/v2/entries/en/" + word)
		if err != nil {
			fmt.Println(err)
		}
		body, err := io.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			fmt.Println(err)
		}

		d := Definition{}
		err = json.Unmarshal(body, &d)
		if err != nil {
			fmt.Println(err)
		}

		output := d[0].Meanings[0].Definitions[0].Definition
		fmt.Println(word + ": " + output)
	}
}
