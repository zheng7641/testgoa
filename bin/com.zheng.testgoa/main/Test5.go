package main

func Test1(a int) int {
	return a
}

func main() {
	var x int = 1
	var a = func(x int) {
		println(x)
	}
	x += 2
	a(x)
}
