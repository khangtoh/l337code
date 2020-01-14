package main

import (
	"fmt"
)

func main() {

	mat := [][]int{
		{1, 0, 0, 1, 0},
		{0, 0, 1, 0, 1},
		{1, 0, 0, 1, 0},
		{1, 0, 1, 0, 1},
	}
	//fmt.Println(hasSameBits(mat[0], mat[1]))
	//fmt.Println(hasSameBits(mat[0], mat[2]))

	//fmt.Println(isRectangle(mat))

	mat = [][]int{
		{1, 0, 1},
		{0, 1, 0},
		{1, 0, 1},
	}

	//fmt.Println(isRectangle(mat))

	mat = [][]int{
		{1, 0, 0},
		{1, 0, 0},
		{1, 1, 0},
	}

	fmt.Println(isRectangle(mat))
}

func isRectangle(grid [][]int) bool {

	rows := len(grid)
	// for each row, check if there is a similar row
	// with same bit patterm

	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			if j != i {
				if hasAtLeastTwoSameBits(grid[i], grid[j]) {
					return true
				}
			}
		}
	}
	return false
}

// func isRectangle(grid [][]int) bool {
// 	rows := len(grid)
// 	cols := len(grid[0])
// 	var rowInts []int
// 	// scans through each row
// 	for i := 0; i < rows; i++ {
// 		row := grid[i]
// 		rowBits := 0
// 		for j := 0; j < cols; j++ {
// 			rowBits = rowBits<<1 | row[j]
// 		}
// 		rowInts = append(rowInts, rowBits)
// 	}
// 	var hasRect int = 0
// 	for _, rowInt := range rowInts {
// 		hasRect = rowInt | hasRect
// 	}
// 	return hasRect != 0
// }

func hasAtLeastTwoSameBits(a []int, b []int) bool {
	cols := len(a)
	aInt := 0
	bInt := 0
	sameBits := 0
	for j := 0; j < cols; j++ {
		if (a[j] & b[j]) == 1 {
			sameBits++
		}
		aInt = aInt<<1 | a[j]
		bInt = bInt<<1 | b[j]
	}
	return (aInt&bInt) != 0 && sameBits > 1
	//return sameBits > 1
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
