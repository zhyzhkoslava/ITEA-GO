package main

import "fmt"

type DayOfWeek int

const (
	Sunday    DayOfWeek = iota // 0
	Monday                     // 1
	Tuesday                    // 2
	Wednesday                  // 3
	Thursday                   // 4
	Friday                     // 5
	Saturday                   // 6
)

func main() {
	// 1. Створити константу з типом int і присвоїти їй результат арифметичної операції
	const result int = 10 + 5

	fmt.Println("Результат арифметичної операції:", result)

	// Приклад використання:
	today := Wednesday
	fmt.Printf("Сьогодні - %s (номер дня: %d)\n", DayOfWeekToString(today), today)
}

// Функція для перетворення значення DayOfWeek до рядка
func DayOfWeekToString(d DayOfWeek) string {
	days := []string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
	}
	return days[d]
}
