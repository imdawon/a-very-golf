package main

import (
	"a-very-golf/hashmap"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	// Ensure we have valid amount of letters.
	letters, err := getValidLettersInput()
	if err != nil {
		log.Fatal(err)
	}

	// High bound of spaces we can inject into our permutation strings.
	maxSpacesNum := getMaxNumOfSpaces(letters)

	fmt.Printf("Letters: %s\n", letters)
	fmt.Printf("Max spaces count: %f \n", maxSpacesNum)
	validWords := hashmap.GetValidWords()

	fmt.Printf("valid word?: %d", validWords["assess2"])
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

func getMaxNumOfSpaces(letters string) float64 {
	if len(letters) == 3 {
		return 0
	}

	ratio := float64(len(letters) / 2)
	return (math.Floor(ratio))
}
