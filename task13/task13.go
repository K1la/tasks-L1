package main

import "fmt"

func main() {
	a := 31
	b := 46

	a ^= b
	b ^= a
	a ^= b

	fmt.Println(a) // 46
	fmt.Println(b) // 31
}
