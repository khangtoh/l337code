package main

import "testing"

var testCases = []struct {
	grid [][]int
}{
	{
		[][]int{{1, 2}, {1, 1}},
	},
	{
		[][]int{{1, 2, 1}, {1, 1, 1}},
	},
	{
		[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
	},
	{
		[][]int{{1, 1, 1}, {1, 3, 2}, {1, 1, 5}},
	},
}

func Test(t *testing.T) {
	tc1 := testCases[0]
	cell := FindEndPoint(tc1.grid)
	if cell.X != 1 && cell.Y != 1 {
		t.Fatalf("endPoint cell is not {1,1}")
	}
	tc2 := testCases[1]
	cell = FindEndPoint(tc2.grid)
	if cell.X != 2 && cell.Y != 1 {
		t.Fatalf("endPoint cell is not {2,1}")
	}

	cost := MinPathSum(tc1.grid, Coord{0, 0}, tc1.grid[0][0])
	if cost != 3 {
		t.Fatalf("tc2 minPath sum %d is not 3", cost)
	}

	tc3 := testCases[2]
	if cost = MinPathSum(tc3.grid, Coord{0, 0}, tc3.grid[0][0]); cost != 7 {
		t.Fatalf("tc2 minPath sum %d is not 7", cost)
	}

	tc4 := testCases[3]
	if cost = MinPathSum(tc4.grid, Coord{0, 0}, tc4.grid[0][0]); cost != 9 {
		t.Fatalf("tc2 minPath sum %d is not 9", cost)
	}
}
