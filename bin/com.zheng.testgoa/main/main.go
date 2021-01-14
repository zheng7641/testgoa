package main

import "runtime"

func main() {
	var i int = 0
	runtime.GOMAXPROCS(20)
	for {
		i++
		go test123()
		if i%1000000 == 0 {
			println(i)
		}
	}
}

func test123() {
	var i int
	i++
}
