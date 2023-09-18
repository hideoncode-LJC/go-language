package main

import (
	"log"
	"math/rand"
	"time"
)

type Signal chan struct{}

/*
轮询定时查询
*/

func doTask(s Signal) {
	t := 0

	for {
		select {
		case <-s:
			log.Println("Task Done !")
			break
		default:
			log.Println("Task Go on!")
		}
		time.Sleep(time.Second)
		log.Printf("Task 已经运行 %d 秒", t)
		t++
	}
}

func main() {

	s := make(Signal)

	go doTask(s)

	for {
		n := rand.Int()

		if n%2 == 0 {
			s <- struct{}{}
			time.Sleep(3 * time.Second)
			break
		}

		time.Sleep(time.Second)

	}
}
