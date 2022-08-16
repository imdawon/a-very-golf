package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	letters, err := getValidLettersInput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(letters)
}

// Validate the user is entering a character string with a length of at least 3.
func getValidLettersInput() (string, error) {
	if len(os.Args) > 1 {
		letters := os.Args[1]
		if len(letters) >= 3 {
			return letters, nil
		} else {
			return "", errors.New("Please enter at least three letters for anagram generation.")
		}
	} else {
		return "", errors.New("Please enter at least three letters for anagram generation.")
	}
}
