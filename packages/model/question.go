package model

type Question struct {
	Text               string
	Options            []string
	CorrectOptionIndex int
}
