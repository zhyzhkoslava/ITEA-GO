package main

import (
	"fmt"
)

type Grades struct {
	Math    int
	Science int
	History int
}

type Student struct {
	FirstName string
	LastName  string
	Grades    Grades
}

func (s Student) FullName() string {
	return s.FirstName + " " + s.LastName
}

func calculateAverage(grades Grades) float64 {
	total := grades.Math + grades.Science + grades.History
	count := 3

	if count == 0 {
		return 0.0
	}

	return float64(total) / float64(count)
}

func main() {
	students := []Student{
		{
			FirstName: "Sam",
			LastName:  "Smith",
			Grades: Grades{
				Math:    90,
				Science: 85,
				History: 78,
			},
		},
		{
			FirstName: "Julia",
			LastName:  "Roberts",
			Grades: Grades{
				Math:    95,
				Science: 88,
				History: 92,
			},
		},
		{
			FirstName: "John",
			LastName:  "Doe",
			Grades: Grades{
				Math:    95,
				Science: 88,
				History: 92,
			},
		},
	}

	fmt.Println("Список студентів та середній бал:")
	for _, student := range students {
		average := calculateAverage(student.Grades)
		fmt.Printf("%s: Середній бал - %.2f\n", student.FullName(), average)
	}
}
