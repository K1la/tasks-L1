package main

import "fmt"

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
