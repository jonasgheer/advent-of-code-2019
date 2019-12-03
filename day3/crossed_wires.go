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

	isect := wireOneVisited.Intersect(wireTwoVisited)

	isectSlice := isect.ToSlice()
	var dists []int
	for _, c := range isectSlice {
		dists = append(dists, manhattanDistanceFromOrigin(c.(cord)))
	}

	fmt.Println("distance from central port to closest intersection:", min(dists))

	w1steps := stepsToIntersection(wires[0], isectSlice[0].(cord))
	w2steps := stepsToIntersection(wires[1], isectSlice[0].(cord))
	minSteps := w1steps + w2steps
	for _, c := range isectSlice {
		w1steps := stepsToIntersection(wires[0], c.(cord))
		w2steps := stepsToIntersection(wires[1], c.(cord))
		steps := w1steps + w2steps
		if steps < minSteps {
			minSteps = steps
		}
	}
	fmt.Println("fewest combined steps to reach intersection:", minSteps)
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

func manhattanDistanceFromOrigin(c cord) int {
	x, y := c.x, c.y
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	return x + y
}

func stepsToIntersection(w wire, c cord) int {
	var steps int
	var currX, currY int

	for _, ins := range w {
		switch ins.direction {
		case "R":
			for i := 0; i < ins.units; i++ {
				currX++
				steps++
				if currX == c.x && currY == c.y {
					return steps
				}
			}
		case "U":
			for i := 0; i < ins.units; i++ {
				currY++
				steps++
				if currX == c.x && currY == c.y {
					return steps
				}
			}
		case "L":
			for i := 0; i < ins.units; i++ {
				currX--
				steps++
				if currX == c.x && currY == c.y {
					return steps
				}
			}
		case "D":
			for i := 0; i < ins.units; i++ {
				currY--
				steps++
				if currX == c.x && currY == c.y {
					return steps
				}
			}
		default:
			log.Fatalf("invalid direction %s", ins.direction)
		}
	}
	panic("path did not reach intersection")
}

func visitedPositions(w wire) mapset.Set {
	visited := mapset.NewSet()
	var currX, currY int

	for _, ins := range w {
		switch ins.direction {
		case "R":
			for i := 0; i < ins.units; i++ {
				currX++
				visited.Add(cord{
					x: currX,
					y: currY,
				})
			}
		case "U":
			for i := 0; i < ins.units; i++ {
				currY++
				visited.Add(cord{
					x: currX,
					y: currY,
				})
			}
		case "L":
			for i := 0; i < ins.units; i++ {
				currX--
				visited.Add(cord{
					x: currX,
					y: currY,
				})
			}
		case "D":
			for i := 0; i < ins.units; i++ {
				currY--
				visited.Add(cord{
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
