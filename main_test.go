package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{1, 2, 3, 4, 5}, 9, []int{3, 4}},
	}

	for _, tt := range tests {
		result := twoSum(tt.nums, tt.target)

		// fmt.Printf("result: %v, want: %v\n", result, tt.want)

		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("result: %v, want: %v", result, tt.want)
		}
	}
}
