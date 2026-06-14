package main

import (
	"bufio"
	"exercise-go/easy"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func InputOddEven() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := getInput("Enter a number: ", reader)
	if err != nil {
		return 0, err
	}
	number, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func InputLargestThree() (int, int, int, error) {
	reader := bufio.NewReader(os.Stdin)

	in1, err := getInput("Enter first number: ", reader)
	if err != nil {
		return 0, 0, 0, err
	}
	n1, err := strconv.Atoi(in1)
	if err != nil {
		return 0, 0, 0, err
	}

	in2, err := getInput("Enter second number: ", reader)
	if err != nil {
		return 0, 0, 0, err
	}
	n2, err := strconv.Atoi(in2)
	if err != nil {
		return 0, 0, 0, err
	}

	in3, err := getInput("Enter third number: ", reader)
	if err != nil {
		return 0, 0, 0, err
	}
	n3, err := strconv.Atoi(in3)
	if err != nil {
		return 0, 0, 0, err
	}

	return n1, n2, n3, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("1. Even or Odd")
		fmt.Println("2. Largest of Three")
		fmt.Println("3. Exit Program")

		choice, err := getInput("Choose a program (1, 2, or 3): ", reader)
		if err != nil {
			fmt.Println("Error reading choice:", err)
			continue 
		}

		
		if choice == "3" {
			fmt.Println("Goodbye!")
			break 
		}

		
		switch choice {
		case "1":
			fmt.Println("\n Running: Even or Odd Program ")
			number, err := InputOddEven()
			if err != nil {
				fmt.Println("Oops! That wasn't a valid whole number.")
			} else {
				easy.CheckEvenOrOdd(number)
			}

		case "2":
			fmt.Println("\n Running: Largest of Three Program ")
			num1, num2, num3, err := InputLargestThree()
			if err != nil {
				fmt.Println("Oops! One of your inputs wasn't a valid whole number.")
			} else {
				easy.FindLargestThree(num1, num2, num3)
			}

		default:
			
			fmt.Printf("\n'%s' is not a valid option. Please try again.\n", choice)
			continue
		}

		
		again, _ := getInput("\nWould you like to return to the main menu? (y/n): ", reader)
		if strings.ToLower(again) != "y" && strings.ToLower(again) != "yes" {
			fmt.Println("Goodbye!")
			break 
		}
	}
}