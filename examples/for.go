package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr); i++ {
		fmt.Println(i)
	}

	//Why?
	fmt.Println(i)

	m := map[string]interface{}{
		"user_id": 123,
		"name":    "test",
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
