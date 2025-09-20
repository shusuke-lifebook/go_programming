package main

import (
	"fmt"
	"sync"
)

func normal(s string) {
	for i := 0; i < 5; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func goroutine(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Add(1)
	go goroutine("world", &wg)
	normal("hello")
	// time.Sleep(2000 * time.Microsecond)
	wg.Wait()
}
