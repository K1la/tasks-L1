package main

import (
	"fmt"
	"time"
)

/*
	Своя функция Sleep

	Реализовать собственную функцию sleep(duration)
	аналогично встроенной функции time.Sleep,
	которая приостанавливает выполнение текущей горутины.
*/

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
