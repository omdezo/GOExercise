package hard

import (
    "bufio"
    "strconv"
)

// InputReadNInts reads N integers from reader and returns the slice.
func InputReadNInts(r *bufio.Reader) ([]int, error) {
    return CalcReadNInts(r)
}

// InputSearchTarget reads an integer target from reader.
func InputSearchTarget(r *bufio.Reader) (int, error) {
    s, err := getInput("Enter number to search for: ", r)
    if err != nil {
        return 0, err
    }
    v, err := strconv.Atoi(s)
    if err != nil {
        return 0, err
    }
    return v, nil
}

// InputStudentCount reads how many students to expect.
func InputStudentCount(r *bufio.Reader) (int, error) {
    s, err := getInput("How many students? ", r)
    if err != nil {
        return 0, err
    }
    v, err := strconv.Atoi(s)
    if err != nil {
        return 0, err
    }
    return v, nil
}

// InputString reads a line string.
func InputString(r *bufio.Reader, prompt string) (string, error) {
    return getInput(prompt, r)
}
