package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "aabcd"
	b := check(s)
	fmt.Println(b)
}

func check(s string) bool {
	m := make(map[rune]struct{})
	for _, v := range strings.ToLower(s) {
		_, inM := m[v]
		if !inM {
			m[v] = struct{}{}
		} else {
			return false
		}
	}
	return true
}
