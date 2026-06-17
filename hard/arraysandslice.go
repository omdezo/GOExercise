package hard

import (
    "bufio"
    "fmt"
    "strconv"
)

func readNInts(prompt string, r *bufio.Reader) ([]int, error) {
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

func initAndOutput(r *bufio.Reader) {
    fmt.Println("\nInitialization & Output")
    arr := [5]int{1, 2, 3, 4, 5}
    sl := []int{10, 20, 30, 40, 50}

    fmt.Println("Array values:")
    for i := 0; i < len(arr); i++ {
        fmt.Printf("arr[%d] = %d\n", i, arr[i])
    }

    fmt.Println("Slice values:")
    for i := range sl {
        fmt.Printf("sl[%d] = %d\n", i, sl[i])
    }
}

func maxMinInSlice(r *bufio.Reader) {
    fmt.Println("\nFind Maximum & Minimum in a Slice")
    nums, err := readNInts("", r)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    if len(nums) == 0 {
        fmt.Println("No numbers provided")
        return
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
    fmt.Printf("Maximum: %d\nMinimum: %d\n", max, min)
}

func countEvenOdd(r *bufio.Reader) {
    fmt.Println("\nCount Even & Odd Numbers")
    nums, err := readNInts("", r)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    even, odd := 0, 0
    for _, v := range nums {
        if v%2 == 0 {
            even++
        } else {
            odd++
        }
    }
    fmt.Printf("Even: %d\nOdd: %d\n", even, odd)
}

func reverseSlice(r *bufio.Reader) {
    fmt.Println("\nReverse a Slice")
    nums, err := readNInts("", r)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
        nums[i], nums[j] = nums[j], nums[i]
    }
    fmt.Println("Reversed:", nums)
}

func searchNumber(r *bufio.Reader) {
    fmt.Println("\nSearch for a Number in a Slice")
    nums, err := readNInts("", r)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    s, err := getInput("Enter number to search for: ", r)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    target, err := strconv.Atoi(s)
    if err != nil {
        fmt.Println("Invalid number")
        return
    }
    found := false
    for i, v := range nums {
        if v == target {
            fmt.Printf("Found %d at index %d\n", target, i)
            found = true
        }
    }
    if !found {
        fmt.Println("Number not found in slice")
    }
}

func sortSlice(r *bufio.Reader) {
    fmt.Println("\nSorting a Slice")
    nums, err := readNInts("", r)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    // sort without importing sort here; use simple sort
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] > nums[j] {
                nums[i], nums[j] = nums[j], nums[i]
            }
        }
    }
    fmt.Println("Ascending:", nums)
    // descending
    for i := 0; i < len(nums)/2; i++ {
        j := len(nums) - 1 - i
        nums[i], nums[j] = nums[j], nums[i]
    }
    fmt.Println("Descending:", nums)
}

func removeDuplicates(r *bufio.Reader) {
    fmt.Println("\nRemove Duplicates from a Slice")
    nums, err := readNInts("", r)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    seen := make(map[int]bool)
    unique := make([]int, 0, len(nums))
    for _, v := range nums {
        if !seen[v] {
            seen[v] = true
            unique = append(unique, v)
        }
    }
    fmt.Println("Unique numbers:", unique)
}

func secondLargest(r *bufio.Reader) {
    fmt.Println("\nSecond Largest Number")
    nums, err := readNInts("", r)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    if len(nums) < 2 {
        fmt.Println("Need at least two numbers")
        return
    }
    first, second := nums[0], nums[0]
    if first < nums[1] {
        first, second = nums[1], nums[0]
    } else {
        first, second = nums[0], nums[1]
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
        fmt.Println("No second largest (all values equal?)")
        return
    }
    fmt.Printf("Second largest: %d\n", second)
}
