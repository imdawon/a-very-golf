package hashmap

import (
	"bufio"
	"fmt"
	"os"
)

// Read local full dictionary file and return a hashmap with the entries.
func GetValidWords() map[string]int {
	validWords := make(map[string]int)

	fileHandle, err := os.Open("fulldict.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(fileHandle)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		validWords[fileScanner.Text()] = 1
	}

	return validWords
}
