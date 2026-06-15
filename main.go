package main

import (
	"bufio"
	"exercise-go/easy"
	"exercise-go/medium"
	"exercise-go/hard"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}


func runMediumMenu(reader *bufio.Reader) {
	fmt.Println("\n -  Medium Exercises Menu -")
	medium.Medium(reader) 
	getInput("\nPress Enter to return to the main menu...", reader)
}
func runHardMenu(reader *bufio.Reader) {
	fmt.Println("\n--- Hard Exercises Menu ---")
	hard.HardPlaceholder() 
	getInput("\nPress Enter to return to the main menu...", reader)
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n CHOOSE YOUR DIFFICULTY ")
		fmt.Println("1. Easy Exercises")
		fmt.Println("2. Medium Exercises")
		fmt.Println("3. Hard Exercises")
		fmt.Println("4. Shut Down Program")

		difficulty, err := getInput("Select an option (1-4): ", reader)
		if err != nil {
			fmt.Println("Error reading input.")
			continue
		}

		if difficulty == "4" {
			fmt.Println("Shutting down CLI app. Goodbye!")
			break
		}

		switch difficulty {
		case "1":
			easy.Easy(reader) 
		case "2":
			runMediumMenu(reader)
		case "3":
			runHardMenu(reader)
		default:
			fmt.Printf("\n'%s' is not a valid tier. Try again.\n", difficulty)
		}
	}
}