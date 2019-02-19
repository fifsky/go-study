package main

import "fmt"

func main() {
	type TString string
	type TMap map[string]string
	type TSlice []map[string]string

	var sl1 TSlice
	var sl2 []map[string]string
	sl1 = append(sl1, map[string]string{"name": "verystar"})
	sl2 = sl1
	fmt.Println(sl1, sl2)
}
