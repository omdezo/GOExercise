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
func InputConverterCF() (float64 ,error){
	reader := bufio.NewReader(os.Stdin)
	input, err := getInput("Enter a Temperature in Celsius: ", reader)
	if err != nil {
		return 0, err
	}
	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return float64(num), nil
}
func InputDiscounter() (float64 ,error){
	reader := bufio.NewReader(os.Stdin)
	input, err := getInput("Enter the price: ", reader)
	if err != nil {
		return 0, err
	}
	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return float64(num), nil
}

func InputGradingSystem() (int, error) {
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

func InputSwapped() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	in1, err := getInput("Enter number (2 Digits) to swap them: ", reader)
	if err != nil {
		return 0, err
	}
	n1, err := strconv.Atoi(in1)
	if err != nil {
		return 0, err
	}

	return n1, nil
}
func InputDays() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	in1, err := getInput("Enter number of days: ", reader)
	if err != nil {
		return 0, err
	}
	n1, err := strconv.Atoi(in1)
	if err != nil {
		return 0, err
	}

	return n1, nil
}
func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("1. Even or Odd")
		fmt.Println("2. Largest of Three")
		fmt.Println("3. Celsius Converter to Fahrenheit")
		fmt.Println("4. Simple Discounter Calculater")
		fmt.Println("5. Grading System")
		fmt.Println("6. Swap Two Numbers")
		fmt.Println("7. Days to Weeks Converter")
		fmt.Println("8. Exit Program")

		choice, err := getInput("Choose a program : ", reader)
		if err != nil {
			fmt.Println("Error reading choice:", err)
			continue 
		}

		
		if choice == "8" {
			fmt.Println("Goodbye!")
			break 
		}

		
		switch choice {
		case "1":
			fmt.Println("\n Running: Even or Odd Program ")
			number, err := InputOddEven()
			if err != nil {
				fmt.Println("Oops haha That wasn't a valid whole number.")
			} else {
				easy.CheckEvenOrOdd(number)
			}

		case "2":
			fmt.Println("\n Running: Largest of Three Program ")
			num1, num2, num3, err := InputLargestThree()
			if err != nil {
				fmt.Println("Oops haha One of your inputs wasn't a valid whole number.")
			} else {
				easy.FindLargestThree(num1, num2, num3)
			}
		case "3":
			fmt.Println("\n Running: Celsius Converter to Fahrenheit ")
			C, err := InputConverterCF()
			if err != nil {
				fmt.Println("Oops haha That wasn't a valid whole number.")
			} else {
				easy.ConvTemp(C)
			}
		case "4":
			fmt.Println("\n Running: Simple Discounter Calculator ")
			score, err := InputDiscounter()
			if err != nil {
				fmt.Println("Oops haha That wasn't a valid whole number.")
			} else {
				easy.Discounter(score)
			}
		case "5":
			fmt.Println("\n Running: Grading System ")
			price, err := InputGradingSystem()
			if err != nil {
				fmt.Println("Oops haha That wasn't a valid whole number.")
			} else {
				easy.GradingSys(price)
			}		
		case "6":
			fmt.Println("\n Running: Swap two numbers ")
			number, err := InputSwapped()
			if err != nil {
				fmt.Println("Oops haha That wasn't a valid number.")
			} else {
				easy.Swaptwo(number)
			}		
		case "7":
			fmt.Println("\n Running: Days to weeks Converter")
			days, err := InputDays()
			if err != nil {
				fmt.Println("Oops haha That wasn't a valid number.")
			} else {
				easy.DaysToweeksconv(days)
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