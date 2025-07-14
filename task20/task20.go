package main

import "fmt"

func main() {
	s := "snow dog sun"
	b := reverse(s)
	fmt.Println(b)
}

func reverse(s string) string {
	r := []rune(s)

	// переворачиваем всю строку "snow dog sun"
	reverseWord(r, 0, len(r)-1) // nus god wons

	/*
		ожидаем либо конца строки, либо знака пробела,
		затем отправляем начало и конец строки, который
		будем переворачивать
	*/
	left := 0
	for i := 0; i <= len(r); i++ {
		/*
			i == len(r) стоит первее,
			чтобы не было проблем с индексом
		*/
		if i == len(r) || r[i] == ' ' {
			reverseWord(r, left, i-1)
			left = i + 1
		}
	}

	return string(r)
}

func reverseWord(s []rune, start, end int) {
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}
