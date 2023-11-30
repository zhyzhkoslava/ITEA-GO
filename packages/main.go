package main

import (
	"context"
	"fmt"

	"github.com/zhyzhkoslava/ITEA-GO/questionproviders"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	connConfig, err := pgx.ParseConfig("postgres://postgres:pass@localhost:5432/lessons")
	if err != nil {
		panic(err)
	}
	conn, err := pgx.ConnectConfig(ctx, connConfig)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	questionsProvider := questionproviders.NewDBQuestionProvider(conn)
	correctAnswers := 0
	incorrectAnswers := 0

	questions, err := questionsProvider.GetAllQuestions()
	if err != nil {
		panic(err)
	}

	for _, question := range questions {
		fmt.Println(question.Text)
		for i, option := range question.Options {
			fmt.Printf("%d. %s\n", i+1, option)
		}

		fmt.Print("Ваша відповідь (введіть номер варіанту): ")
		var userChoice int
		_, scanErr := fmt.Scan(&userChoice)

		if scanErr != nil {
			fmt.Println("Помилка при зчитуванні відповіді. Спробуйте ще раз.")
			continue
		}

		if userChoice == question.CorrectOptionIndex+1 {
			fmt.Println("Правильно!")
			correctAnswers++
		} else {
			correctAnswer := question.CorrectOptionIndex + 1
			fmt.Printf("Неправильно. Правильна відповідь: %d\n\n", correctAnswer)
			incorrectAnswers++
		}
	}

	fmt.Printf("Правильних відповідей: %d\n", correctAnswers)
	fmt.Printf("Неправильних відповідей: %d\n", incorrectAnswers)
}
