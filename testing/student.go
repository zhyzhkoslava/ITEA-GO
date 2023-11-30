package testing

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
