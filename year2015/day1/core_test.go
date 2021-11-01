package main

import "testing"

func TestCalculateFloor(t *testing.T) {
	var tests = []struct {
		input string
		want int
	} {
		{"(())", 0},
		{"()()", 0},
		{"(((", 3},
		{"(()(()(", 3},
		{"))(((((", 3},
		{"())", -1},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
	}
	for _, test := range tests {
		if got := CalculateFloor(test.input); got != test.want {
			t.Errorf("CalculateFloor(%s) = %d, should be: %d", test.input, got, test.want)
		}
	}
}

func TestBasementPosition(t *testing.T) {
	var tests = []struct {
		input string
		want int
	} {
		{")", 1},
		{"()())", 5},
	}
	for _, test := range tests {
		if got := BasementPosition(test.input); got != test.want {
			t.Errorf("BasementPosition(%s) == %d, should be %d", test.input, got, test.want)
		}
	}
}
