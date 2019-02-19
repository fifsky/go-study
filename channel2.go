package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Sleep:5s")
		stop <- true
	}()

	<-stop
	close(stop)
}
