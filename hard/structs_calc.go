package hard

// Structs helpers: functions operating on Student/GradeStudent/Employee slices.

// CalcFindStudentByID searches slice for a student by ID and returns index or -1.
func CalcFindStudentByID(students []Student, id string) int {
	for i, s := range students {
		if s.ID == id {
			return i
		}
	}
	return -1
}

// CalcHighestGrade returns the index of the student with highest grade.
func CalcHighestGrade(grades []GradeStudent) int {
	if len(grades) == 0 {
		return -1
	}
	bestIdx := 0
	for i := 1; i < len(grades); i++ {
		if grades[i].Grade > grades[bestIdx].Grade {
			bestIdx = i
		}
	}
	return bestIdx
}

// CalcYearlySalary returns yearly salary for given monthly salary.
func CalcYearlySalary(monthly float64) float64 {
	return monthly * 12
}

