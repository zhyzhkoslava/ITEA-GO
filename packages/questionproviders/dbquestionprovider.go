package questionproviders

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/zhyzhkoslava/ITEA-GO/packages/model"
)

type DBQuestionProvider struct {
	conn *pgx.Conn
}

func NewDBQuestionProvider(conn *pgx.Conn) *DBQuestionProvider {
	return &DBQuestionProvider{conn: conn}
}

func (p *DBQuestionProvider) GetAllQuestions() ([]*model.Question, error) {
	rows, err := p.conn.Query(context.Background(), "SELECT id, text, options, correct_option_index FROM questions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var allQuestions []*model.Question

	for rows.Next() {
		var id int
		var text string
		var options []string
		var correctOptionIndex int

		if err := rows.Scan(&id, &text, &options, &correctOptionIndex); err != nil {
			return nil, err
		}

		question := &model.Question{
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
