package main

import "fmt"

func main() {

	fmt.Println("a")

	i, err := tests(1)
	fmt.Println(i)
	fmt.Println(err)
}

func tests(a int) (int, error) {

	return a, nil
}
