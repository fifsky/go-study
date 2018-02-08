package main

import "fmt"

type IModel interface {
	GetTableName() string
	Get(key string) interface{}
}

//var _ IModel = (*UserModel)(nil)

type UserModel struct {
	UserId int
	Name   string
}

func (u *UserModel) GetTableName() string {
	return "users"
}

func (u *UserModel) Get(key string) interface{} {
	if key == "user_id" {
		return u.UserId
	} else if key == "name" {
		return u.Name
	}

	return nil
}

//var _ IModel = (*ProviderModel)(nil)
//var _ IModel = &ProviderModel{}

type ProviderModel struct {
	UserId int
	Name   string
}

func (u *ProviderModel) Get(key string) interface{} {
	if key == "user_id" {
		return u.UserId
	} else if key == "name" {
		return u.Name
	}
	return nil
}

func main() {
	user := &UserModel{
		UserId: 123,
		Name:   "test",
	}

	user_id := user.Get("user_id")

	var id int
	id = 123
	//id = user_id

	if user_id == id {
		fmt.Println("OK", user, user_id)
	} else {
		fmt.Println("Error", user_id)
	}
}
