package main

import (
	"fazztrack/demo/calc"
	"fmt"
)
func fazz(harga int,voucher string,jarak int,pajak bool){
	discount := calc.Dicount(harga ,voucher)
	delivery := calc.DeliveryFee(jarak)
	taxs := calc.IsTaxed(harga,pajak)
	fmt.Printf("harga  :%d\n",harga)
	fmt.Printf("voucher : %d\n",discount)
	fmt.Printf("biaya pengiriman : %d\n",delivery)
	fmt.Printf("pajak : %d\n",taxs)
	
}

	
func main() {
	var harga int
	var voucher string
	var delivery int
	var taxs bool
	fmt.Print("harga :") // Just for beauty
	fmt.Scanln(&harga)
	fmt.Print("voucher :") // Just for beauty
	fmt.Scanln( &voucher)
	fmt.Print("biaya pengiriman :") // Just for beauty
	fmt.Scanln(&delivery)
	fmt.Print("pajak :") // Just for beauty
	fmt.Scanln(&taxs)
	fazz(harga,voucher,delivery,taxs)
}

// func main()  {
// 	fazz(75000,"DITRAKTIR60",1,false)
	
// }