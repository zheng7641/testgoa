package main

func main() {
	var m = make(map[string]int)
	m["a"] = 1
	delete(m, "a")

}
