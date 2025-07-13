package main

import "fmt"

func main() {
	nums := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	ans := make(map[int][]float32)

	for _, n := range nums {
		key := int(n) / 10 * 10
		ans[key] = append(ans[key], n)
	}

	for k, n := range ans {
		fmt.Println(k, n)
	}
}
