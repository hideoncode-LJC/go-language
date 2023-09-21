package main

import (
	"context"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()

	go handle(ctx, )
}