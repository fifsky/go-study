package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	go func() {
		i := rand.Intn(10)
		time.Sleep(time.Duration(i) * time.Second)
		fmt.Println("Sleep:" + strconv.Itoa(i) + " s")
	}()

	for {
		select {
		case <-time.NewTicker(1 * time.Second).C:
		}
	}
}
