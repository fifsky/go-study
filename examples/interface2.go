package main

import "fmt"

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
}