package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func reqTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("Task Done!")
			return
		default:
			fmt.Println("Task Go on!")
			time.Sleep(time.Second * 1)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go reqTask(ctx)

	time.Sleep(time.Second * 3)

	var num int

	for {
		num = rand.Int()

		if num%2 == 0 {
			cancel()
			time.Sleep(time.Second * 1)
			break
		}
	}

}
