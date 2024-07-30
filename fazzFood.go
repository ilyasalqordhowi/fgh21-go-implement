package main

import (
	"fazztrack/demo/fazz"
	"fmt"
)


	
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
	fazz.Fazz(harga,voucher,delivery,taxs)
}

// func main()  {
// 	fazz(75000,"DITRAKTIR60",1,false)
	
// }