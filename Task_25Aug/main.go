package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Roll  int
	Marks map[string]int
	Avg   float64
	Grade string
}

func (s *Student) CalculateResult() {
	total := 0
	for _, mark := range s.Marks {
		total += mark
	}
	s.Avg = float64(total) / float64(len(s.Marks))

	if s.Avg <= 50 {
		s.Grade = "Fail"
	} else if s.Avg > 50 && s.Avg < 70 {
		s.Grade = "B"
	} else {
		s.Grade = "A"
	}
}

func main() {

	students := []Student{
		{Name: "Shubhi", Roll: 1, Marks: map[string]int{"Math": 80, "Science": 70, "English": 90}},
		{Name: "Somu", Roll: 2, Marks: map[string]int{"Math": 40, "Science": 55, "English": 50}},
		{Name: "Anu", Roll: 3, Marks: map[string]int{"Math": 60, "Science": 65, "English": 58}},
	}

	results := make(map[int]Student)

	for i := range students {
		students[i].CalculateResult()
		results[students[i].Roll] = students[i]
	}

	fmt.Printf("%-5s %-10s %-10s %-10s\n", "Roll", "Name", "Average", "Grade")
	fmt.Println("---------------------------------------------------")
	for _, s := range results {
		fmt.Printf("%-5d %-10s %-10.2f %-10s\n", s.Roll, s.Name, s.Avg, s.Grade)
	}
}
