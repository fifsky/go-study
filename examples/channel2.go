package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	stop := make(chan bool)
	go func() {
		i := rand.Intn(10)
		time.Sleep(time.Duration(i) * time.Second)
		fmt.Println("Sleep:" + strconv.Itoa(i) + " s")
		stop <- true
	}()

	<-stop
	close(stop)
}
