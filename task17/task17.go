package main

import "fmt"

/*
Бинарный поиск

Реализовать алгоритм бинарного поиска встроенными методами языка.
Функция должна принимать отсортированный слайс и искомый элемент,
возвращать индекс элемента или -1, если элемент не найден.
*/

func main() {

	nums := []int{1, 10, 11, 22, 24, 24, 35, 90}
	// искомый элемент
	k := 11
	res := bSort(nums, k)
	fmt.Println(res)
}

func bSort(nums []int, k int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		midIndex := (left + right) / 2
		midVal := nums[midIndex]
		if midVal == k {
			return midIndex
		} else if midVal > k {
			right = midIndex - 1
		} else {
			left = midIndex + 1
		}

	}

	return -1
}
