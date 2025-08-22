package studentgrading

// GetGrade returns the grade and status of a student
func GetGrade(marks int, firstAttempt bool) (string, string) {
	if marks >= 90 {
		if firstAttempt {
			return "A", "pass"
		}
		return "A", "passed"
	} else if marks >= 70 {
		if firstAttempt {
			return "B", "pass"
		}
		return "B", "passed"
	} else if marks >= 50 {
		if firstAttempt {
			return "C", "pass"
		}
		return "C", "passed"
	}
	return "Fail", "fail"
}

// Max returns the maximum of two numbers
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
