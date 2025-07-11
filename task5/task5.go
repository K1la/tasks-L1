package main

import (
	"fmt"
	"os"
	"time"
)

func worker(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
		time.Sleep(500 * time.Millisecond) // 0,5с
	}
}

func main() {

	ch := make(chan int, 2)

	go worker(ch)

	timeout := time.After(5 * time.Second)

	for {

		select {
		case i := <-ch:
			fmt.Printf("%d\n", i)
		case <-timeout:
			fmt.Println("timeout")
			os.Exit(0)
		}
	}
}
