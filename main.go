package main

import (
	"a-very-golf/anagram"
	"a-very-golf/hashmap"
	"a-very-golf/permutation"
	"fmt"
	"log"
	"os"
	"time"
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

	timeStart := time.Now()
	// Generate anagrams and extract valid strings of English words.
	permutation.FindAnagramCandidates(unsortedAnagram, 0, len(unsortedAnagram))

	wordsPerLine := 6
	wordsPrinted := 0
	var wordsToPrint []string

	// Print anagrams.
	if len(hashmap.AnagramCandidates) > 0 {
		anagramsCount := len(hashmap.AnagramCandidates) - 1
		i := 0
		for anagram := range hashmap.AnagramCandidates {
			wordsToPrint = append(wordsToPrint, anagram)
			wordsPrinted++

			if wordsPrinted%wordsPerLine == 0 {
				fmt.Println(assembleAnagramsOnSameLine(wordsToPrint))
				// Empty our words to print slice.
				wordsToPrint = nil
				// Print all remaining words if we are iterating over the last entry in `AnagramCandidates`.
			} else if anagramsCount == i {
				fmt.Println(assembleAnagramsOnSameLine(wordsToPrint))
			}
			i++
		}

		timeFinish := time.Now()
		fmt.Printf("Time elapsed: %s\n", timeFinish.Sub(timeStart))

		// We didn't find any anagrams.
	} else {
		fmt.Println("Not a single anagram was found. oof.")
	}

}

// Assembles however many words are in the given slice on one string.
// We can outut more results into a grid in the terminal this way.
func assembleAnagramsOnSameLine(anagrams []string) string {
	if len(anagrams) == 0 {
		return ""
	}

	var output string
	wordsOnLine := 0
	for _, anagram := range anagrams {
		if wordsOnLine == 0 {
			output = anagram
		} else {
			output += "\t" + anagram
		}
		wordsOnLine++
	}

	return output
}
