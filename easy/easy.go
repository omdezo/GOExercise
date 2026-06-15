package easy

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

func InputOddEven(reader *bufio.Reader) (int, error) {
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

func InputLargestThree(reader *bufio.Reader) (int, int, int, error) {
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
func InputConverterCF(reader *bufio.Reader) (float64 ,error){
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
func InputDiscounter(reader *bufio.Reader) (float64 ,error){
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

func InputGradingSystem(reader *bufio.Reader) (int, error) {
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

func InputSwapped(reader *bufio.Reader) (int, error) {

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
func InputDays(reader *bufio.Reader) (int, error) {

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

func InputELec(reader *bufio.Reader) (float64, error) {

	in1, err := getInput("Enter number of Units: ", reader)
	if err != nil {
		return 0, err
	}
	n1, err := strconv.Atoi(in1)
	if err != nil {
		return 0, err
	}

	return float64(n1), nil
}

func InputCalculator(reader *bufio.Reader) (float64, float64, string, error) {

	input, err := getInput("Enter expression : ", reader)
	if err != nil {
		return 0, 0, "", err
	}

	
	parts := strings.Fields(input)
	
	if len(parts) != 3 {
		return 0, 0, "", fmt.Errorf("please use spaces (Format: number [space] operator [space] number)")
	}

	n1, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, 0, "", fmt.Errorf("'%s' is not a valid number", parts[0])
	}

	op := parts[1]

	n2, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return 0, 0, "", fmt.Errorf("'%s' is not a valid number", parts[2])
	}

	return n1, n2, op, nil
}

func Easy(reader *bufio.Reader) {

	for {
		fmt.Println("1. Even or Odd")
		fmt.Println("2. Largest of Three")
		fmt.Println("3. Celsius Converter to Fahrenheit")
		fmt.Println("4. Simple Discounter Calculater")
		fmt.Println("5. Grading System")
		fmt.Println("6. Swap Two Numbers")
		fmt.Println("7. Days to Weeks Converter")
		fmt.Println("8. Electricity Bill Calculator")
		fmt.Println("9. Simple Calculator")
		fmt.Println("10. Exit Program")

		choice, err := getInput("Choose a program : ", reader)
		if err != nil {
			fmt.Println("Error reading choice:", err)
			continue 
		}

		
		if choice == "10" {
			return
		}

		
		switch choice {
		case "1":
			fmt.Println("\n Running: Even or Odd Program ")
			number, err := InputOddEven(reader)
			if err != nil {
				fmt.Println("Oops haha That wasn't a valid whole number.")
			} else {
				CheckEvenOrOdd(number)
			}

		case "2":
			fmt.Println("\n Running: Largest of Three Program ")
			num1, num2, num3, err := InputLargestThree(reader)
			if err != nil {
				fmt.Println("Oops haha One of your inputs wasn't a valid whole number.")
			} else {
				FindLargestThree(num1, num2, num3)
			}
		case "3":
			fmt.Println("\n Running: Celsius Converter to Fahrenheit ")
			C, err := InputConverterCF(reader)
			if err != nil {
				fmt.Println("Oops haha That wasn't a valid whole number.")
			} else {
				ConvTemp(C)
			}
		case "4":
			fmt.Println("\n Running: Simple Discounter Calculator ")
			score, err := InputDiscounter(reader)
			if err != nil {
				fmt.Println("Oops haha That wasn't a valid whole number.")
			} else {
				Discounter(score)
			}
		case "5":
			fmt.Println("\n Running: Grading System ")
			price, err := InputGradingSystem(reader)
			if err != nil {
				fmt.Println("Oops haha That wasn't a valid whole number.")
			} else {
				GradingSys(price)
			}		
		case "6":
			fmt.Println("\n Running: Swap two numbers ")
			number, err := InputSwapped(reader)
			if err != nil {
				fmt.Println("Oops haha That wasn't a valid number.")
			} else {
				Swaptwo(number)
			}		
		case "7":
			fmt.Println("\n Running: Days to weeks Converter")
			days, err := InputDays(reader)
			if err != nil {
				fmt.Println("Oops haha That wasn't a valid number.")
			} else {
				DaysToweeksconv(days)
			}
		case "8":
			fmt.Println("\n Running: Electricity Bill Calculator")
			amount, err := InputELec(reader)
			if err != nil {
				fmt.Println("Oops haha That wasn't a valid number.")
			} else {
				ElecBill(amount)
			}
		case "9":
			fmt.Println("\n Running: Simple Calculator")
			num1, num2, op, err := InputCalculator(reader)
			if err != nil {
				fmt.Println("Oops haha That wasn't a valid number.")
			} else {
				Calcul(num1, num2, op)
			}
		default:
			
			fmt.Printf("\n'%s' is not a valid option. Please try again.\n", choice)
			continue
		}

		
		again, _ := getInput("\nWould you like to return to the main menu? (y/n): ", reader)
		if strings.ToLower(again) != "y" && strings.ToLower(again) != "yes" {
			fmt.Println("Goodbye!")
			return
		}
	}
}