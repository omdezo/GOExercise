package hard

import (
    "bufio"
    "fmt"
    "strconv"
)

// CalcReadNInts reads N integers from the reader and returns them.
func CalcReadNInts(r *bufio.Reader) ([]int, error) {
    nStr, err := getInput("Enter how many numbers (N): ", r)
    if err != nil {
        return nil, err
    }
    n, err := strconv.Atoi(nStr)
    if err != nil || n < 0 {
        return nil, fmt.Errorf("invalid number")
    }
    nums := make([]int, 0, n)
    for i := 0; i < n; i++ {
        s, err := getInput(fmt.Sprintf("Enter number %d: ", i+1), r)
        if err != nil {
            return nil, err
        }
        v, err := strconv.Atoi(s)
        if err != nil {
            return nil, fmt.Errorf("invalid integer")
        }
        nums = append(nums, v)
    }
    return nums, nil
}

func CalcMaxMin(nums []int) (int, int) {
    if len(nums) == 0 {
        return 0, 0
    }
    max, min := nums[0], nums[0]
    for _, v := range nums[1:] {
        if v > max {
            max = v
        }
        if v < min {
            min = v
        }
    }
    return max, min
}

func CalcCountEvenOdd(nums []int) (int, int) {
    even, odd := 0, 0
    for _, v := range nums {
        if v%2 == 0 {
            even++
        } else {
            odd++
        }
    }
    return even, odd
}

func CalcReverse(nums []int) []int {
    out := make([]int, len(nums))
    copy(out, nums)
    for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
        out[i], out[j] = out[j], out[i]
    }
    return out
}

func CalcSearchIndices(nums []int, target int) []int {
    idx := []int{}
    for i, v := range nums {
        if v == target {
            idx = append(idx, i)
        }
    }
    return idx
}

func CalcSortAsc(nums []int) []int {
    out := make([]int, len(nums))
    copy(out, nums)
    for i := 0; i < len(out); i++ {
        for j := i + 1; j < len(out); j++ {
            if out[i] > out[j] {
                out[i], out[j] = out[j], out[i]
            }
        }
    }
    return out
}

func CalcRemoveDuplicates(nums []int) []int {
    seen := make(map[int]bool)
    unique := make([]int, 0, len(nums))
    for _, v := range nums {
        if !seen[v] {
            seen[v] = true
            unique = append(unique, v)
        }
    }
    return unique
}

func CalcSecondLargest(nums []int) (int, error) {
    if len(nums) < 2 {
        return 0, fmt.Errorf("need at least two numbers")
    }
    first, second := nums[0], nums[0]
    if len(nums) >= 2 {
        if nums[0] < nums[1] {
            first, second = nums[1], nums[0]
        } else {
            first, second = nums[0], nums[1]
        }
    }
    for _, v := range nums[2:] {
        if v > first {
            second = first
            first = v
        } else if v > second && v < first {
            second = v
        }
    }
    if first == second {
        return 0, fmt.Errorf("no second largest (all values equal?)")
    }
    return second, nil
}
