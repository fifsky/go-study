package main

import "fmt"

type I2 interface {
	Get() int
}

type SI2 struct {
	Age int
}

func (s SI2) Get() int {
	return s.Age
}

func getAge(i I2) int {
	return i.Get()
}

func main() {
	ss := SI2{12}
	fmt.Println(getAge(&ss)) //ponter
	fmt.Println(getAge(ss))  //value

	//if func(s *SI2) ?
}
