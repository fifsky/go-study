package main

import (
	"errors"
	"fmt"
)

type Oooops struct {
	code int
	msg  string
}

func Recover()  {
	defer func() {
		if re := recover(); re != nil {
			op := re.(Oooops)
			fmt.Println(op.msg)
		}
	}()
}


func main() {
	//Recover()
	defer func() {
		if re := recover(); re != nil {
			op := re.(Oooops)
			fmt.Println(op.msg)
		}
	}()

	var err error
	err = errors.New("this is error")
	fmt.Println(err)

	panic(&Oooops{
		code: 123,
		msg:  "this is panic",
	})
}
