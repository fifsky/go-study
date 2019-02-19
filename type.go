package main

import "fmt"

func main() {
	var i int
	i = 1
	fmt.Println(i)

	var i64 int64
	i64 = 1
	//i = i64
	fmt.Println(i64)

	//Why?
	a := 1
	i = a

	var s string
	s = "abcddddd"
	fmt.Println(s)

	var b byte
	b = 's'
	fmt.Println(b)

	var f float64
	f = 0.15
	fmt.Println(f)

	var arr [10]int
	arr[0] = 1
	arr[1] = 2
	arr[3] = 5

	var sarr []int
	//Why?
	sarr[0] = 1
	sarr[1] = 2

	var m map[string]int
	//Why?
	m["c"] = 1
	m["d"] = 2

	var it interface{}
	it = 1
	//Why not error?
	it = "cccc"
	fmt.Println(it)

	var m2 map[string]interface{}
	m2["user_id"] = 123
	m2["user_name"] = "Test"

	//------
	//https://stackoverflow.com/questions/19334542/why-can-i-type-alias-functions-and-use-them-without-casting
	type TString string
	type TMap map[string]string
	type TSlice []map[string]string

	var s1 TString = "verystar"
	var s2 string
	s2 = s1
	fmt.Println(s1, s2)

	var m1 TMap = TMap{"name": "verystar"}
	var m2 map[string]string
	m2 = m1
	fmt.Println(m1, m2)

	var sl1 TSlice
	var sl2 []map[string]string
	sl1 = append(sl1, map[string]string{"name": "verystar"})
	sl2 = sl1
	fmt.Println(sl1, sl2)
}
