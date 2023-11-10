package main

import (
	"fmt"
	"time"
)

func main() {
	weekdays := map[string]int{
		"Mon": 1,
		"Tue": 2,
		"Wed": 3,
		"Thu": 4,
		"Fri": 5,
		"Sut": 6,
		"Sun": 7,
	}

	currentDay := time.Now().Weekday().String()[:3]
	currentDayNumber := weekdays[currentDay]

	fmt.Printf("Сьогодні - %s (номер дня: %d)\n", currentDay, currentDayNumber)

}
