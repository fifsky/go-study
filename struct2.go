package main

import "fmt"

type People struct{}

func (p *People) Say() {
	p.Hello()
}
func (p *People) Hello() {
	fmt.Println("Hello")
}

type Teacher struct {
	People
}

func (t *Teacher) Hello() {
	fmt.Println("Teacher Hello")
}

func main() {
	t := Teacher{}
	t.Say()
}
