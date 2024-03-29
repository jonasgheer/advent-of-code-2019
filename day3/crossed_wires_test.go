package main

import (
	"strings"
	"testing"

	"github.com/deckarep/golang-set"
)

func TestVisitedPositions(t *testing.T) {
	w := wire{
		ins{
			direction: "R",
			units:     3,
		},
		ins{
			direction: "D",
			units:     2,
		},
	}
	expectedCords := mapset.NewSetFromSlice([]interface{}{cord{1, 0}, cord{2, 0}, cord{3, 0}, cord{3, -1}, cord{3, -2}})
	actualCords := visitedPositions(w)
	if !expectedCords.Equal(actualCords) {
		t.Errorf("visitedPositions(%v) == %v, expected %v", w, actualCords, expectedCords)
	}
}

func TestStepsToIntersection(t *testing.T) {
	wire := readWires(strings.NewReader("R8,U5,L5,D3"))[0]
	expectedSteps := 20
	intersectCord := cord{3, 3}

	actualSteps := stepsToIntersection(wire, intersectCord)
	if actualSteps != expectedSteps {
		t.Errorf("stepsToIntersection(%v) == %d, expected %d", wire, actualSteps, expectedSteps)
	}
}

func TestCompleteRoute(t *testing.T) {
	examples := map[string]int{
		"R8,U5,L5,D3\nU7,R6,D4,L4": 6,
		"R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83": 159,
	}
	for ws, dist := range examples {
		wires := readWires(strings.NewReader(ws))
		t.Log("wires", wires)

		w1Visited := visitedPositions(wires[0])
		w2Visited := visitedPositions(wires[1])

		t.Log("w1Visited", w1Visited)
		t.Log("w2Visited", w2Visited)

		isect := w1Visited.Intersect(w2Visited)
		t.Log("intersection:", isect)

		isectSlice := isect.ToSlice()
		var dists []int
		for _, c := range isectSlice {
			dists = append(dists, manhattanDistanceFromOrigin(c.(cord)))
		}

		actual := min(dists)
		if actual != dist {
			t.Errorf("min(%v) == %d, expected %d", dists, actual, dist)
		}
	}
}

func TestManhattanDistance(t *testing.T) {
	c := cord{
		x: -3,
		y: 4,
	}
	dist := manhattanDistanceFromOrigin(c)
	expected := 7
	if dist != expected {
		t.Errorf("manhattanDistance(%v) == %d, expected %d", c, dist, expected)
	}
}

func TestMin(t *testing.T) {
	vals := []int{13, 4, 23, 1, 33}
	expected := 1
	actual := min(vals)
	if actual != expected {
		t.Errorf("min(%v) == %d, expected %d", vals, actual, expected)
	}
}
