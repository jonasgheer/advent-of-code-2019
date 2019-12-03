package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/deckarep/golang-set"
)

type wire []ins

type ins struct {
	direction string
	units     int
}

type cord struct {
	x int
	y int
}

func main() {
	var wires []wire
	file, err := os.Open("wires.csv")
	if err != nil {
		log.Fatal(err)
	}

	wires = readWires(file)

	wireOneVisited := visitedPositions(wires[0])
	wireTwoVisited := visitedPositions(wires[1])

	w1 := mapset.NewSet()
	w2 := mapset.NewSet()

	for _, c := range wireOneVisited {
		w1.Add(c)
	}
	for _, c := range wireTwoVisited {
		w2.Add(c)
	}
	isect := w1.Intersect(w2)
	fmt.Println(isect.Cardinality())

	isectSlice := isect.ToSlice()
	var dists []int
	for _, c := range isectSlice {
		dists = append(dists, manhattanDistance(c.(cord)))
	}

	fmt.Println(min(dists))

}

func min(values []int) int {
	m := values[0]
	for _, n := range values {
		if n < m {
			m = n
		}
	}
	return m
}

// from origin 0,0
func manhattanDistance(c cord) int {
	x, y := c.x, c.y
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	return x + y
}

// bruteforce O(n^2)
func intersections(a, b []cord) []cord {
	var match []cord
	for _, c1 := range a {
		for _, c2 := range b {
			if c1 == c2 {
				match = append(match, c1)
			}
		}
	}
	return match
}

func visitedPositions(w wire) []cord {
	var visited []cord
	var currX, currY int

	for _, ins := range w {
		switch ins.direction {
		case "R":
			for i := 0; i < ins.units; i++ {
				currX++
				visited = append(visited, cord{
					x: currX,
					y: currY,
				})
			}
		case "U":
			for i := 0; i < ins.units; i++ {
				currY++
				visited = append(visited, cord{
					x: currX,
					y: currY,
				})
			}
		case "L":
			for i := 0; i < ins.units; i++ {
				currX--
				visited = append(visited, cord{
					x: currX,
					y: currY,
				})
			}
		case "D":
			for i := 0; i < ins.units; i++ {
				currY--
				visited = append(visited, cord{
					x: currX,
					y: currY,
				})
			}
		default:
			log.Fatalf("invalid direction %s", ins.direction)
		}
	}
	return visited
}

func readWires(r io.Reader) []wire {
	var wires []wire
	reader := csv.NewReader(r)
	reader.FieldsPerRecord = -1
	ws, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var currentWire wire
	for _, w := range ws {
		currentWire = nil
		for _, instruction := range w {
			direction := string(instruction[0])
			units, err := strconv.Atoi(instruction[1:])
			if err != nil {
				log.Fatal(err)
			}
			currentWire = append(currentWire, ins{
				direction: direction,
				units:     units,
			})
		}
		wires = append(wires, currentWire)
	}
	return wires
}
