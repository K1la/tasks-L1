package main

import "fmt"

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
