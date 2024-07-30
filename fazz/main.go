package fazz

import (
	"fazztrack/demo/calc"
	"fmt"
)

func Fazz(harga int,voucher string,jarak int,pajak bool){
	discount := calc.Dicount(harga ,voucher)
	delivery := calc.DeliveryFee(jarak)
	taxs := calc.IsTaxed(harga,pajak)
	fmt.Printf("harga  :%d\n",harga)
	fmt.Printf("voucher : %d\n",discount)
	fmt.Printf("biaya pengiriman : %d\n",delivery)
	fmt.Printf("pajak : %d\n",taxs)
	
}