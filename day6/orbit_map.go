// planets are stored in a map[string]string where
// the key is a planet orbiting the value
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("planets.csv")
	if err != nil {
		panic(err)
	}
	planets := readCsvInput(file)
	numDir := numDirectOrbits(planets)
	numIndir := numIndirectOrbits(planets)

	fmt.Println("total number of direct and indirect orbits:", numDir + numIndir)

	yourIndirect := indirectOrbits("YOU", planets)
	santaIndirect := indirectOrbits("SAN", planets)

	transfers := minNumberOfTransfers(yourIndirect, santaIndirect)

	fmt.Println("minimum number of orbital transfers from you to santa:", transfers)
}

func numDirectOrbits(planets map[string]string) int {
	return len(planets)
}

func numIndirectOrbits(planets map[string]string) int {
	var indirectOrbits int
	for _, orbits := range planets {
		for {
			// if key does not exist it means that planet does not orbit another
			if o, ok := planets[orbits]; ok {
				orbits = o
				indirectOrbits++
				continue
			}
			break
		}
	}
	return indirectOrbits
}

func indirectOrbits(planet string, planets map[string]string) []string {
	var indirectOrbits []string
	orbits := planets[planet]
	indirectOrbits = append(indirectOrbits, orbits)
	for {
		if o, ok:= planets[orbits]; ok {
			orbits = o
			indirectOrbits = append(indirectOrbits, o)
			continue
		}
		break
	}	
	return indirectOrbits
}

func minNumberOfTransfers(a, b []string) int {
	totalSteps := 0
	var commonPlanet string
	for _, planet := range a {
		if containsPlanet(b, planet) {
			commonPlanet = planet
			break	
		}
		totalSteps++
	}
	for _, planet := range b {
		if planet == commonPlanet {
			return totalSteps
		}
		totalSteps++
	}
	panic("a and b does not have a planet in common")
}

func containsPlanet(planets []string, planet string) bool {
	for _, s := range planets {
		if s == planet {
			return true
		}
	}
	return false
}

func readCsvInput(r io.Reader) map[string]string {
	result := make(map[string]string)
	reader := csv.NewReader(r)
	reader.FieldsPerRecord = -1
	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, l := range lines {
		for _, pair := range l {
			pair := strings.Split(pair, ")")
			result[pair[1]] = pair[0]
		}
	}
	return result
}
