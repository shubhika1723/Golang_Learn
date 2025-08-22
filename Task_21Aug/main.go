package main

import (
	"Task_21Aug/studentgrading"
	"fmt"
)

func main() {
	// Example marks for 10 students
	marks := [10]int{95, 82, 67, 45, 30, 76, 54, 92, 49, 88}

	totalPass, totalFail := 0, 0

	for i, m := range marks {
		grade, status := studentgrading.GetGrade(m, true)

		if status == "fail" {
			// Failed → ask for re-exam
			var reMark int
			fmt.Printf("Enter re-exam marks for student %d (original %d): ", i+1, m)
			fmt.Scan(&reMark)

			final := studentgrading.Max(m, reMark)
			grade, status = studentgrading.GetGrade(final, false)
			fmt.Printf("Student %d → Final Marks: %d | Grade: %s | Status: %s\n",
				i+1, final, grade, status)

			if status == "fail" {
				totalFail++
			} else {
				totalPass++
			}
		} else {
			fmt.Printf("Student %d → Marks: %d | Grade: %s | Status: %s\n", i+1, m, grade, status)
			totalPass++
		}
	}

	fmt.Println("\nSummary:")
	fmt.Println("Total Passed:", totalPass)
	fmt.Println("Total Failed:", totalFail)
}
