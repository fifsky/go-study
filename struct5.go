package main

import "fmt"

type S5 struct {
	Age int
}

func (s S5) Get() int {
	return s.Age
}

func (s S5) Set(age int) {
	s.Age = age
}

func main() {
	ss := S5{}
	ss.Set(12)
	fmt.Println(ss.Get())

	//ss.Age = 12
	//fmt.Println(ss.Age)
	//if func(s *S5) why ?  PS:Foo.Int(123)
	// 等价于
	//S5.Set(ss,12)
	//fmt.Println(S5.Get(ss))
}
