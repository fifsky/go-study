package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)

	go func() {
		for ch := range channel {
			time.Sleep(time.Second)
			fmt.Println(ch)
		}
	}()

	for i := 0; i < 10; i++ {
		channel <- i
	}
	close(channel)
}
