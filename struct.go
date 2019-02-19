package main

import (
	"fmt"
	"sync"
)

type TestStruct struct {
	UserId int
	Name   string
}

type TestStruct2 struct {
	mu *sync.Mutex
	i  int
}

func main() {
	var user *TestStruct
	user = new(TestStruct)
	user.UserId = 1
	user.Name = "test"

	user2 := new(TestStruct)
	user2.UserId = 2
	user2.Name = "test2"

	user3 := &TestStruct{}
	user3.UserId = 3
	user3.Name = "test3"

	user4 := &TestStruct{
		UserId: 4,
		Name:   "test4",
	}

	fmt.Println(user, user2, user3, user4)

	//What? TestStruct
	user5 := TestStruct{
		UserId: 5,
		Name:   "test5",
	}

	fmt.Println(user5)

	t := TestStruct2{}
	//t.mu.Unlock()
	fmt.Println(t)
}
