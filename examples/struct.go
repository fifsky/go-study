package main

import "fmt"

type User struct {
	UserId int
	Name   string
}

func main() {
	var user *User
	user = new(User)
	user.UserId = 1
	user.Name = "test"

	user2 := new(User)
	user2.UserId = 2
	user2.Name = "test2"

	user3 := &User{}
	user3.UserId = 3
	user3.Name = "test3"

	user4 := &User{
		UserId: 4,
		Name:   "test4",
	}

	fmt.Println(user, user2, user3, user4)

	//What? *User
	user5 := User{
		UserId: 5,
		Name:   "test5",
	}

	fmt.Println(user5)
}
