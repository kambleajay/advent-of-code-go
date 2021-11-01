package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Dimensions struct {
	l, w, h int
}

func MinOf2(a1, a2 int) int {
	if a1 < a2 {
		return a1
	} else {
		return a2
	}
}

func MinOf3(a1, a2, a3 int) int {
	return MinOf2(MinOf2(a1, a2), a3)
}

func PaperRequired(d Dimensions) int {
	areaOfSide1 := d.l * d.w
	areaOfSide2 := d.w * d.h
	areaOfSide3 := d.h * d.l
	smallestArea := MinOf3(areaOfSide1, areaOfSide2, areaOfSide3)
	surfaceArea := (2 * areaOfSide1) + (2 * areaOfSide2) + (2 * areaOfSide3)
	return surfaceArea + smallestArea
}

func Min2(x, y, z int) (int, int) {
	nums := []int{x, y, z}
	if nums[0] > nums[1] {
		nums[0], nums[1] = nums[1], nums[0]
	}
	if nums[1] > nums[2] {
		nums[1], nums[2] = nums[2], nums[1]
	}
	return nums[0], nums[1]
}

func RibbonRequired(d Dimensions) int {
	min1, min2 := Min2(d.l, d.w, d.h)
	smallestPerimeter := (2 * min1) + (2 * min2)
	volume := d.l * d.w * d.h
	return smallestPerimeter + volume
}

func ParseDimensions(line string) Dimensions {
	parts := strings.Split(line, "x")
	var partsInt []int
	for _, part := range parts {
		d, err := strconv.Atoi(part)
		if err != nil {
			panic(fmt.Sprintf("could not parse: %s, in %s", part, line))
		}
		partsInt = append(partsInt, d)
	}
	return Dimensions{partsInt[0], partsInt[1], partsInt[2]}
}

func ReadDimensions(data string) []Dimensions {
	var dims []Dimensions
	for _, line := range strings.Split(data, "\n") {
		if len(line) > 0 {
			dims = append(dims, ParseDimensions(line))
		}
	}
	return dims
}

func PaperRequiredForAll(dimensions []Dimensions) int {
	var paperNeeded int
	for _, dimension := range dimensions {
		paperNeeded += PaperRequired(dimension)
	}
	return paperNeeded
}

func RibbonRequiredAll(dimensions []Dimensions) int {
	var ribbonNeeded int
	for _, dimension := range dimensions {
		ribbonNeeded += RibbonRequired(dimension)
	}
	return ribbonNeeded
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(fmt.Sprintf("could not open file input.txt: %v", err))
	}
	dimensions := ReadDimensions(string(data))

	fmt.Printf("answer 1: paper needed: %d\n", PaperRequiredForAll(dimensions))
	fmt.Printf("answer 2: ribbon needed: %d\n", RibbonRequiredAll(dimensions))
}
