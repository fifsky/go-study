package main

import (
	"fmt"
	"time"
)

func main() {
	ctrl := make(chan int, 3)
	for i := 0; i < 50; i++ {
		ctrl <- 1
		go func(i int) {
			defer func() {
				<-ctrl
			}()
			time.Sleep(time.Second)
			//all print ?
			fmt.Println(i)
		}(i)
	}
}
