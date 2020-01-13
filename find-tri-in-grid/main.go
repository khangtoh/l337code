package main

import (
	"fmt"
	"strconv"
)

func main() {

	mat := [][]int{
		{1, 0, 0, 1, 0},
		{0, 0, 1, 0, 1},
		{0, 0, 0, 1, 0},
		{1, 0, 1, 0, 1},
	}

	fmt.Println(isRectangle(mat) == true)
}

func isRectangle(grid [][]int) bool {
	rows := len(grid)
	cols := len(grid[0])
	var rowInts []int
	// scans through each row
	for i := 0; i < rows; i++ {
		row := grid[i]
		rowBits := 0
		for j := 0; j < cols; j++ {
			rowBits = rowBits | row[j]<<j
			// if row[j] == 1 {
			// 	rowBits = rowBits | row[j]<<j
			// } else {
			// 	rowBits = rowBits << 1
			// }

		}
		rowInts = append(rowInts, reverseBits(rowBits))
		fmt.Println("Row ", i, strconv.FormatInt(int64(rowBits), 2))
	}
	var hasRect int = 0
	for _, rowInt := range rowInts {
		hasRect = rowInt & hasRect
	}
	return hasRect != 0
}
func reverseBits(b byte) byte {
	var d byte
	for i := 0; i < 8; i++ {
		d <<= 1
		d |= b & 1
		b >>= 1
	}
	return d
}
