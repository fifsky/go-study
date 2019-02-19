package main

import "fmt"

type S struct{}

func (s S) Get() {
	fmt.Println("hello")
}

func main() {
	var a interface{}
	a = 1
	var b int
	b = 2

	//can not plus
	c := a + b

	//aa := a.(int)
	//c := aa + b
	fmt.Println(c)

	//aa := a.(string)
	//fmt.Println(aa,ok)

	switch aa := a.(type) {
	case int:
		fmt.Println(aa + b)
	}

	//------
	var s interface{}
	s = S{}

	//Why?
	s.Get()

	//------
	var i1 interface{}
	var i2 int
	var i3 int32
	i1 = 1
	i2 = 1
	i3 = 1
	if i1 == i2 {
		fmt.Println("OK")
	} else {
		fmt.Println("Error")
	}

	//if i2 == i3 {
	//	fmt.Println("OK")
	//} else {
	//	fmt.Println("Error")
	//}
}
