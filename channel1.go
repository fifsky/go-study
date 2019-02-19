package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			fmt.Println(i)
		}
	}()

	//<-time.After(15 * time.Second)
}
