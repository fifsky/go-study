package main

import "fmt"

type Foo int

func (f *Foo) Int() {
	fmt.Println(f)
}

func main() {
	//Foo(123).Int()
	Foo.Int(123)
}
