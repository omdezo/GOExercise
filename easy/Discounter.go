package easy

import "fmt"

func Discounter(price float64){
	off := 0.1

	if price < 100 {
		fmt.Printf("No discount applied the final price is %0.2f", price)
	} else { 
		price = price + (price * off)
		fmt.Printf("10 precent discount applied the final price is %0.2f", price)
	}
}