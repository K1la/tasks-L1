package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("start: %v\n", time.Now())
	sleep(2)
	fmt.Printf("end: %v\n", time.Now())
}

func sleep(duration int) {
	ch := make(chan struct{})

	d := time.Duration(duration) * time.Second
	time.AfterFunc(d, func() {
		ch <- struct{}{}
	})

	<-ch
}
