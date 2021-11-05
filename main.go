package main

import "fmt"

type person struct {
	first string
}

func (p person) speak () {
	fmt.Println("from a person -- this is my name", p.first)
}

type human interface {
	speak()
}

func main() {
	p1 := person{
		first: "James",
	}

	fmt.Printf("%T\n", p1)

	var x human
	x = p1
	x.speak()
	fmt.Printf("%T\n", x)
}

/*
//go:noinline
func foo() int {
	x := 42
	return x
}

//go:noinline
func bar() *int {
	y := 43
	return &y
}

func main() {
	_ = foo()
	_ = bar()
}
*/