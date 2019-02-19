package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Second)
			fmt.Println(i)
			//What print?
		}(i)
	}

	<-time.After(10 * time.Second)
}
