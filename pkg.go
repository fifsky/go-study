package main

import (
	"encoding/json"
	"fmt"

	"examples/util"

	"github.com/verystar/logger"
)

type JsonTest struct {
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
}

func main() {

	m := util.Md5("test")
	fmt.Println(m)

	h := &JsonTest{
		UserId: 123,
		Name:   "test",
	}

	b, err := json.Marshal(h)

	if err != nil {
		fmt.Println(b)
	}

	hh := &JsonTest{}
	err = json.Unmarshal(b, hh)

	if err != nil {
		logger.Debug("HH%+v", hh)
	}
}
