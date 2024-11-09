package main

import (
	"design-patterns/creational/internal/singleton"
	"sync"
)

func runGoroutine(wg *sync.WaitGroup) {
	singleton.GetInstance("value in singleInstance")
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	n := 10
	wg.Add(n)
	for i := 0; i < n; i++ {
		go runGoroutine(wg)
	}
	wg.Wait()
}
