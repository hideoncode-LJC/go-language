package main

import (
	"context"
	"fmt"
	"time"
)

type Options struct {
	Interval time.Duration
}

func reqTask(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop", name)
			return
		}
		default:
			fmt.Println(name, )
	}
}
