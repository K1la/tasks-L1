package main

import "fmt"

/*
Поменять местами два числа без использования временной переменной.

Используется XOR-обмен (исключающее или)
*/

func main() {
	a := 31
	b := 46

	a ^= b
	b ^= a
	a ^= b

	fmt.Println(a) // 46
	fmt.Println(b) // 31
}
