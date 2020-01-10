package main

import "fmt"

func main() {
	var grid = [][]int{{1, 2, 2}, {13, 1, 1}}
	end := FindEndPoint(grid)
	fmt.Println(end.X, end.Y)
}

// Coord stores the X, Y values for a coord
// of a M by N matric array
type Coord struct {
	X int
	Y int
}

// MinPathSum func
func MinPathSum(grid [][]int) int {
	endCoord := FindEndPoint(grid)
	if endCoord.X == 0 && endCoord.Y == 0 {
		return 0
	}
	return 0

}

// FindEndPoint func
func FindEndPoint(grid [][]int) Coord {
	depth := len(grid)
	breadth := len(grid[0])
	if breadth != 0 && depth != 0 {
		return Coord{breadth, depth}
	}
	return Coord{0, 0}
}
