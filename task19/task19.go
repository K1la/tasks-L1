package main

import "fmt"

/*
	Разворот строки

	Разработать программу, которая переворачивает подаваемую на вход строку.

	PS: учтено, что символы могут быть в Unicode (русские буквы, emoji и пр.)
	Используется массив rune
*/

func main() {
	str := ""
	fmt.Scan(&str)
	runestr := []rune(str)
	for i := len(runestr) - 1; i >= 0; i-- {
		fmt.Print(string(runestr[i]))
	}
}
