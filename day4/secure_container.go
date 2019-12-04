package main

import "fmt"

type password []int

func main() {
	validCount := 0
	strictValidCount := 0
	for i := 108457; i < 562041; i++ {
		pw := intToPasswordSlice(i)
		if pw.meetsAdjacentCriteria() && pw.meetsIncreasingCriteria() {
			validCount++
		}
		if pw.meetsStrictAdjacentCriteria() && pw.meetsIncreasingCriteria() {
			strictValidCount++
		}
	}
	fmt.Println("valid number of passwords:", validCount)
	fmt.Println("valid number of strict passwords:", strictValidCount)
}

func (pw password) meetsAdjacentCriteria() bool {
	for i := 0; i < len(pw)-1; i++ {
		if pw[i] == pw[i+1] {
			return true
		}
	}
	return false
}

func (pw password) meetsStrictAdjacentCriteria() bool {
	matches := []bool{}
	for i := 0; i < len(pw)-1; i++ {
		if pw[i] == pw[i+1] {
			matches = append(matches, true)
		} else {
			matches = append(matches, false)
		}
	}

	for i := 0; i < len(matches); i++ {
		if matches[i] == true {
			//check neighbors
			if (i - 1) >= 0 {
				if matches[i-1] == true {
					continue
				}
			}
			if (i + 1) <= len(matches)-1 {
				if matches[i+1] == true {
					continue
				}
			}
			return true
		}
	}
	return false
}

func (pw password) meetsIncreasingCriteria() bool {
	curr := pw[0]
	for i := 1; i < len(pw); i++ {
		if pw[i] < curr {
			return false
		}
		curr = pw[i]
	}
	return true
}

func countDigits(n int) int {
	count := 0
	for n != 0 {
		n /= 10
		count++
	}
	return count
}

func intToPasswordSlice(n int) password {
	numDigits := countDigits(n)
	s := make(password, numDigits)
	for i := 0; i < numDigits; i++ {
		s[numDigits-i-1] = n % 10
		n /= 10
	}
	return s
}
