package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	originalIntcodes := readIntcodes()
	intcodes := make([]int, len(originalIntcodes))

	copy(intcodes, originalIntcodes)
	intcodes[1] = 12
	intcodes[2] = 2

	fmt.Println("Position 0 value:", intcodeProgram(intcodes))

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			copy(intcodes, originalIntcodes)
			intcodes[1] = noun
			intcodes[2] = verb
			if intcodeProgram(intcodes) == 19690720 {
				fmt.Printf("100 * %d + %d = %d\n", noun, verb, 100*noun+verb)
				return
			}
		}
	}
}

func intcodeProgram(intcodes []int) int {
	for i := 0; i < len(intcodes); i += 4 {
		switch intcodes[i] {
		case 1: // add
			intcodes[intcodes[i+3]] = intcodes[intcodes[i+1]] + intcodes[intcodes[i+2]]
		case 2: // multiply
			intcodes[intcodes[i+3]] = intcodes[intcodes[i+1]] * intcodes[intcodes[i+2]]
		case 99: // halt
			return intcodes[0]
		default:
			log.Fatalf("invalid intcode %d", intcodes[i])
		}
	}
	panic("intcode 99 not found")
}

func readIntcodes() []int {
	file, err := os.Open("intcodes.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(file)
	temp, err := reader.Read()
	var intcodes = []int{}
	if err != nil {
		log.Fatal(err)
	}
	for _, i := range temp {
		intcode, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		intcodes = append(intcodes, intcode)
	}
	return intcodes
}
