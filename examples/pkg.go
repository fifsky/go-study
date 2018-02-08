package main

import (
	"encoding/json"
	"fmt"

	"app/core/util"

	"github.com/verystar/golib/logger"
)

type Users struct {
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
}

func main() {

	m := util.Md5("test")
	fmt.Println(m)

	h := &Users{
		UserId: 123,
		Name:   "test",
	}

	b, err := json.Marshal(h)

	if err != nil {
		fmt.Println(b)
	}

	hh := &Users{}
	err = json.Unmarshal(b, hh)

	if err != nil {
		logger.Debug("HH%+v", hh)
	}
}
