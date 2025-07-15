package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.

Классические подходы: выход по условию, через канал уведомления,
через контекст, прекращение работы runtime.Goexit() и др.
*/

func main() {
	fmt.Println("start main")
	var wg sync.WaitGroup
	wg.Add(4)

	// горутина с выходом через условие
	go ifgo(&wg)

	// горутина с выходом через канал
	done := make(chan struct{})
	go chanDone(done, &wg)

	// горутина с выходом через контекст
	ctx, stop := context.WithCancel(context.Background())
	go ctxFunc(ctx, stop, &wg)

	// горутина с выходом через runtime.GoExit()
	go goexitFunc(&wg)

	wg.Wait()
	fmt.Println("end main")
}

func ifgo(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		if i == 3 {
			log.Printf("stopped ifgo i=%d\n", i)
			return
		} else {
			log.Printf("running ifgo i=%d\n", i)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func chanDone(done chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	go func() {
		// закрываем канал done через 1с,
		// чтобы горутина chanDone закончила работать
		time.Sleep(1 * time.Second)
		close(done)
	}()
	for {
		select {
		case <-done:
			log.Println("goroutine `done` exit, after 1s")
			return
		default:
			log.Println("goroutine `done` running")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func ctxFunc(ctx context.Context, stop context.CancelFunc, wg *sync.WaitGroup) {
	defer wg.Done()
	go func() {
		time.Sleep(2 * time.Second)
		stop()
	}()

	for {
		select {
		case <-ctx.Done():
			log.Println("goroutine `ctxFunc` exit, after 2s")
			return
		default:
			log.Println("goroutine `ctxFunc` running")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func goexitFunc(wg *sync.WaitGroup) {
	for i := 0; i < 2; i++ {
		log.Println("goroutine `goexitFunc` running i=", i)
	}
	wg.Done()
	log.Println("goroutine `goexitFunc` exit")
	runtime.Goexit()
	log.Println("goroutine `goexitFunc` after goexit")
}
