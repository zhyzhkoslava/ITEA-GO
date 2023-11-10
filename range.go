package main

import "fmt"

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

	for key := range weekdays {
		fmt.Println(key)
	}
}
