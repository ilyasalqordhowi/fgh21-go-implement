package fazz

import (
	"fazztrack/demo/calc"
	"fmt"
)

func FazzFood(harga int,voucher string,jarak int,pajak bool){
	discount := calc.Dicount(harga ,voucher)
	delivery := calc.DeliveryFee(jarak)
	taxs := calc.IsTaxed(harga,pajak)
	fmt.Printf("harga  :%d\n",harga)
	fmt.Printf("voucher : %d\n",discount)
	fmt.Printf("biaya pengiriman : %d\n",delivery)
	fmt.Printf("pajak : %d\n",taxs)

}
func CalcFazzFood(){
	var harga int
	var voucher string
	var delivery int
	var taxs bool
	fmt.Print("harga :")
	fmt.Scanln(&harga)
	fmt.Print("voucher :")
	fmt.Scanln( &voucher)
	fmt.Print("biaya pengiriman :")
	fmt.Scanln(&delivery)
	fmt.Print("pajak :")
	fmt.Scanln(&taxs)
	FazzFood(harga,voucher,delivery,taxs)
		
}
