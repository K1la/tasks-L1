package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(arr))
	wg := &sync.WaitGroup{}

	for _, v := range arr {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			ch <- v * v
		}(v)

	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
