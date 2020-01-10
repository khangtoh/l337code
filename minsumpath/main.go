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
func MinPathSum(grid [][]int, currentCoord Coord, cost int) int {
	endCoord := FindEndPoint(grid)
	var coord1 Coord
	var coord2 Coord
	var cost1 int
	var cost2 int

	if currentCoord.X > endCoord.X || currentCoord.Y > endCoord.Y {
		return 0
	}

	if endCoord.X == 0 && endCoord.Y == 0 {
		return grid[0][0]
	}

	coord1 = Coord{currentCoord.X + 1, currentCoord.Y}
	coord2 = Coord{currentCoord.X, currentCoord.Y + 1}

	// if already at endpoint
	if currentCoord == endCoord {
		return cost
	}

	// if no more left moves
	if currentCoord.X == endCoord.X {
		return MinPathSum(grid, coord2, cost+grid[coord2.Y][coord2.X])
	}
	// if no more down moves
	if currentCoord.Y == endCoord.Y {
		return MinPathSum(grid, coord1, cost+grid[coord1.Y][coord1.X])
	}

	// try both moves and return the lower cost
	cost1 = MinPathSum(grid, coord1, cost+grid[coord1.Y][coord1.X])
	cost2 = MinPathSum(grid, coord2, cost+grid[coord2.Y][coord2.X])

	if cost2 < cost1 {
		return cost2
	}
	return cost1

	// if currentCoord.X+1 == endCoord.X {
	// 	coord2 = Coord{currentCoord.X, currentCoord.Y + 1}
	// 	return MinPathSum(grid, coord2, cost+grid[coord2.Y][coord2.X])
	// }
	// if currentCoord.Y+1 == endCoord.Y {
	// 	coord1 = Coord{currentCoord.X + 1, currentCoord.Y}
	// 	return MinPathSum(grid, coord1, cost+grid[coord1.Y][coord1.X])
	// }

	//return 0

}

// FindEndPoint func
func FindEndPoint(grid [][]int) Coord {
	depth := len(grid) - 1
	breadth := len(grid[0]) - 1
	if breadth != 0 && depth != 0 {
		return Coord{breadth, depth}
	}
	return Coord{0, 0}
}
