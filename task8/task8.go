package main

import "fmt"

/*
Дана переменная типа int64. Разработать программу,
которая устанавливает i-й бит этого числа в 1 или 0.
*/

func setBit(val int64, pos int) int64 {
	return val ^ (1 << int64(pos))

}

/*
	ПРИМЕР:

	val=9
	val в байтах = 1001
	какой i-ый бит поменять в val
	n: 1

	num = 11 | 1011
                 ^
	1 бит был изменен с 0 на 1
*/

func main() {
	var (
		n   int
		val int64
	)
	fmt.Printf("val=")
	fmt.Scan(&val)
	fmt.Printf("val в байтах = %b\n", val)
	fmt.Printf("какой i-ый бит поменять в val\nn: ")
	fmt.Scan(&n)

	num := setBit(val, n)
	fmt.Printf("\nnum = %d | %b\n", num, num)
}
