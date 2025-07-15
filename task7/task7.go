package main

import (
	"fmt"
	"log"
	"sync"
)

/*
Реализовать безопасную для конкуренции запись данных в структуру map.

Используется: sync.Mutex
*/

type RaceMap struct {
	mu sync.Mutex
	m  map[int]string
}

func (s *RaceMap) Set(key int, val string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = val
}

func (s *RaceMap) Get(key int) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	val, ok := s.m[key]
	return val, ok
}

func main() {
	rmap := RaceMap{m: make(map[int]string)}

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			val := fmt.Sprintf("go chan=%d", i)
			rmap.Set(i, val)
			if v, ok := rmap.Get(i); ok {
				log.Printf("Read %d = %s\n", i, v)
			}
		}(i)
	}

	wg.Wait()

}
