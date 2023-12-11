package questionproviders

import (
	"encoding/json"

	"github.com/zhyzhkoslava/ITEA-GO/packages/model"
)

type JSONQuestionProvider struct {
	questions []*model.Question
}

func NewJSONQuestionProvider(jsonData []byte) (*JSONQuestionProvider, error) {
	var questions []*model.Question
	err := json.Unmarshal(jsonData, &questions)
	if err != nil {
		return nil, err
	}

	return &JSONQuestionProvider{questions: questions}, nil
}

func (p *JSONQuestionProvider) GetAllQuestions() ([]*model.Question, error) {
	return p.questions, nil
}
