package hard

import (
    "bufio"
    "fmt"
    "strconv"
)

type Student struct {
    ID   string
    Name string
    Age  int
}

func studentInformation(r *bufio.Reader) {
    id, _ := getInput("Enter ID: ", r)
    name, _ := getInput("Enter Name: ", r)
    ageStr, _ := getInput("Enter Age: ", r)
    age, _ := strconv.Atoi(ageStr)
    s := Student{ID: id, Name: name, Age: age}
    fmt.Printf("Student: %+v\n", s)
}

func multipleStudents(r *bufio.Reader) {
    nStr, _ := getInput("How many students? ", r)
    n, _ := strconv.Atoi(nStr)
    students := make([]Student, 0, n)
    for i := 0; i < n; i++ {
        fmt.Printf("\nStudent %d:\n", i+1)
        id, _ := getInput("ID: ", r)
        name, _ := getInput("Name: ", r)
        ageStr, _ := getInput("Age: ", r)
        age, _ := strconv.Atoi(ageStr)
        students = append(students, Student{ID: id, Name: name, Age: age})
    }
    fmt.Println("Students:")
    for _, s := range students {
        fmt.Printf("%+v\n", s)
    }
}

func searchStudentByID(r *bufio.Reader) {
    nStr, _ := getInput("How many students to store? ", r)
    n, _ := strconv.Atoi(nStr)
    students := make([]Student, 0, n)
    for i := 0; i < n; i++ {
        id, _ := getInput(fmt.Sprintf("ID for student %d: ", i+1), r)
        name, _ := getInput("Name: ", r)
        ageStr, _ := getInput("Age: ", r)
        age, _ := strconv.Atoi(ageStr)
        students = append(students, Student{ID: id, Name: name, Age: age})
    }
    searchID, _ := getInput("Enter ID to search: ", r)
    for _, s := range students {
        if s.ID == searchID {
            fmt.Printf("Found: %+v\n", s)
            return
        }
    }
    fmt.Println("Student not found")
}

type GradeStudent struct {
    Name  string
    Grade float64
}

func highestGradeStudent(r *bufio.Reader) {
    nStr, _ := getInput("How many students? ", r)
    n, _ := strconv.Atoi(nStr)
    if n <= 0 {
        fmt.Println("No students")
        return
    }
    var best GradeStudent
    for i := 0; i < n; i++ {
        name, _ := getInput(fmt.Sprintf("Name for student %d: ", i+1), r)
        gStr, _ := getInput("Grade: ", r)
        g, _ := strconv.ParseFloat(gStr, 64)
        if i == 0 || g > best.Grade {
            best = GradeStudent{Name: name, Grade: g}
        }
    }
    fmt.Printf("Highest grade student: %s -> %.2f\n", best.Name, best.Grade)
}

type Employee struct {
    Name          string
    MonthlySalary float64
}

func employeeSalaryCalculator(r *bufio.Reader) {
    name, _ := getInput("Employee name: ", r)
    sStr, _ := getInput("Monthly salary: ", r)
    s, _ := strconv.ParseFloat(sStr, 64)
    emp := Employee{Name: name, MonthlySalary: s}
    yearly := emp.MonthlySalary * 12
    fmt.Printf("%s's yearly salary: %.2f\n", emp.Name, yearly)
}
