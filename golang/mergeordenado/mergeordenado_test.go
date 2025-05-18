package mergeordenado

import (
	"reflect"
	"testing"
)

func TestMergeOrdenado(t *testing.T) {
	tests := []struct {
		name   string
		array1 []int
		array2 []int
		want   []int
	}{
		{
			name:   "Both arrays empty",
			array1: []int{},
			array2: []int{},
			want:   []int{},
		},
		{
			name:   "First array empty",
			array1: []int{},
			array2: []int{1, 2, 3},
			want:   []int{1, 2, 3},
		},
		{
			name:   "Second array empty",
			array1: []int{1, 2, 3},
			array2: []int{},
			want:   []int{1, 2, 3},
		},
		{
			name:   "No overlap",
			array1: []int{1, 2, 3},
			array2: []int{4, 5, 6},
			want:   []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:   "Interleaved arrays",
			array1: []int{1, 3, 5},
			array2: []int{2, 4, 6},
			want:   []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:   "Arrays with duplicates",
			array1: []int{1, 2, 2, 3},
			array2: []int{2, 2, 4},
			want:   []int{1, 2, 2, 2, 2, 3, 4},
		},
		{
			name:   "Arrays with negative numbers",
			array1: []int{-3, -1, 2},
			array2: []int{-2, 0, 3},
			want:   []int{-3, -2, -1, 0, 2, 3},
		},
		{
			name:   "Arrays of different lengths",
			array1: []int{1, 4, 5, 8},
			array2: []int{2, 3, 6},
			want:   []int{1, 2, 3, 4, 5, 6, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeOrdenado(tt.array1, tt.array2)
			if (len(got) != 0 && len(tt.want) != 0) && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeOrdenado(%v, %v) = %v; want %v", tt.array1, tt.array2, got, tt.want)
			}
		})
	}
}
