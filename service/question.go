package service

import (
	"errors"
	"fmt"
	"rankCompanies/domain"
	"strconv"
	"strings"
)

type IQuestionService interface {
	GetQuestion(questionsString []string) ([]domain.Question, error)
	Teste(i string)
}
type QuestionService struct {
}

func (qs *QuestionService) GetQuestion(questionsString []string) ([]domain.Question, error) {
	var questions []domain.Question
	for _, questionString := range questionsString {
		questionSplit := strings.Split(questionString, " ")
		id, answer, err := convertQuestionString(questionSplit)
		if err != nil {
			return nil, err
		}
		if answer < 5 {
			var question domain.Question
			question.Id = id
			question.Answers = append(question.Answers, answer)
			index, err := seachQuestion(&question, &questions)
			if err != nil {
				questions = append(questions, question)
			} else {
				questions[index].Answers = append(questions[index].Answers, answer)
			}
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

func seachQuestion(question *domain.Question, questions *[]domain.Question) (int, error) {

	for i, x := range *questions {
		if x.Id == question.Id {
			return i, nil
		}
	}
	return 0, errors.New("Not Found")
}

func (qs *QuestionService) Teste(i string) {
	fmt.Println(i)
}
