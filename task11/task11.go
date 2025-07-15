package main

import "fmt"

/*
	Реализовать пересечение двух неупорядоченных множеств
	(например, двух слайсов) — т.е. вывести элементы,
	присутствующие и в первом, и во втором.

	Пример:
	A = {1,2,3}
	B = {2,3,4}
	Пересечение = {2,3}
*/

func main() {
	sl1 := []int{1, 2, 3}
	sl2 := []int{2, 3, 4}
	setsSl(sl1, sl2)
}

func setsSl(s1 []int, s2 []int) {
	set1 := make(map[int]struct{})
	for _, v := range s1 {
		set1[v] = struct{}{}
	}

	for _, v := range s2 {
		if _, inSet1 := set1[v]; inSet1 {
			fmt.Printf("%v ", v)
		}
	}

}
