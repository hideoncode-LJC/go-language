package main

import (
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup

func doTask(n int) {
	time.Sleep(time.Duration(n))
	fmt.Printf("Task %d Done\n", n)
	wg.Done()
}

func main() {
	for i := 0 ; i < 3 ; i++ {
		wg.Add(1)
		go doTask(i)
	}

	wg.Wait()
	fmt.Println("All Task Done")
}