package main

import (
	"a-very-golf/anagram"
	"a-very-golf/hashmap"
	"a-very-golf/permutation"
	"fmt"
	"log"
	"os"
)

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

	hashmap.AnagramCandidates = make(map[string]int)

	// High bound count of spaces we can inject into our permutation strings.
	maxSpaces := anagram.GetMaxNumOfSpaces(unsortedAnagram)

	// Inject max number of spaces into our user input string.
	unsortedAnagramWithSpaces := anagram.AddSpaces(unsortedAnagram, int(maxSpaces))

	// Generate anagrams and extract valid strings of English words.
	permutation.FindAnagramCandidates(unsortedAnagramWithSpaces, 0, len(unsortedAnagramWithSpaces))

	// Print anagrams.
	if len(hashmap.AnagramCandidates) > 0 {
		for anagram := range hashmap.AnagramCandidates {
			fmt.Println(anagram)
		}
		// We didn't find any anagrams.
	} else {
		fmt.Println("Not a single anagram was found. oof.")
	}

}
