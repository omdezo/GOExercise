package easy
import (
	"fmt"
)

func ElecBill(amount float64){
	if amount < 0 {
		fmt.Println("Please Enter Valid number")
		return
	} else if amount < 101 { 
		amount *= 0.5 
		fmt.Printf("You should pay %0.2f $", amount)
	} else if amount < 301 { 
		amount *= 0.75 
		fmt.Printf("You should pay %0.2f $", amount)
	} else if amount > 300 { 
		amount *= 1 
		fmt.Printf("You should pay %0.2f $", amount)
	}
}