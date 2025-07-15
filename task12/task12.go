package main

import "fmt"

/*
	Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree").
	Создать для неё собственное множество.

	Ожидается: получить набор уникальных слов.
	Для примера, множество = {"cat", "dog", "tree"}.
*/

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
