package main

import (
	"fmt"
)

func out(ch chan int, i int) {
	for v := range ch {
		fmt.Printf("\nworker: %d\tvalue: %v\n", i, v)

	}
}

func main() {
	var n, val int
	fmt.Print("Enter n: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Error: ", err)
		return
	}
	ch := make(chan int, 10000)

	go func() {
		for {
			fmt.Print("\nEnter val: ")
			if _, err := fmt.Scan(&val); err != nil {
				fmt.Println("Invalid input. Try again.")
				continue
			}
			ch <- val
		}
	}()

	for i := 1; i <= n; i++ {
		go out(ch, i)

	}

	// select для того, чтобы main не завершался до того,
	// пока сами не захотим завершить программу
	select {}
}
