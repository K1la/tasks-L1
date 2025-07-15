package main

import "fmt"

/*
	Удаление элемента слайса

	Удалить i-ый элемент из слайса.
	Продемонстрируйте корректное удаление без утечки памяти.
*/

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	k := 7
	s := deleteK(nums, k)
	fmt.Printf("AFTER\tlen=%v\tcap=%v\ns=%v\n", len(s), cap(s), s)
}

func deleteK(nums []int, k int) []int {
	if k < 0 || k >= len(nums) {
		return nums
	}

	copy(nums[k:], nums[k+1:])
	return nums[:len(nums)-1]
}
