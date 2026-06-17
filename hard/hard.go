package hard

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

// HardPlaceholder launches the hard exercises interactive menu.
func HardPlaceholder() {
	reader := bufio.NewReader(os.Stdin)
	Hard(reader)
}

// Hard is the exported entry point that follows the same style as easy.Medium(reader)
func Hard(reader *bufio.Reader) {
	for {
		fmt.Println("\n--- Hard Exercises Menu ---")
		fmt.Println("1. Arrays & Slices")
		fmt.Println("2. Strings")
		fmt.Println("3. Maps")
		fmt.Println("4. Structs")
		fmt.Println("5. Exit Hard Menu")

		choice, err := getInput("Choose a category: ", reader)
		if err != nil {
			fmt.Println("Error reading choice:", err)
			continue
		}
		switch choice {
		case "1":
			arraysMenu(reader)
		case "2":
			stringsMenu(reader)
		case "3":
			mapsMenu(reader)
		case "4":
			structsMenu(reader)
		case "5":
			return
		default:
			fmt.Printf("\n'%s' is not a valid option. Please try again.\n", choice)
			continue
		}
	}
}

func arraysMenu(r *bufio.Reader) {
	for {
		fmt.Println("\n-- Arrays & Slices --")
		fmt.Println("1. Initialization & Output")
		fmt.Println("2. Find Max & Min in Slice")
		fmt.Println("3. Count Even & Odd Numbers")
		fmt.Println("4. Reverse a Slice")
		fmt.Println("5. Search for a Number in a Slice")
		fmt.Println("6. Sorting a Slice")
		fmt.Println("7. Remove Duplicates from a Slice")
		fmt.Println("8. Second Largest Number")
		fmt.Println("9. Back to Hard Menu")
		choice, _ := getInput("Choose an option: ", r)
		switch choice {
		case "1":
			initAndOutput(r)
			_, _ = getInput("\nPress Enter to continue...", r)
		case "2":
			maxMinInSlice(r)
			_, _ = getInput("\nPress Enter to continue...", r)
		case "3":
			countEvenOdd(r)
			_, _ = getInput("\nPress Enter to continue...", r)
		case "4":
			reverseSlice(r)
			_, _ = getInput("\nPress Enter to continue...", r)
		case "5":
			searchNumber(r)
			_, _ = getInput("\nPress Enter to continue...", r)
		case "6":
			sortSlice(r)
			_, _ = getInput("\nPress Enter to continue...", r)
		case "7":
			removeDuplicates(r)
			_, _ = getInput("\nPress Enter to continue...", r)
		case "8":
			secondLargest(r)
			_, _ = getInput("\nPress Enter to continue...", r)
		case "9":
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func stringsMenu(r *bufio.Reader) {
	for {
		fmt.Println("\n-- Strings --")
		fmt.Println("1. Count Vowels")
		fmt.Println("2. Reverse a String")
		fmt.Println("3. Palindrome Checker")
		fmt.Println("4. Count Character Frequency")
		fmt.Println("5. Back to Hard Menu")
		choice, _ := getInput("Choose an option: ", r)
		switch choice {
		case "1":
			countVowels(r)
			_, _ = getInput("\nPress Enter to continue...", r)
		case "2":
			reverseString(r)
			_, _ = getInput("\nPress Enter to continue...", r)
		case "3":
			palindromeChecker(r)
			_, _ = getInput("\nPress Enter to continue...", r)
		case "4":
			countCharFreq(r)
			_, _ = getInput("\nPress Enter to continue...", r)
		case "5":
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func mapsMenu(r *bufio.Reader) {
	for {
		fmt.Println("\n-- Maps --")
		fmt.Println("1. Student Grades Lookup")
		fmt.Println("2. Word Frequency Counter")
		fmt.Println("3. Phone Book System")
		fmt.Println("4. Inventory System")
		fmt.Println("5. Back to Hard Menu")
		choice, _ := getInput("Choose an option: ", r)
		switch choice {
		case "1":
			studentGradesLookup(r)
			_, _ = getInput("\nPress Enter to continue...", r)
		case "2":
			wordFreqCounter(r)
			_, _ = getInput("\nPress Enter to continue...", r)
		case "3":
			phoneBookSystem(r)
			_, _ = getInput("\nPress Enter to continue...", r)
		case "4":
			inventorySystem(r)
			_, _ = getInput("\nPress Enter to continue...", r)
		case "5":
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func structsMenu(r *bufio.Reader) {
	for {
		fmt.Println("\n-- Structs --")
		fmt.Println("1. Student Information")
		fmt.Println("2. Multiple Students")
		fmt.Println("3. Search Student by ID")
		fmt.Println("4. Highest Grade Student")
		fmt.Println("5. Employee Salary Calculator")
		fmt.Println("6. Back to Hard Menu")
		choice, _ := getInput("Choose an option: ", r)
		switch choice {
		case "1":
			studentInformation(r)
			_, _ = getInput("\nPress Enter to continue...", r)
		case "2":
			multipleStudents(r)
			_, _ = getInput("\nPress Enter to continue...", r)
		case "3":
			searchStudentByID(r)
			_, _ = getInput("\nPress Enter to continue...", r)
		case "4":
			highestGradeStudent(r)
			_, _ = getInput("\nPress Enter to continue...", r)
		case "5":
			employeeSalaryCalculator(r)
			_, _ = getInput("\nPress Enter to continue...", r)
		case "6":
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

