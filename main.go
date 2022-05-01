package main

import (
	"fmt"

	"github.com/mamcer/tech-test/palindrome"
)

func main() {
	var w string = ""
	fmt.Printf("enter a word to see if it is palindrome: ")
	fmt.Scanln(&w)
	if palindrome.IsPalindromeFor(w) {
		fmt.Printf("%q is palindrome\n", w)
	} else {
		fmt.Printf("%q is not palindrome\n", w)
	}
}
