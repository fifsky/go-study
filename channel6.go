package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			channel1 <- 1
		}
	}()

	go func() {
		for {
			time.Sleep(3 * time.Second)
			channel1 <- 3
		}
	}()

	for {
		select {
		case i := <-channel1:
			fmt.Println(i)
		case j := <-channel2:
			fmt.Println(j)
			//default:
			//fmt.Println(time.Now())
		}
	}
}
