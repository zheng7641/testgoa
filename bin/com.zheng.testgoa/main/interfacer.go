package main

import "fmt"

func main() {
	var d1 data = data{
		"a",
	}
	fmt.Println(d1.test())

	var d2 tester = &d1
	fmt.Println(d2.test())
}

type tester interface {
	test() string
}

type data struct {
	a string
}

func (f data) test() string {
	return f.a
}
