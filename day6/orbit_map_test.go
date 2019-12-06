package main

import (
	"strings"
	"testing"
)

var exampleInput = `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`

var exampleSantaInput = `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN`

func TestReadCsvInput(t *testing.T) {
	planets := readCsvInput(strings.NewReader(exampleInput))

	if planets["D"] != "C" {
		t.Error("D should orgit C")
	}
	if len(planets) != 11 {
		t.Error("map should contain 11 planets")
	}
}

func TestNumDirectOrbits(t *testing.T) {
	planets := readCsvInput(strings.NewReader(exampleInput))

	if numDirectOrbits(planets) != 11 {
		t.Errorf("number of direct orbits should equal %d, go %d", 11, numDirectOrbits(planets))
	}
}

func TestIndirectOrbits(t *testing.T) {
	planets := readCsvInput(strings.NewReader(exampleSantaInput))

	expected := []string{"I", "D", "C", "B", "COM"}
	actual := indirectOrbits("SAN", planets)
	if len(expected) != len(actual) {
		t.Errorf("indirectOrbits(%v) == %v, wanted %v", planets, actual, expected)
	}
	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("indirectOrbits(%v) == %v, wanted %v", planets, actual, expected)
			break
		}
	}
}

func TestYouOrbitsSanta(t *testing.T) {
	planets := readCsvInput(strings.NewReader(exampleSantaInput))
	yourIndirect := indirectOrbits("YOU", planets)
	santaIndirect := indirectOrbits("SAN", planets)

	expected := 4
	actual := minNumberOfTransfers(yourIndirect, santaIndirect)

	if actual != expected {
		t.Errorf("minNumberOfTransfers(%v, %v) == %d, wanted %d", yourIndirect, santaIndirect, actual, expected)
	}
}

func TestCompleteExample(t *testing.T) {
	planets := readCsvInput(strings.NewReader(exampleInput))
	numDir := numDirectOrbits(planets)
	numIndir := numIndirectOrbits(planets)

	actual := numDir + numIndir
	expected := 42
	if actual != expected {
		t.Errorf("number of indirect and direct orbits are %d, wanted %d", actual, expected)
	}
}

func TestNumIndirectOrbits(t *testing.T) {
	planets := map[string]string{
		"D": "C",
		"C": "B",
		"B": "COM",
	}
	expected := 3
	actual := numIndirectOrbits(planets)
	if actual != expected {
		t.Errorf("numIndirectOrbits(%v) == %d, wanted %d", planets, actual, expected)
	}
}
