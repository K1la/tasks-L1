package main

import (
	"fmt"
	"runtime"
	"strings"
)

func printMem(label string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%s:\tAlloc = %v KB\n", label, m.Alloc/1024)
}

func createHugeString(size int) string {
	return strings.Repeat("A", size)
}

var justString string

func someFunc() {
	printMem("Before allocation")
	/*
		чтобы было нагляднее можно использовать не 1 << 10, а 1 << 30
	*/
	v := createHugeString(1 << 10) // 1024 байт / 1КБ, (1024 символов)
	/*
		в justString сохраняется ссылка на `v` и используется 100 байт
		здесь просто копия `v`, в памяти все еще будет использоваться 1КБ,
		до тех пор, пока строка justString используется
	*/
	// старая версия
	//justString = v[:100]

	// исправление
	// 1 вариант
	// justString = string([]byte(v[:100]))
	// 2 вариант
	justString = strings.Clone(v[:100])

	/*
		в переменную `v` присваиваю пустую строку,
		чтобы обнулить 1КБ в ее памяти
		и запускаю сборщик мусора
	*/
	v = ""
	runtime.GC()
	printMem("After GC")
}

func main() {
	someFunc()
}
