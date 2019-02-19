package main

import "fmt"

type T struct{}

func (t T) say() {
	fmt.Println("hello")
}

func main() {
	T{}.say()

	//(&T{}).say()
	//((*T)(nil)).say()
}
