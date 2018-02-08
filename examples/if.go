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
	case 2:
	case 3:
		//not php break
		break
	}
}
