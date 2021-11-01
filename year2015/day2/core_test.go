package main

import "testing"

func TestWrappingPaper(t *testing.T) {
	var tests = []struct {
		l, w, h int
		want    int
	}{
		{2, 3, 4, 58},
		{1, 1, 10, 43},
	}
	for _, test := range tests {
		if got := PaperRequired(Dimensions{test.l, test.w, test.h}); got != test.want {
			t.Errorf("PaperRequired(%d, %d, %d) = %d, was %d", test.l, test.w, test.h, test.want, got)
		}
	}
}

func Contains(xs []int, x int) bool {
	for _, nextX := range xs {
		if nextX == x {
			return true
		}
	}
	return false
}

func ContainsAll(s1, s2 []int) bool {
	for _, nextS2 := range s2 {
		if !Contains(s1, nextS2) {
			return false
		}
	}
	return true
}

func TestMin2(t *testing.T) {
	var tests = []struct {
		nums []int
		want []int
	}{
		{[]int{3, 1, 2}, []int{1, 2}},
		{[]int{3, 3, 4}, []int{3, 3}},
		{[]int{1, 10, 100}, []int{1, 10}},
		{[]int{5, 10, 10}, []int{5, 10}},
		{[]int{19, 24, 12}, []int{19, 12}},
	}
	for _, test := range tests {
		if m1, m2 := Min2(test.nums[0], test.nums[1], test.nums[2]); !ContainsAll([]int{m1, m2}, test.want) {
			t.Errorf("Min2(%d, %d, %d) = %d, %d", test.nums[0], test.nums[1], test.nums[2], m1, m2)
		}
	}
}

func TestRibbonRequired(t *testing.T) {
	var tests = []struct {
		input Dimensions
		want int
	} {
		{Dimensions{2, 3, 4}, 34},
		{Dimensions{1, 1, 10}, 14},
		{Dimensions{28, 11, 10}, 3122},
		{Dimensions{2, 7, 21}, 312},
		{Dimensions{24, 6, 6}, 888},
		{Dimensions{7, 9, 9}, 599},
		{Dimensions{4, 3, 8}, 110},
		{Dimensions{23, 23, 3}, 1639},
		{Dimensions{19, 24, 12}, 5534},
	}
	for _, test := range tests {
		if got := RibbonRequired(test.input); got != test.want {
			t.Errorf("RibbonRequired(%v) = %d", test.input, got)
		}
	}
}
