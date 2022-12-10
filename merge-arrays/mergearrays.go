package mergearrays

func MergeArrays(a []int, b []int) []int {
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
