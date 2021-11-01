package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Point struct {
	x, y int
}

func housesVisited(instructions string) map[Point]bool {
	housesVisited := make(map[Point]bool)
	currentLocation := Point{0, 0}
	for _, instruction := range instructions {
		if instruction == '^' {
			nextLocation := Point{currentLocation.x, currentLocation.y + 1}
			housesVisited[nextLocation] = true
			currentLocation = nextLocation
		} else if instruction == '>' {
			nextLocation := Point{currentLocation.x + 1, currentLocation.y}
			housesVisited[nextLocation] = true
			currentLocation = nextLocation
		} else if instruction == 'v' {
			nextLocation := Point{currentLocation.x, currentLocation.y - 1}
			housesVisited[nextLocation] = true
			currentLocation = nextLocation
		} else if instruction == '<' {
			nextLocation := Point{currentLocation.x - 1, currentLocation.y}
			housesVisited[nextLocation] = true
			currentLocation = nextLocation
		}
	}
	return housesVisited
}

func answer1(instructions string) int {
	houses := housesVisited(instructions)
	houses[Point{0, 0}] = true
	return len(houses)
}

func answer2(instructions string) int {
	var santaInstructions, roboSantaInstructions strings.Builder
	for i, instruction := range instructions {
		if i%2 == 0 {
			santaInstructions.WriteRune(instruction)
		} else {
			roboSantaInstructions.WriteRune(instruction)
		}
	}
	santaHouses := housesVisited(santaInstructions.String())
	roboSantaHouses := housesVisited(roboSantaInstructions.String())
	for k, v := range roboSantaHouses {
		santaHouses[k] = v
	}
	santaHouses[Point{0,0}] = true
	return len(santaHouses)
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(fmt.Sprintf("could not open file input.txt: %v", err))
	}
	fmt.Printf("answer 1: %d\n", answer1(string(data)))
	fmt.Printf("answer 2: %d\n", answer2(string(data)))
}
