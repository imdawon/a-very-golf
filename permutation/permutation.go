package permutation

import (
	"a-very-golf/anagram"
	"fmt"
)

func GetAnagramCandidates(str string, i, n int) map[string]int {
	// base condition
	if i == n-1 {
		fmt.Println(str)

		return nil
	}

	anagramCandidates := make(map[string]int)

	// process each character of the remaining string
	for j := i; j < n; j++ {
		// swap character at index i with current character
		str = swap(str, i, j)

		_, hashExists := anagramCandidates[str]

		if !hashExists {
			anagramCandidates[anagram.RemoveMultipleAdjacentSpaces(str)] = 1
		} else {
			continue
		}

		// recursion for string [i+1:n]
		GetAnagramCandidates(str, i+1, n)

		// backtrack (restore the string to its original state)
		str = swap(str, i, j)
	}

	return anagramCandidates
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
