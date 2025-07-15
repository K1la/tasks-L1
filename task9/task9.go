package main

import "fmt"

/*
	Разработать конвейер чисел.
	Даны два канала: в первый пишутся числа x из массива,
	во второй – результат операции x*2. После этого данные из
	второго канала должны выводиться в stdout.
	То есть, организуйте конвейер из двух этапов с
	горутинами: генерация чисел и их обработка.
	Убедитесь, что чтение из второго канала корректно завершается.
*/

func ch1(arr []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, v := range arr {
			out <- v
		}
	}()
	return out
}

func ch2(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			out <- v * 2
		}
	}()
	return out
}

func main() {
	arr := []int{2, 3, 4, 5, 6}

	cha1 := ch1(arr)
	cha2 := ch2(cha1)

	for v := range cha2 {
		fmt.Println(v)
	}
}
