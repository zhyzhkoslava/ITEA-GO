package questionproviders

import (
	"context"
	"packages/questions"

	"github.com/jackc/pgx/v5"
)

type DBQuestionProvider struct {
	conn *pgx.Conn
}

func NewDBQuestionProvider(conn *pgx.Conn) *DBQuestionProvider {
	return &DBQuestionProvider{conn: conn}
}

func InsertQuestions(conn *pgx.Conn, questionsToAdd []*questions.Question) error {
	ctx := context.Background()
	for _, q := range questionsToAdd {
		_, err := conn.Exec(ctx,
			"INSERT INTO questions (text, options, correct_option_index) VALUES ($1, $2, $3)",
			q.Text, q.Options, q.CorrectOptionIndex,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *DBQuestionProvider) GetAllQuestions() ([]*questions.Question, error) {
	rows, err := p.conn.Query(context.Background(), "SELECT id, text, options, correct_option_index FROM questions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var allQuestions []*questions.Question

	for rows.Next() {
		var id int
		var text string
		var options []string
		var correctOptionIndex int

		if err := rows.Scan(&id, &text, &options, &correctOptionIndex); err != nil {
			return nil, err
		}

		question := &questions.Question{
			Text:               text,
			Options:            options,
			CorrectOptionIndex: correctOptionIndex,
		}

		allQuestions = append(allQuestions, question)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return allQuestions, nil
}
