package main

import (
	"fmt"

	"github.com/mamcer/tech-test/palindrome"
)

func main() {
	var w string = ""
	fmt.Printf("Please type a word to see if it is a palindrome: ")
	fmt.Scanln(&w)
	if palindrome.IsPalindromeFor(w) {
		fmt.Printf("%q is a palindrome.\n", w)
	} else {
		fmt.Printf("%q is not a palindrome.\n", w)
	}
}
