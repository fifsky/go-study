package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	x int64
}

func (c *Counter) Inc() {
	c.x++
}

func main() {
	c := &Counter{}
	var wait sync.WaitGroup
	wait.Add(4)
	for n := 0; n < 4; n++ {
		go func() {
			defer wait.Done()

			for i := 0; i < 2500; i++ {
				c.Inc()
			}
		}()
	}
	wait.Wait()
	fmt.Println(c.x)
}

//go run -race lock1.go
