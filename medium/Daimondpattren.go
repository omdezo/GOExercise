package medium
import (
	"fmt"
)

func PrintDiamond(n int) {
    for i := 1; i <= n; i++ {
        for j := 1; j <= n - i; j++ {
            fmt.Print(" ")
        }
        for k := 1; k <= (2*i-1) ; k++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
    for i := n - 1; i >= 1; i-- {
        for j := 1; j <= n-i; j++ {
            fmt.Print(" ")
        }
        for k := 1; k <= (2*i - 1); k++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
}