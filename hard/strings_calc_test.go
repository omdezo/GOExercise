package hard

import "testing"

func TestCalcCountVowels(t *testing.T) {
    got := CalcCountVowels("Hello World")
    if got != 3 { // e,o,o
        t.Fatalf("expected 3 vowels, got %d", got)
    }
}

func TestCalcIsPalindrome(t *testing.T) {
    if !CalcIsPalindrome("A man, a plan, a canal: Panama") {
        t.Fatalf("expected palindrome")
    }
    if CalcIsPalindrome("Not Palindrome") {
        t.Fatalf("expected not palindrome")
    }
}
