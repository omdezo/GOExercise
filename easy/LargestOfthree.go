package easy

import (
	"fmt"
)

func FindLargestThree(num1, num2, num3 int) {
	if num1 >= num2 && num1 >= num3 {
		fmt.Printf("The largest number is: %d\n", num1)
	} else if num2 >= num1 && num2 >= num3 {
		fmt.Printf("The largest number is: %d\n", num2)
	} else {
		fmt.Printf("The largest number is: %d\n", num3)
	}
}