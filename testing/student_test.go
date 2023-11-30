package testing

import (
	"fmt"
	"testing"
)

func TestCalculateAverage(t *testing.T) {
	testCases := []struct {
		name   string
		grades Grades
		result float64
		err    error
	}{
		{
			name: "Valid grades",
			grades: Grades{
				"Math":    90,
				"Science": 80,
				"History": 70,
			},
			result: 80,
			err:    nil,
		},
		{
			name:   "No grades available",
			grades: Grades{},
			result: 0.0,
			err:    fmt.Errorf("no grades available"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			student := Student{
				FirstName: "Test",
				LastName:  "Student",
				Grades:    tc.grades,
			}

			average, err := student.calculateAverage()

			if err != nil && err.Error() != tc.err.Error() {
				t.Errorf("Expected error %v, got %v", tc.err, err)
			}

			if average != tc.result {
				t.Errorf("Expected average %f, got %f", tc.result, average)
			}
		})
	}
}
