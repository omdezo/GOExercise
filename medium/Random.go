package medium

import (
	"bufio"
	"fmt"
	"math/rand"
)

func Randomguess(reader *bufio.Reader) {
	secretNum := rand.Intn(100) + 1 // +1 makes it 1 to 100 instead of 0 to 99
	fmt.Println("I've chosen a number between 1 and 100. Try to guess it!")

	for {
		guess, err := InputGameGuess(reader)
		if err != nil {
			fmt.Println("Invalid input, try entering a number.")
			continue
		}

		
		if guess > secretNum {
			fmt.Println("Try lower DUDE")
		} else if guess < secretNum {
			fmt.Println("Try Higher COME ON")
		} else {
			fmt.Println(" You Win! ")
			break 
		}
	}
}