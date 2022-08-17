package anagram

import (
	"errors"
	"math"
	"os"
	"strings"
)

func AddSpaces(unsortedAnagram string, maxSpaces int) string {
	unsortedAnagramWithSpaces := unsortedAnagram
	for i := 0; i < maxSpaces; i++ {
		unsortedAnagramWithSpaces = unsortedAnagramWithSpaces + string(" ")
	}

	return unsortedAnagramWithSpaces
}

// Validate the user is entering a character string with a length of at least 3.
func IsValidInput() (bool, error) {
	if len(os.Args) > 1 {
		letters := os.Args[1]
		if len(letters) >= 3 {
			return true, nil
		} else {
			return false, errors.New("Please enter at least three letters for anagram generation.")
		}
	} else {
		return false, errors.New("Please enter at least three letters for anagram generation.")
	}
}

func GetMaxNumOfSpaces(letters string) float64 {
	if len(letters) == 3 {
		return 0
	}

	ratio := float64(len(letters) / 2)
	return (math.Floor(ratio))
}

func RemoveMultipleAdjacentSpaces(str string) string {
	return strings.Join(strings.Fields(str), " ")
}
