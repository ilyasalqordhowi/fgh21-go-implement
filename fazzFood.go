package main

import "fmt"
func fazz(harga int,voucher string,jarak int,pajak bool){
	discount := dicount(harga ,voucher)
	delivery := deliveryFee(jarak)
	taxs := isTaxed(harga,pajak)
	fmt.Printf("harga  :%d\n",harga)
	fmt.Printf("voucher : %d\n",discount)
	fmt.Printf("biaya pengiriman : %d\n",delivery)
	fmt.Printf("pajak : %d\n",taxs)
	
}
func dicount(harga int,voucher string)int{
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
        if (discount * 60/100 > 25000) {
          discount = 30000;
        }
      }
    }
	return discount
}
func deliveryFee(jarak int) int {
	var delivery int = 5000
	if jarak > 2 {
		delivery += 3000*(jarak-2)
	}
	return delivery
	
}
func isTaxed(harga int,pajak bool) int {
var taxs int 
	if pajak {
		taxs = harga * 5/100 
	}  
return taxs
	
}
	


func main()  {
	fazz(75000,"DITRAKTIR60",1,false)
	
}