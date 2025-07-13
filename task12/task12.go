package main

import "fmt"

func main() {
	sl := []string{"cat", "cat", "dog", "cat", "tree"}
	ans := setsSl(sl)
	fmt.Println(ans)
}

func setsSl(s []string) (ans []string) {
	m := make(map[string]struct{})

	for _, v := range s {
		if _, inM := m[v]; !inM {
			m[v] = struct{}{}
			ans = append(ans, v)
		}
	}
	return
}
