package main

import (
	"fmt"
)

type Student struct {
	FirstName string
	LastName  string
	Grades    map[string]int
}

func calculateAverage(grades map[string]int) float64 {
	total := 0
	count := 0

	for _, grade := range grades {
		total += grade
		count++
	}

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
			Grades: map[string]int{
				"Math":    90,
				"Science": 85,
				"History": 78,
			},
		},
		{
			FirstName: "Julia",
			LastName:  "Roberts",
			Grades: map[string]int{
				"Math":    95,
				"Science": 88,
				"History": 92,
			},
		},
		{
			FirstName: "John",
			LastName:  "Doe",
			Grades: map[string]int{
				"Math":    95,
				"Science": 88,
				"History": 92,
			},
		},
	}

	fmt.Println("Список студентів та середній бал:")
	for _, student := range students {
		average := calculateAverage(student.Grades)
		fmt.Printf("%s %s: Середній бал - %.2f\n", student.FirstName, student.LastName, average)
	}
}
