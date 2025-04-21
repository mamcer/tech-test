package main

import "fmt"

func maximumDifference(a []int32) int32 {
	var md int32 = -1
	var i, j int
	for i = 0; i < len(a); i++ {
		for j = i + 1; j < len(a); j++ {
			if a[i] < a[j] {
				var cd = a[j] - a[i]
				if cd > md {
					md = cd
				}
			}
		}
	}
	return md
}

func main() {
	var md = maximumDifference([]int32{15, 3, 6, 10})
	fmt.Printf("maximum difference: %d\n", md)
}
