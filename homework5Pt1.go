package main

import (
	"fmt"
)

type Grades map[string]int

type Student struct {
	FirstName string
	LastName  string
	Grades    Grades
}

func (s Student) FullName() string {
	return s.FirstName + " " + s.LastName
}

func (s Student) calculateAverage() (float64, error) {
	total := 0
	count := 0

	for _, grade := range s.Grades {
		total += grade
		count++
	}

	if count == 0 {
		return 0.0, fmt.Errorf("no grades available")
	}

	return float64(total) / float64(count), nil
}

func main() {
	students := []Student{
		{
			FirstName: "Sam",
			LastName:  "Smith",
			Grades: Grades{
				"Math":    90,
				"Science": 85,
				"History": 78,
			},
		},
		{
			FirstName: "Julia",
			LastName:  "Roberts",
			Grades: Grades{
				"Math":    95,
				"Science": 88,
				"History": 92,
			},
		},
		{
			FirstName: "John",
			LastName:  "Doe",
			Grades: Grades{
				"Math":    95,
				"Science": 88,
				"History": 92,
			},
		},
	}

	fmt.Println("Список студентів та середній бал:")
	for _, student := range students {
		average, err := student.calculateAverage()
		if err != nil {
			fmt.Printf("%s: %v\n", student.FullName(), err)
			continue
		}
		fmt.Printf("%s: Середній бал - %.2f\n", student.FullName(), average)
	}
}
