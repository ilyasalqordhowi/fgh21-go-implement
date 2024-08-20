package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Data struct {
	name string
}

func Out(in []string)  {
	if len(in)%2 != 0 {
		panic("Error: slice should be even")
	}
	left := 0
	right := 0
	 for  _,v := range in[:len(in)/2] {
        num,_  := strconv.Atoi(v)
        left += num
    }
	 for _,v := range in[len(in)/2:] {
        num,_ := strconv.Atoi(v)
        right += num
    }
	fmt.Println(left)
	fmt.Println(right)

}

func Error(){
	if e:=recover();e != nil{
		if strings.Contains(e.(string),"should be even") {
			fmt.Println("Angka dalam slice berjumlah ganjil")
		}
	}
}

func main() {
	defer Error()
	s := []string{"2", "3", "4", "5","6","7"}
	Out(s)
}

// data := Data{
// 	name: "mas syarif",
// }

// dataUpdate := &data.name
// *dataUpdate = "oke"
// fmt.Print(data)
// fmt.Print(*dataUpdate)
// fazz.CalcFazzFood()
// fazz.Survey()
// fazz.Matrix()
// 	a := 123
// 	b := &a
// 	*b = *b - 23
// 	fmt.Print(a)
// 	fmt.Print(*b)