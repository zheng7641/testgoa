package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	newFile *os.File
	err     error
)

func main() {

	fmt.Println("a1")

	newFile, err := os.Open("/Users/zct/Desktop/baocuo.txt")
	if err!=nil{
		log.Fatal(err)
	}
	data, err1 := ioutil.ReadAll(newFile)
	if err1 !=nil{
		log.Fatal(err1)
	}
	fmt.Println(string(data))

	newFile.Close()
}
