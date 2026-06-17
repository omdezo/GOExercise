package hard

import "unicode"

func CalcCountVowels(s string) int {
    count := 0
    for _, ch := range s {
        switch unicode.ToLower(ch) {
        case 'a', 'e', 'i', 'o', 'u':
            count++
        }
    }
    return count
}

func CalcReverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func CalcIsPalindrome(s string) bool {
    filtered := make([]rune, 0, len(s))
    for _, ch := range s {
        if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
            filtered = append(filtered, unicode.ToLower(ch))
        }
    }
    for i, j := 0, len(filtered)-1; i < j; i, j = i+1, j-1 {
        if filtered[i] != filtered[j] {
            return false
        }
    }
    return true
}

func CalcCharFreq(s string) map[rune]int {
    freq := make(map[rune]int)
    for _, ch := range s {
        freq[ch]++
    }
    return freq
}
