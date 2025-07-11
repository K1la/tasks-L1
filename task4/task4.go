package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

func out(ch chan int, i int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {

		select {
		case val := <-ch:
			fmt.Printf("\nworker: %d\tvalue: %v\n", i, val)

		case <-ctx.Done():
			fmt.Printf("\nworker %d done\n", i)
			return
		}
	}

}

func main() {
	var n int
	fmt.Print("Enter n: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Error: ", err)
		return
	}
	ch := make(chan int, 10000)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go out(ch, i, ctx, &wg)

	}

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for {
			fmt.Printf("\nEnter number: ")
			if !scanner.Scan() {
				break
			}

			text := scanner.Text()
			val, err := strconv.Atoi(text)
			if err != nil {
				fmt.Println("Invalid input, must be a number.")
				continue
			}
			ch <- val
		}
	}()
	<-ctx.Done()
	fmt.Println("Stopped in main!")

	wg.Wait()
	fmt.Println("MAIN all workers done")

}
