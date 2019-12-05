package main

import (
	"strings"
	"testing"
)

func TestReadCsvInput(t *testing.T) {
	validInput := []string{
		"23,35,12\n45,2,5,22",
		"2,55,2,1,5,6",
	}
	expectedOutput := [][]int{
		[]int{23, 35, 12, 45, 2, 5, 22},
		[]int{2, 55, 2, 1, 5, 6},
	}

	for i := 0; i < len(validInput); i++ {
		expected := expectedOutput[i]
		actual := readCsvInput(strings.NewReader(validInput[i]))

		if len(expected) != len(actual) {
			t.Errorf("readCsvInput(%s) == %v, wanted %v", validInput[i], actual, expected)
		}
		for j := 0; j < len(expected); j++ {
			if expected[j] != actual[j] {
				t.Errorf("readCsvInput(%s) == %v, wanted %v", validInput[i], actual, expected)
			}
		}
	}
}

func TestParseInstruction(t *testing.T) {
	inputs := []int{
		1002,
		10199,
	}
	outputs := [][4]int{
		[4]int{0, 1, 0, 2},
		[4]int{1, 0, 1, 99},
	}
	for i := 0; i < len(inputs); i++ {
		actual := parseInstruction(inputs[i])
		expected := outputs[i]
		if len(actual) != len(expected) {
			t.Errorf("parseInstruction(%v) == %v, wanted %v", inputs[i], actual, expected)
		}
		for j := 0; j < len(actual); j++ {
			if actual[j] != expected[j] {
				t.Errorf("parseInstruction(%v) == %v, wanted %v", inputs[i], actual, expected)
				break
			}
		}
	}
}

// test that the resulting slice is correct
/*
func TestIntCodeProgram(t *testing.T) {
	validInput := []int{1002, 4, 3, 4, 33}

	expected := []int{1002, 4, 3, 4, 99}
	actual := intcodeProgram(validInput)

	if len(actual) != len(expected) {
		t.Errorf("intcodeProgram(%v) == %v, wanted %v", validInput, actual, expected)
	}
	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("intcodeProgram(%v) == %v, wanted %v", validInput, actual, expected)
		}
	}
}
*/
