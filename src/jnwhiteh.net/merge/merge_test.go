package merge

import "math/rand"
import "sort"
import "testing"

type mergeTest struct {
	input []int
	result []int
}

var mergeTests = []mergeTest{
	{input: []int{5, 4, 3, 2, 1}, result: []int{1, 2, 3, 4, 5},},
	{input: []int{3,2,1}, result: []int{1,2,3},},
}

func Equal(left []int, right []int) bool {
	if len(left) != len(right) {
		return false
	}
	for i := 0; i < len(left); i++ {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}

func TestMerge(T *testing.T) {
	for _, test := range mergeTests {
		orig := make([]int, len(test.input))
		copy(orig, test.input)
		result := make(chan []int)
		go Sort(test.input, result)
		sorted := <-result
		if !Equal(sorted, test.result) {
			T.Errorf("When testing %v: expected %v, got %v", orig, test.result, sorted)
		}
	}
}

func TestRandoms(T *testing.T) {
	// Run a sort on 100 random integers
	for i:= 0; i < 100; i++ {
		length := 100

		var data []int
		for j := 0; j < length; j++ {
			data = append(data, rand.Intn(100))
		}

		rchan := make(chan []int)
		go Sort(data, rchan)

		result := <-rchan

		check := make([]int, length)
		copy(check, data)
		sort.Ints(check)

		if !Equal(result, check) {
			T.Errorf("When testing %v: expected %v, got %v", data, check, result)
		}
	}
}
