package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan int)

	for i := 1; i <= 50; i++ {
		go worker(channel)
	}

	for i := 1; i <= 100; i++ {
		channel <- i
	}

}

func worker(channel chan int) {
	for i := range channel {
		fmt.Println(i)
		time.Sleep(time.Second * 5)
	}
}
