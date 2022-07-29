package main

import "fmt"

func mergeArrays(a []int, b []int) []int {
	var k int = 0
	var r []int
	var i = 0
	var j = 0
	for ; k < len(a)+len(b); k++ {
		if i >= len(a) {
			r = append(r, b[j])
			j++
			continue
		}
		if j >= len(b) {
			r = append(r, a[i])
			i++
			continue
		}
		if a[i] < b[j] {
			r = append(r, a[i])
			i++
			continue
		} else {
			r = append(r, b[j])
			j++
		}
	}
	return r
}

func main() {
	var a = []int{1, 4, 7, 9}
	var b = []int{2, 3, 8, 11, 12}
	var c = mergeArrays(a, b)
	fmt.Printf("result: %v\n", c)
}
