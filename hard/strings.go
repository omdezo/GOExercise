package hard

import (
    "bufio"
    "fmt"
    "unicode"
)

func countVowels(r *bufio.Reader) {
    s, _ := getInput("Enter a string: ", r)
    count := 0
    for _, ch := range s {
        switch unicode.ToLower(ch) {
        case 'a', 'e', 'i', 'o', 'u':
            count++
        }
    }
    fmt.Printf("Vowel count: %d\n", count)
}

func reverseString(r *bufio.Reader) {
    s, _ := getInput("Enter a string: ", r)
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    fmt.Println("Reversed:", string(runes))
}

func palindromeChecker(r *bufio.Reader) {
    s, _ := getInput("Enter a string: ", r)
    filtered := make([]rune, 0, len(s))
    for _, ch := range s {
        if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
            filtered = append(filtered, unicode.ToLower(ch))
        }
    }
    isPal := true
    for i, j := 0, len(filtered)-1; i < j; i, j = i+1, j-1 {
        if filtered[i] != filtered[j] {
            isPal = false
            break
        }
    }
    if isPal {
        fmt.Println("The string is a palindrome")
    } else {
        fmt.Println("The string is NOT a palindrome")
    }
}

func countCharFreq(r *bufio.Reader) {
    s, _ := getInput("Enter a string: ", r)
    freq := make(map[rune]int)
    for _, ch := range s {
        freq[ch]++
    }
    fmt.Println("Character frequencies:")
    for k, v := range freq {
        fmt.Printf("'%c': %d\n", k, v)
    }
}
