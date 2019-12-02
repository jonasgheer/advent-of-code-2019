package main

import (
	"testing"
)

func TestRequiredFuel(t *testing.T) {
	requiedFuelByMass := map[int]int{
		12:     2,
		14:     2,
		1969:   654,
		100756: 33583,
	}
	for mass, reqFuel := range requiedFuelByMass {
		if requiredFuel(mass) != reqFuel {
			t.Errorf("requiredFuel(%d) == %d, wanted %d", mass, requiredFuel(mass), reqFuel)
		}
	}
}
