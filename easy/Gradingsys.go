package easy

import "fmt"

func GradingSys(score int){
	if score > 100 || score < 0 {
		fmt.Println("THAT'S WRONG NUBER BOY")
	} else if score >= 90 {
		fmt.Println("Excelent You got an A")
	} else if score >= 80 {
		fmt.Println("verygood You got an B")
	} else if score >= 70 {
		fmt.Println("Good You got an C")
	} else if score >= 60 {
		fmt.Println("You need to wake up You got an D")
	} else if score < 60 {
		fmt.Println("Try harder You got an F")
	}
}