package easy

import "fmt"

func Calcul(num1 , num2 float64, opration string){
	switch opration {
	case "+":
		fmt.Printf("%0.2f + %0.2f = %0.2f ",num1,num2, num1 + num2)
	case "-":
		fmt.Printf("%0.2f - %0.2f = %0.2f ",num1,num2, num1 - num2)
	case "*":
		fmt.Printf("%0.2f * %0.2f = %0.2f ",num1,num2, num1 * num2)
	case "/":
		fmt.Printf("%0.2f / %0.2f = %0.2f ",num1,num2, num1 / num2)
	}
}