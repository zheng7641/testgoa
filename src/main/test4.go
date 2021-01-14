package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type testt struct {
		a string `json:"a"`
	}

	//{"a":"a"}

	var a string = `{"a":"a"}`
	var testtt testt
	var err = json.Unmarshal([]byte(a), &testtt)
	if err == nil {
		fmt.Println(testtt.a)
	}
}
