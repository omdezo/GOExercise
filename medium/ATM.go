package medium

import "fmt"

var balance float64 = 1000.0

func ATM(num1, num2 float64, choice string) {
	switch choice {
	case "1", "Deposit": 
		balance = balance + num1 
		fmt.Printf("Deposit successful! Your current balance is $%.2f\n", balance)
        
	case "2", "Withdraw":
		if balance < num2 { 
			fmt.Println("Sorry, insufficient balance!")
			return
		} else { 
			balance = balance - num2 
			fmt.Printf("Great! You withdrew $%.2f, now your balance is $%.2f\n", num2, balance)
		}

	case "3", "Check balance":
		fmt.Printf("Your current balance is $%.2f\n", balance)
    
	default:
		fmt.Printf("\n'%s' is not a valid option. Please try again.\n", choice)
	}
}