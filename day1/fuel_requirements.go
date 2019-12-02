package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("modules.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	totalFuelRequiredPartOne := 0
	totalFuelRequiredPartTwo := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		fuel := requiredFuel(mass)
		totalFuelRequiredPartOne += fuel

		for fuel > 0 {
			totalFuelRequiredPartTwo += fuel
			fuel = requiredFuel(fuel)
		}
	}
	fmt.Printf("required fuel for Part One: %d\n", totalFuelRequiredPartOne)
	fmt.Printf("required fuel for Part Two: %d\n", totalFuelRequiredPartTwo)
}

func requiredFuel(mass int) int {
	return mass/3 - 2
}
