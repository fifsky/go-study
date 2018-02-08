package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	channel := make(chan int)
	ctrl := make(chan int, 3)

	rand.Seed(time.Now().UnixNano())
	go func() {
		for ch := range channel {
			ctrl <- 1
			go func(ch int) {
				fmt.Println("LEN:", len(ctrl))
				i := rand.Intn(3)
				time.Sleep(time.Duration(i) * time.Second)
				fmt.Println("Printer:" + strconv.Itoa(ch) + " Sleep:" + strconv.Itoa(i))
				<-ctrl
			}(ch)
		}
	}()

	//all exec?
	for i := 1; i <= 50; i++ {
		channel <- i
	}
	close(channel)
}
