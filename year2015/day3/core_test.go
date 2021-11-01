package main

import "testing"

func TestAnswer1(t *testing.T) {
	var tests = []struct {
		instructions string
		want         int
	}{
		{">", 2},
		{"^>v<", 4},
		{"^v^v^v^v^v", 2},
		{"^>^>^>^>v<", 10},
		{">^^v^<>v<<", 7},
		{"^>vv<^", 6},
		{"^^^>>>vvv<<<", 12},
	}
	for _, test := range tests {
		if got := answer1(test.instructions); got != test.want {
			t.Errorf("answer1(%s) = %d, expected %d", test.instructions, got, test.want)
		}
	}
}

func TestAnswer2(t *testing.T) {
	var tests = []struct {
		instructions string
		want         int
	}{
		{">", 2},
		{"^v", 3},
		{"^^", 2},
		{"><", 3},
		{">>", 2},
		{"^>>^", 4},
		{"^v><", 5},
		{"^>v<", 3},
		{"^^>>vv", 4},
		{"^v><v^", 7},
		{"^v^v^v^v^v", 11},
		{"^>^>^>^>v<", 9},
		{">^^v^<>v<<", 9},
		{"^>vv<^", 5},
		{"^^^>>>vvv<<<", 8},
		{"^^>>vv", 4},
		{"^^>>vv<<", 4},
	}
	for _, test := range tests {
		if got := answer2(test.instructions); got != test.want {
			t.Errorf("answer2(%s) = %d, expected %d", test.instructions, got, test.want)
		}
	}
}
