package main

import (
	"fmt"
	"os"
)

func main() {
	// get word string input
	if len(os.Args) > 1 {
		letters := os.Args[1]
		if len(letters) >= 3 {
			fmt.Println(letters)
		} else {
			fmt.Println("Please supply at least 3 letters for anagram generation.")
		}
	}
}
