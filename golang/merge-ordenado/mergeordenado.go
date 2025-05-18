package mergeordenado;

func MergeOrdenado(int[] array1, array2) int[] {
	var k int = 0
	var r []int
	var i = 0
	var j = 0
	for ; k < len(array1)+len(array2); k++ {
		if i >= len(array1) {
			r = append(r, array2[j])
			j++
			continue
		}
		if j >= len(array2) {
			r = append(r, array1[i])
			i++
			continue
		}
		if array1[i] < array2[j] {
			r = append(r, array1[i])
			i++
			continue
		} else {
			r = append(r, array2[j])
			j++
		}
	}
	
	return r
}