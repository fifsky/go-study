package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/verystar/golib/task"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	task.NewTaskPool(3, func(pool *task.TaskPool) {
		for i := 1; i <= 50; i++ {
			pool.Add(1)
			go func(i int) {
				s := rand.Intn(3)
				time.Sleep(time.Duration(s) * time.Second)
				fmt.Println("Printer:" + strconv.Itoa(i) + " Sleep:" + strconv.Itoa(s))
				pool.Done()
			}(i)
		}
	}).Run()
}
