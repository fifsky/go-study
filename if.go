package main

import "fmt"

func main() {
	var a int
	var b int32
	a = 1
	b = 1

	//Can not be
	if a == b {
		fmt.Println("ok")
	}

	c := 1
	//Why?
	if c == b {
	}

	c := 1
	//Why?
	if c == a {
	}

	var bb bool
	bb = true
	//Cant not be
	if bb == c {
	}

	//&& || else

	switch b {
	case 1:
		fmt.Printf("1")
	case 2:
		fmt.Printf("2")
		fallthrough
	case 3:
		fmt.Printf("3")
		//not php break
		break
	case 4, 5, 6:
		fmt.Printf("4, 5, 6")
	default:
		fmt.Printf("Default")
	}
}
