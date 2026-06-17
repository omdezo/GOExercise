package easy

import "fmt"

func Swaptwo(number int){
	if number < 10 || number > 100 {
		fmt.Printf("Must enter only 2 digits")
		return
	}
	tens := number / 10
	ones := number %10
	rev := (ones*10) + tens
	fmt.Printf("original: %d --> Swapped: %d\n ",number, rev)
}