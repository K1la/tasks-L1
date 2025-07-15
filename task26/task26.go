package main

import (
	"fmt"
	"strings"
)

/*
	Уникальные символы в строке

	Разработать программу, которая проверяет,
	что все символы в строке встречаются один раз
	(т.е. строка состоит из уникальных символов).

	Вывод: true, если все символы уникальны, false,
	если есть повторения. Проверка должна быть
	регистронезависимой, т.е. символы в разных
	регистрах считать одинаковыми.

	Например: "abcd" -> true,
	"abCdefAaf" -> false (повторяются a/A),
	"aabcd" -> false.
*/

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
