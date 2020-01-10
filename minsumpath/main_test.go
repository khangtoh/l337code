package main

import "testing"

var testCases = []struct {
	grid [][]int
}{
	{
		[][]int{
			{1, 2}, {1, 1},
			{1, 2, 1}, {1, 1, 1},
		},
	},
}

func Test(t *testing.T) {
	tc := testCases[0]
	cell := FindEndPoint(tc.grid)
	if cell.X != 2 && cell.Y != 2 {
		t.Fatalf("Winner is not X and should be X")
	}
	tc = testCases[1]
	cell = FindEndPoint(tc.grid)
	if cell.X != 3 && cell.Y != 3 {
		t.Fatalf("Winner is not X and should be X")
	}

}
