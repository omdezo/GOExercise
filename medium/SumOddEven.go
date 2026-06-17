package medium

import (
	"fmt"
)

func SumOddEven(n int) {
	
	if n < 1 {
		fmt.Println("Please enter a positive integer greater than 0")
		return
	}

	evenSum := 0
	oddSum := 0
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			evenSum += i
		} else {
			oddSum += i
		}
	}

	fmt.Printf("Sum of even numbers from 1 to %d is: %d\n", n, evenSum)
	fmt.Printf("Sum of odd numbers from 1 to %d is: %d\n", n, oddSum)
}