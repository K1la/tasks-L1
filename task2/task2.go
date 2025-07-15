package main

import (
	"fmt"
	"sync"
)

// Написать программу, которая конкурентно рассчитает
// значения квадратов чисел, взятых из массива [2,4,6,8,10],
// и выведет результаты в stdout.

func main() {
	arr := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(arr))
	wg := &sync.WaitGroup{}

	for _, v := range arr {
		// Каждой горутине инкрементим счетчик WaitGroup
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			ch <- v * v
		}(v)

	}

	// ждем выполнения всех WaitGroup и закрываем канал
	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
