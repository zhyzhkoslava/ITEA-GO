package main

import (
	"fmt"
	"os"

	"github.com/zhyzhkoslava/ITEA-GO/packages/questionproviders"
)

func main() {
	jsonData, err := os.ReadFile("homework9/questions.json")
	if err != nil {
		panic(err)
	}

	jsonQuestionsProvider, err := questionproviders.NewJSONQuestionProvider(jsonData)
	if err != nil {
		panic(err)
	}

	questionsProvider := jsonQuestionsProvider

	correctAnswers := 0
	incorrectAnswers := 0

	questions, err := questionsProvider.GetAllQuestions()
	if err != nil {
		fmt.Println("Error:", err)
		return
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
