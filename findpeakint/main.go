package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(findPeakElement(nums))

	nums = []int{1, 2, 1, 3, 5, 6, 4}
	fmt.Println(findPeakElement(nums))

}

func findPeakElement(nums []int) int {
	var peak int

	if len(nums) == 1 {
		peak = 0
	} else if len(nums) < 3 {
		if nums[0] > nums[1] {
			peak = 0
		} else {
			peak = 1
		}
	} else {
		for i := 1; i < len(nums); i++ {
			if i+1 >= len(nums) {
				break
			}
			if nums[i] > nums[i-1] || nums[i] > nums[i+1] {
				peak = i
			}
		}
	}
	return peak
}
