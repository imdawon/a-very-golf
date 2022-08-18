package permutation

import (
	"a-very-golf/anagram"
	"a-very-golf/hashmap"
	"strings"
)

var Iterations = 0

func FindAnagramCandidates(curPermutation string, i int, n int) {
	// Base case.
	if i == n-1 {
		return
	}
	Iterations++

	// Process each character of the remaining string.
	for j := i; j < n; j++ {
		// Swap char at index i with current char.
		curPermutation = swap(curPermutation, i, j)

		_, hashExists := hashmap.SeenAnagrams[curPermutation]

		if !hashExists {
			// Split on spaces to invidually check if each word is a valid English word.
			permutationWords := strings.Split(curPermutation, " ")
			if isValidPermutation(permutationWords, hashmap.ValidWords) {
				// Add valid phrase to our candidates hashmap.
				hashmap.AnagramCandidates[anagram.RemoveMultipleAdjacentSpaces(curPermutation)] = 1

				hashmap.SeenAnagrams[curPermutation] = 1
			}
		} else {
			return
		}

		FindAnagramCandidates(curPermutation, i+1, n)

		// Restore the string to its original state.
		curPermutation = swap(curPermutation, i, j)
	}
}

func swap(s string, i, j int) string {
	var result []byte
	for k := 0; k < len(s); k++ {
		if k == i {
			result = append(result, s[j])
		} else if k == j {
			result = append(result, s[i])
		} else {
			result = append(result, s[k])
		}
	}
	return string(result)
}

func isValidPermutation(permutationWords []string, validWords map[string]int) bool {
	for _, word := range permutationWords {
		_, isValidWord := validWords[word]
		if !isValidWord && word != "" {
			return false
		}
	}
	return true
}
