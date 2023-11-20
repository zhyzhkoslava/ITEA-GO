package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Question struct {
	Text               string
	Options            []string
	CorrectOptionIndex int
}

type QuestionProvider interface {
	GetNextQuestion() *Question
}

type HardcodedQuestionProvider struct {
	questions  []*Question
	currentIdx int
}

func NewHardcodedQuestionProvider() *HardcodedQuestionProvider {
	questions := []*Question{
		{
			Text: "Яка столиця України?",
			Options: []string{
				"Київ",
				"Львів",
				"Одеса",
			},
			CorrectOptionIndex: 0,
		},
		{
			Text: "Скільки планет в Сонячній системі?",
			Options: []string{
				"7",
				"8",
				"9",
			},
			CorrectOptionIndex: 1,
		},
		{
			Text: "Як називається найбільший океан на Землі?",
			Options: []string{
				"Атлантичний",
				"Індійський",
				"Тихий",
			},
			CorrectOptionIndex: 2,
		},
	}

	return &HardcodedQuestionProvider{
		questions:  questions,
		currentIdx: 0,
	}
}

func (q *HardcodedQuestionProvider) GetNextQuestion() *Question {
	if q.currentIdx >= len(q.questions) {
		return nil
	}

	question := q.questions[q.currentIdx]
	q.currentIdx++
	return question
}

func main() {
	questionProvider := NewHardcodedQuestionProvider()
	correctAnswers := 0
	incorrectAnswers := 0

	for {
		question := questionProvider.GetNextQuestion()

		if question == nil {
			break
		}

		fmt.Println(question.Text)
		for i, option := range question.Options {
			fmt.Printf("%d. %s\n", i+1, option)
		}

		fmt.Print("Ваша відповідь (введіть номер варіанту): ")
		reader := bufio.NewReader(os.Stdin)
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		userChoice := parseUserChoice(userInput)
		if userChoice == question.CorrectOptionIndex {
			fmt.Println("Правильно!\n")
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

func parseUserChoice(input string) int {
	userChoice, err := strconv.Atoi(input)
	if err != nil {
		return -1
	}
	return userChoice - 1
}
