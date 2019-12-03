package main

import "testing"

func TestIntcodeProgram(t *testing.T) {
	initialStates := [][]int{
		[]int{1, 0, 0, 0, 99},
		[]int{2, 3, 0, 3, 99},
		[]int{2, 4, 4, 5, 99, 0},
		[]int{1, 1, 1, 4, 99, 5, 6, 0, 99},
	}
	finalStates := [][]int{
		[]int{2, 0, 0, 0, 99},
		[]int{2, 3, 0, 6, 99},
		[]int{2, 4, 4, 5, 99, 9801},
		[]int{30, 1, 1, 4, 2, 5, 6, 0, 99},
	}

	for i := range initialStates {
		actual := intcodeProgram(initialStates[i])
		expected := finalStates[i][0]

		if actual != expected {
			t.Errorf("intcodeProgram(%v) == %d, wanted %d", initialStates[i], actual, expected)
		}
	}
}
