package hard

import (
    "bufio"
    "fmt"
    "strconv"
    "strings"
    "unicode"
)

func studentGradesLookup(r *bufio.Reader) {
    nStr, _ := getInput("How many students to enter? ", r)
    n, _ := strconv.Atoi(nStr)
    grades := make(map[string]string)
    for i := 0; i < n; i++ {
        name, _ := getInput(fmt.Sprintf("Enter student %d name: ", i+1), r)
        grade, _ := getInput("Enter grade: ", r)
        grades[name] = grade
    }
    name, _ := getInput("Enter a student's name to lookup: ", r)
    if g, ok := grades[name]; ok {
        fmt.Printf("%s -> %s\n", name, g)
    } else {
        fmt.Println("Student not found")
    }
}

func wordFreqCounter(r *bufio.Reader) {
    s, _ := getInput("Enter a sentence: ", r)
    s = strings.ToLower(s)
    // simple split on spaces and punctuation
    fields := strings.FieldsFunc(s, func(r rune) bool {
        return !unicode.IsLetter(r) && !unicode.IsNumber(r)
    })
    freq := make(map[string]int)
    for _, w := range fields {
        freq[w]++
    }
    fmt.Println("Word frequencies:")
    for k, v := range freq {
        fmt.Printf("%s: %d\n", k, v)
    }
}

func phoneBookSystem(r *bufio.Reader) {
    phonebook := make(map[string]string)
    for {
        fmt.Println("\nPhone Book Menu")
        fmt.Println("1. Add Contact")
        fmt.Println("2. Search Contact")
        fmt.Println("3. Delete Contact")
        fmt.Println("4. Display Contacts")
        fmt.Println("5. Exit Phone Book")
        choice, _ := getInput("Choose: ", r)
        switch choice {
        case "1":
            name, _ := getInput("Contact name: ", r)
            num, _ := getInput("Phone number: ", r)
            phonebook[name] = num
            fmt.Println("Contact added")
        case "2":
            name, _ := getInput("Name to search: ", r)
            if num, ok := phonebook[name]; ok {
                fmt.Printf("%s -> %s\n", name, num)
            } else {
                fmt.Println("Contact not found")
            }
        case "3":
            name, _ := getInput("Name to delete: ", r)
            delete(phonebook, name)
            fmt.Println("If existed, contact deleted")
        case "4":
            fmt.Println("Contacts:")
            for k, v := range phonebook {
                fmt.Printf("%s: %s\n", k, v)
            }
        case "5":
            return
        default:
            fmt.Println("Invalid choice")
        }
    }
}

func inventorySystem(r *bufio.Reader) {
    inventory := make(map[string]int)
    for {
        fmt.Println("\nInventory Menu")
        fmt.Println("1. Add Product")
        fmt.Println("2. Update Quantity")
        fmt.Println("3. Search Product")
        fmt.Println("4. Display Inventory")
        fmt.Println("5. Exit Inventory")
        choice, _ := getInput("Choose: ", r)
        switch choice {
        case "1":
            name, _ := getInput("Product name: ", r)
            qStr, _ := getInput("Quantity: ", r)
            q, _ := strconv.Atoi(qStr)
            inventory[name] = q
            fmt.Println("Product added")
        case "2":
            name, _ := getInput("Product name: ", r)
            qStr, _ := getInput("New quantity: ", r)
            q, _ := strconv.Atoi(qStr)
            inventory[name] = q
            fmt.Println("Quantity updated")
        case "3":
            name, _ := getInput("Product to search: ", r)
            if q, ok := inventory[name]; ok {
                fmt.Printf("%s -> %d\n", name, q)
            } else {
                fmt.Println("Product not found")
            }
        case "4":
            fmt.Println("Inventory:")
            for k, v := range inventory {
                fmt.Printf("%s: %d\n", k, v)
            }
        case "5":
            return
        default:
            fmt.Println("Invalid choice")
        }
    }
}
