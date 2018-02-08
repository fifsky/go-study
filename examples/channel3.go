package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	channel := make(chan int)

	go func() {
		for ch := range channel {
			i := rand.Intn(3)
			time.Sleep(time.Duration(i) * time.Second)
			fmt.Println("Printer:" + strconv.Itoa(ch) + " Sleep:" + strconv.Itoa(i))
		}
	}()

	for i := 1; i <= 100; i++ {
		channel <- i
	}
	close(channel)
}
