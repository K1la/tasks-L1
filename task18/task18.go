package main

import (
	"fmt"
	"sync"
)

type Count struct {
	count int
	mu    sync.Mutex
}

func main() {
	c := new(Count)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 4; i++ {
			c.inc()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			c.inc()
		}
	}()

	wg.Wait()
	fmt.Println(c.count)
}

func (c *Count) inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Count) get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
