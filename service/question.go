package service

import (
	"errors"
	"rankCompanies/domain"
	"strconv"
	"strings"
)

type QuestionService struct {
}

func (qs *QuestionService) GetQuestion(questionsString []string) ([]domain.Question, error) {
	var questions []domain.Question
	for _, questionString := range questionsString {
		questionString := strings.Split(questionString, " ")
		id, answer, err := convertQuestionString(questionString)
		if err != nil {
			return nil, err
		}
		question := domain.Question{id, make([]int32, answer)}

		questionFound, err := seachQuestion(&question, &questions)

		if err != nil {
			questions = append(questions, question)
		} else {
			questionFound.Answers = append(questionFound.Answers, question.Answers[0])
		}

	}

	return questions, nil
}

func convertQuestionString(questionString []string) (int64, int32, error) {
	id, err := strconv.Atoi(questionString[0])
	if err != nil {
		return 0, 0, err
	}
	answers, err := strconv.Atoi(questionString[1])
	if err != nil {
		return 0, 0, err
	}
	return int64(id), int32(answers), nil
}

func seachQuestion(question *domain.Question, questions *[]domain.Question) (*domain.Question, error) {
	for _, x := range *questions {
		if x.Id == question.Id {
			return &x, nil
		}
	}
	return nil, errors.New("Not Found")
}
