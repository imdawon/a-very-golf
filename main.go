package main

import (
	"a-very-golf/anagram"
	"a-very-golf/hashmap"
	"a-very-golf/permutation"
	"fmt"
	"log"
	"os"
)

var validWords = hashmap.GetValidWords()

func main() {
	// Ensure user entered valid unsorted anagram.
	success, err := anagram.IsValidInput()
	if err != nil {
		log.Fatal(err)
	}

	var unsortedAnagram string

	if success {
		unsortedAnagram = os.Args[1]
	}

	// High bound of spaces we can inject into our permutation strings.
	maxSpaces := anagram.GetMaxNumOfSpaces(unsortedAnagram)

	unsortedAnagramWithSpaces := anagram.AddSpaces(unsortedAnagram, int(maxSpaces))

	fmt.Printf("Letters: %s\n", unsortedAnagramWithSpaces)
	fmt.Printf("Max spaces count: %f\n", maxSpaces)

	fmt.Printf("valid word?: %d\n", validWords["assess"])
	fmt.Printf("stripped adjacent spaces: %s\n", anagram.RemoveMultipleAdjacentSpaces(" |test  test2  |"))

	anagramCandidates := permutation.GetAnagramCandidates(unsortedAnagramWithSpaces, 0, len(unsortedAnagramWithSpaces))

	if len(anagramCandidates) > 0 {
		for _, anagram := range anagramCandidates {
			fmt.Println(anagram)
		}
	} else {
		fmt.Println("Not a single anagram was found. oof.")
	}

}
