package main

func main() {
	var i int = 0
	for {
		i++
		go println("go")
		println(i)
	}

}
