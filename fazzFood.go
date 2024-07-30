package main

import "fmt"
func fazz(harga int,voucher string,jarak int,pajak bool){
var discount int 
  if voucher == "FAZZFOOD50" {
      if (harga >= 50000) {
        discount = harga *50/100
        if (discount * 50/100 > 50000) {
          discount = 50000
        }
      }
    }
    if voucher == "DITRAKTIR60" {
      if (harga >= 25000) {
        discount = harga * 60/100
        if (discount * 50/100 > 25000) {
          discount = 30000;
        }
      }
    }
	var delivery int = 5000
	if jarak > 2 {
		delivery += 3000*(jarak-2)
	}
	var taxs int 
	if pajak {
		taxs = harga * 5/100 
	}
	fmt.Printf("harga %d,voucher %d,jarak %d,pajak %d",harga,discount,delivery,taxs)
}
	


func main()  {
	fazz(25000,"DITRAKTIR60",1,false)
	
}