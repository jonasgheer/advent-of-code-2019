package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("intcodes.csv")
	if err != nil {
		panic(err)
	}
	intcodes := readCsvInput(file)

	diagnosticProgram(intcodes, 1)
}

type instruct struct {
	opcode     int
	paramOne   int
	paramTwo   int
	paramThree int
}

// converts a number instruction to corresponding array
// ex: 1002 -> [0, 1, 0, 2]
func parseInstruction(ins int) [4]int {
	result := [4]int{}
	// handle two rightmost digits
	a := ins % 10
	ins /= 10
	b := ins % 10
	ins /= 10
	result[3] = (b * 10) + a

	numDigits := countDigits(ins)
	nextPos := 2
	for i := 0; i < numDigits; i++ {
		result[nextPos] = ins % 10
		ins /= 10
		nextPos--
	}
	return result
}

func countDigits(n int) int {
	count := 0
	for n != 0 {
		n /= 10
		count++
	}
	return count
}

func diagnosticProgram(intcodes []int, inputInstruction int) {
	for i := 0; i < len(intcodes); {
		ins := parseInstruction(intcodes[i])
		param3, param2, param1, opcode := ins[0], ins[1], ins[2], ins[3]
		switch opcode {
		case 1: // add
			var a, b int
			if param1 == 0 {
				a = intcodes[intcodes[i+1]]
			} else {
				a = intcodes[i+1]
			}
			if param2 == 0 {
				b = intcodes[intcodes[i+2]]
			} else {
				b = intcodes[i+2]
			}
			if param3 == 0 {
				intcodes[intcodes[i+3]] = a + b
			} else {
				intcodes[i+3] = a + b
			}
			i += 4
		case 2: // multiply
			var a, b int
			if param1 == 0 {
				a = intcodes[intcodes[i+1]]
			} else {
				a = intcodes[i+1]
			}
			if param2 == 0 {
				b = intcodes[intcodes[i+2]]
			} else {
				b = intcodes[i+2]
			}
			if param3 == 0 {
				intcodes[intcodes[i+3]] = a * b
			} else {
				intcodes[i+3] = a * b
			}
			i += 4
		case 3: // input
			if param1 == 0 {
				intcodes[intcodes[i+1]] = inputInstruction
			} else {
				intcodes[i+1] = inputInstruction
			}
			i += 2
		case 4: // output
			fmt.Println(intcodes[intcodes[i+1]])
			i += 2
		case 99: // halt
			return
		default:
			log.Fatalf("invalid intcode %d", intcodes[i])
		}
	}
	panic("intcode 99 not found")
}

// readCsvInput reads csv format of numbers and returns a single []int
func readCsvInput(r io.Reader) []int {
	var result []int
	reader := csv.NewReader(r)
	reader.FieldsPerRecord = -1
	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, l := range lines {
		for _, number := range l {
			n, err := strconv.Atoi(number)
			if err != nil {
				panic(fmt.Sprintf("could not convert %s to int", number))
			}
			result = append(result, n)
		}
	}
	return result
}
