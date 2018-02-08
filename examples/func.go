package main

import (
	"fmt"
	"errors"
)

func foo(a int, b int) int {
	return a + b
}

func foo2(a int, b int) (c int) {
	c = a + b
	return c
}

func foo3(a int, b int) (int, error) {

	if a == 0 {
		return 0, errors.New("a must > 0")
	}

	return a + b, nil
}

func main() {

	c := foo(1, 2)
	fmt.Println(c)

	var fn = func() {
		//var scope
		fmt.Println(c)
	}

	fn()
}
