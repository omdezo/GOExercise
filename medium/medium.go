package medium

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func InputAtm(reader *bufio.Reader) (float64, float64, string, error) {
	fmt.Println("\n - ATM -")
	fmt.Println("1. Deposit")
	fmt.Println("2. Withdraw")
	fmt.Println("3. Check Balance")

	choice, err := getInput("Select an option (1-3): ", reader)
	if err != nil {
		return 0, 0, "", err
	}

	var num1, num2 float64

	if choice == "1" {
		amtStr, err := getInput("Enter deposit amount: $", reader)
		if err != nil {
			return 0, 0, "", err
		}
		num1, err = strconv.ParseFloat(amtStr, 64)
		if err != nil {
			return 0, 0, "", fmt.Errorf("invalid amount entered")
		}
	} else if choice == "2" {
		amtStr, err := getInput("Enter withdrawal amount: $", reader)
		if err != nil {
			return 0, 0, "", err
		}
		num2, err = strconv.ParseFloat(amtStr, 64)
		if err != nil {
			return 0, 0, "", fmt.Errorf("invalid amount entered")
		}
	} else if choice != "3" {
		return 0, 0, "", fmt.Errorf("invalid action choice")
	}

	return num1, num2, choice, nil
}
func InputFact(reader *bufio.Reader) (int, error) {

	in1, err := getInput("Enter the number: ", reader)
	if err != nil {
		return 0, err
	}
	n1, err := strconv.Atoi(in1)
	if err != nil {
		return 0, err
	}

	return n1, nil
}
func InputSumEvenOdd(reader *bufio.Reader) (int, error) {

	in1, err := getInput("Enter a Number: ", reader)
	if err != nil {
		return 0, err
	}
	n1, err := strconv.Atoi(in1)
	if err != nil {
		return 0, err
	}

	return n1, nil
}
func Medium(reader *bufio.Reader) {
	for {
		fmt.Println("\n--- Medium Exercises Menu ---")
		fmt.Println("1. Basic ATM System")
		fmt.Println("2. Factorial of a Number")
		fmt.Println("3. Sum of Even and Odd Numbers")
		fmt.Println("4. Guess Game with helper")
		fmt.Println("5. Print Pattern")
		fmt.Println("6. Print a Pyramid Pattern")
		fmt.Println("7. Print a Diamond Pattern")
		fmt.Println("8. Exit Program")

		choice, err := getInput("Choose a program: ", reader)
		if err != nil {
			fmt.Println("Error reading choice:", err)
			continue
		}

		if choice == "8" {
			return 
		}

		switch choice {
		case "1":
			fmt.Println("\n=== Running: Basic ATM System ===")
			num1, num2, atmChoice, err := InputAtm(reader)
			if err != nil {
				fmt.Println("ATM Error:", err)
			} else {
			
				ATM(num1, num2, atmChoice)
			}

		case "2":
			fmt.Println("\n=== Running: Factorial of a Number ===")
			num, err := InputFact(reader)
			if err != nil {
				fmt.Println("Factorial Error:", err)
			} else {
				result := fact(num)
				fmt.Printf("The Factorial of %d is: %d\n", num, result)
			}

		case "3":
			fmt.Println("\n=== Running: Factorial of a Number ===")
			num, err := InputSumEvenOdd(reader)
			if err != nil {
				fmt.Println("Factorial Error:", err)
			} else {
				SumOddEven(num)
			}
			

		case "4":
			fmt.Println("\n Running: Guess Game with helper (Coming Soon)")

		case "5":
			fmt.Println("\n Running: Print Pattern (Coming Soon)")

		case "6":
			fmt.Println("\n Running: Print a Pyramid Pattern (Coming Soon)")

		case "7":
			fmt.Println("\n Running: Print a Diamond Pattern (Coming Soon)")

		default:
			fmt.Printf("\n'%s' is not a valid option. Please try again.\n", choice)
			continue
		}

		again, _ := getInput("\nWould you like to return to the medium menu? (y/n): ", reader)
		if strings.ToLower(again) != "y" && strings.ToLower(again) != "yes" {
			return
		}
	}
}