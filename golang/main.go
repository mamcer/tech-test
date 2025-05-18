package main

import (
	"fmt"
	"github.com/mamcer/tech-test/mergeordenado"

	// "github.com/mamcer/tech-test/palindrome"
)

func main() {
	// var w string = ""
	// fmt.Printf("Please type a word to see if it is a palindrome: ")
	// fmt.Scanln(&w)
	// if palindrome.IsPalindromeFor(w) {
	// 	fmt.Printf("%q is a palindrome.\n", w)
	// } else {
	// 	fmt.Printf("%q is not a palindrome.\n", w)
	// }

	var a []int = []int{1, 3, 5, 7, 9}
	var b []int = []int{2, 4, 6, 8, 10}
	var r []int = mergeordenado.MergeOrdenado(a, b)
	fmt.Printf("Merged array: %v\n", r)	
}
