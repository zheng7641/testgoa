package main

import (
	"fmt"
)
import _ "os"

func main() {
	fmt.Println("Hello, World!")

	var s, sep string
	for i := 1; i < 10; i++ {
		s += sep + ("" + "1")
		sep = ""
	}
	fmt.Print(s)
	fmt.Println(1)
	var age int32 = 9
	fmt.Println(age)

	var a, b int = 2, 2
	fmt.Println(a + b)
	fmt.Println(a, b)

	var stra string
	fmt.Println(stra)

	var inta int
	fmt.Println(inta)

	var bc *int
	fmt.Println(bc)

	var ara []int
	fmt.Println(ara)

	var ma map[string]string = make(map[string]string)
	mb := make(map[string]string)
	ma["a"] = "a"
	mb["b"] = "b"
	fmt.Println(ma)
	fmt.Println(mb)

	var dd = 2
	fmt.Println(dd)

	fmt.Println(&dd)
	var ss *int = &dd
	*ss += 10
	fmt.Println(*ss)

	var p *string
	fmt.Println(p)

	const ca int = 1
	fmt.Println(ca)

	const cb = "cb"
	fmt.Println(cb)

	var adda int = 1
	var addb int = 2
	var addc int = adda + addb
	fmt.Println(addc)
	var addd int = adda / addb
	fmt.Println(addd)
	var adde float32 = float32(adda / addb)
	fmt.Println(adde)

	var addf float64 = 22.22
	var addg float64 = float64(adda) / addf

	var addh = 1.1 / 1.1
	fmt.Println(addh)
	fmt.Println(addg)

	var ia int = 110
	if ia < 10 {
		ia++
		fmt.Println(ia)
	} else {
		fmt.Println("b")
	}

	for ii := 1; ii < 10; ii++ {
		fmt.Println(ii)
	}

	var changdu = len("aaa")
	fmt.Println(changdu)
	var wb = 1
	var result int = max("max", &wb)
	fmt.Println("wb", wb)
	fmt.Println(result)

	//var fuc = max
	//var resulta = fuc("fuc",1)
	//fmt.Println(resulta)
	var arraya = [4]int{1, 2, 3, 4}
	fmt.Println(arraya)

	var arrayb = [3]string{"a", "b", "c"}
	fmt.Println(arrayb)

	var arrayc = [...]int{1, 2, 3}
	fmt.Println(arrayc)

	var arrayd [2][2]int
	fmt.Println(arrayd)
	var arraye = [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(arraye)
	for e1 := 0; e1 < 2; e1++ {
		for e2 := 0; e2 < 2; e2++ {
			fmt.Println(arraye[e1][e2])
		}
	}
	var strua = structa{1, 2, "c", "d"}
	fmt.Println(strua)
	fmt.Println(strua.stra)

	var strub *structa = &strua
	fmt.Println(strub)

	var sli = make([]int, 3, 20)
	var newslii []int
	for slii := 0; slii < 20; slii++ {
		newslii = append(sli, slii)
	}
	fmt.Println(newslii)
	//fmt.Println(sli)

	var mapa map[string]string
	fmt.Println(mapa)
	mapa = make(map[string]string)
	mapa["a"] = "a"
	fmt.Println(mapa)
	delete(mapa, "a")
	fmt.Println(mapa)

	var sli1 = make([]int, len(sli), (cap(sli) * 2))
	fmt.Println(sli1)
	sli1[1] = 1

	for as, num := range sli1 {
		fmt.Println(as, num, sli1[num], sli1[as])
	}
	fmt.Println()
	for num := range sli1 {
		fmt.Println(num, sli1[num])
	}

	//errors.New("a")
	var lenarr [3]int
	lenarr[0] = 0
	fmt.Println(len(lenarr))

}

func max(wa string, wb *int) int {
	fmt.Println(wa, *wb)
	*wb++
	return 1
}

type structa struct {
	inta int
	intb int
	stra string
	strb string
}
