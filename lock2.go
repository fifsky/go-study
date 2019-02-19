package main

import (
	"fmt"
	"sync"
)

type Counter2 struct {
	mu sync.Mutex
	x  int64
}

func (c *Counter2) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.x++
}

func main() {
	c := &Counter2{}
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
