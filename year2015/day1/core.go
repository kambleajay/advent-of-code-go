package main

import (
	"fmt"
	"io/ioutil"
)

func nextFloor(nextInstruction byte, currentFloor int) int {
	if nextInstruction == '(' {
		return currentFloor + 1
	} else if nextInstruction == ')' {
		return currentFloor - 1
	} else {
		panic(fmt.Sprintf("unexpected char found: %c", nextInstruction))
	}
}

func CalculateFloor(instructions string) int {
	var f int
	for i := 0; i < len(instructions); i++ {
		f = nextFloor(instructions[i], f)
	}
	return f
}

func BasementPosition(instructions string) int {
	var p int
	for i := 0; i < len(instructions); i++ {
		p = nextFloor(instructions[i], p)
		if p == -1 {
			return i + 1
		}
	}
	panic("basement position not found")
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(fmt.Sprintf("could not open file: %v", err))
	}
	fmt.Printf("answer 1: %d\n", CalculateFloor(string(input)))
	fmt.Printf("answer 2: %d\n", BasementPosition(string(input)))
}
