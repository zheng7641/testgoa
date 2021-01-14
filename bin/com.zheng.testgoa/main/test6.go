package main

import fmt "fmt"

func main() {
	//var m = make(map[string]int,1000)
	//m["a"] = 1
	//delete(m, "a")
	//
	//var m1 map[string]int
	//println(m1)

	u := struct {
		a int
		b string
	}{
		a: 1,
		b: "b",
	}
	fmt.Println(u)

	type data struct {
		a int
		b int
	}

	var a data = data{
		1,
		1,
	}
	fmt.Println(a)

	type data1 struct {
		string
		int
	}
	var d1 data1 = data1{
		"a",
		1,
	}

	fmt.Println(d1)

	var t1 test
	t1.testa()

	f1 := test.testa
	fmt.Println(f1)
}

type test struct {
	a string
}

func (a test) testa() int {
	fmt.Println("a")
	return 1
}
