package main

import "testing"

// valid password
var pw = password{1, 2, 2, 3, 4, 5}

func TestMeetsAdjacentCriteria(t *testing.T) {
	if pw.meetsAdjacentCriteria() != true {
		t.Errorf("pw.meetsAdjacentCriteria is %t, wanted %t", pw.meetsAdjacentCriteria(), true)
	}
}

func TestMeetsStrictAdjacentCriteria(t *testing.T) {
	validPws := []password{
		password{1, 1, 1, 1, 2, 2},
		password{1, 1, 2, 2, 3, 3},
	}
	invalidPws := []password{
		password{1, 2, 3, 4, 4, 4},
	}
	for _, pw := range validPws {
		if pw.meetsStrictAdjacentCriteria() != true {
			t.Log("valid password:", pw)
			t.Errorf("pw.meetsStrictAdjacentCriteria is %t, wanted %t", pw.meetsStrictAdjacentCriteria(), true)
		}
	}
	for _, pw := range invalidPws {
		if pw.meetsStrictAdjacentCriteria() == true {
			t.Log("invalid password:", pw)
			t.Errorf("pw.meetsStrictAdjacentCriteria is %t, wanted %t", pw.meetsStrictAdjacentCriteria(), false)
		}
	}
}

func TestMeetsIncreasingCriteria(t *testing.T) {
	if pw.meetsIncreasingCriteria() != true {
		t.Errorf("pw.meetsIncreasingCriteria is %t, wanted %t", pw.meetsIncreasingCriteria(), true)
	}
}

func TestIntToSlice(t *testing.T) {
	n := 135
	expected := []int{1, 3, 5}
	actual := intToPasswordSlice(n)
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Errorf("intToSlice(%d) == %v, wanted %v", n, actual, expected)
		}
	}
	if len(expected) != len(actual) {
		t.Errorf("intToSlice(%d) == %v, wanted %v", n, actual, expected)
	}
}

func TestCountDigits(t *testing.T) {
	n := 1234
	expected := 4
	actual := countDigits(n)
	if expected != actual {
		t.Errorf("countDigits(%d) == %d, wanted %d", n, actual, expected)
	}
}
