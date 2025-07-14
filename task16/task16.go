package main

import "fmt"

func main() {
	nums := []int{90, 10, 24, 22, 35, 11, 100, 1, 24}

	fmt.Println("before:\t", nums)
	numse := quickSort(nums)
	fmt.Println("after:\t", numse)
}

func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	pivot := nums[len(nums)/2]

	var less, equal, greater []int

	for _, v := range nums {
		switch {
		case v < pivot:
			less = append(less, v)
		case v == pivot:
			equal = append(equal, v)
		default:
			greater = append(greater, v)
		}
	}

	sortedLess := quickSort(less)
	sortedGreat := quickSort(greater)

	result := make([]int, 0, len(nums))
	result = append(result, sortedLess...)
	result = append(result, equal...)
	result = append(result, sortedGreat...)
	return result
}
