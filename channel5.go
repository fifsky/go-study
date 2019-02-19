package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ctrl := make(chan int, 3)
	wg := new(sync.WaitGroup)

	for i := 0; i < 50; i++ {
		ctrl <- 1
		wg.Add(1)
		go func(i int) {
			time.Sleep(time.Second)
			fmt.Println(i)
			wg.Done()
			<-ctrl
		}(i)
	}

	wg.Wait()
}
