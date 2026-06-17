package easy

import "fmt"

func ConvTemp(C float64) {
	C = (C*9 / 5)+32
	fmt.Printf("The temperature in Fahrenheit is %0.1f\n", C)
}