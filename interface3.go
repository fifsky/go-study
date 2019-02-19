package main

import "fmt"

type I3 interface {
	Get() string
	Set(string)
}

//var _ I3 = (*SI3)(nil)

type SI3 struct {
	Name string
}

func (s *SI3) Set(name string) {
	s.Name = name
}

func (s *SI3) Get() string {
	return s.Name
}

func main() {
	ss := &SI3{
		Name: "fifsky",
	}

	fmt.Println(ss.Get())
}
