package main

func main() {
	var i interface{}
	i = &UserModel{}

	//Why?
	i.Get("name")
}
