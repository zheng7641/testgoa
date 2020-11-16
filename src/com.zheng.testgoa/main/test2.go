package main

import (
	"fmt"
	"time"
)

type phone interface {
	call()
}

type iphone struct {
}

type nokia struct {
}

func (i iphone) call() {
	fmt.Println("iphone")
}

func (n nokia) call() {
	fmt.Println("nokia")
}

func gotest(b string, a *int) {
	*a = *a + 1
	fmt.Println(b, *a)
}

func chantest(a *int, chana chan int) {
	*a += 1
	fmt.Println("chantest", *a)
	chana <- *a
}

func chantest2(chana chan int) {
	var a int = <-chana
	a += 1
	fmt.Println("chantest2", a)
	chana <- a
}

func main() {
	//var p phone
	//p = new(iphone)
	//p.call()
	//p=new(nokia)
	//p.call()

	//var goint *int =new(int)
	//*goint = 1
	//go gotest("go",goint)
	//gotest("not",goint)
	//gotest("not",goint)
	//time.Sleep(100 * time.Millisecond)
	//
	//var ch chan int = make(chan int,100)
	//ch <- 1
	//var aw =<- ch
	//fmt.Println(aw)

	var ch2 chan int = make(chan int, 100)
	var chani *int = new(int)
	*chani = 1
	go chantest(chani, ch2)
	time.Sleep(100 * time.Millisecond)
	go chantest2(ch2)
	time.Sleep(100 * time.Millisecond)
	var chanre int = <-ch2
	fmt.Println("chanre", chanre)
}
