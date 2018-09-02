package service

import (
	"rankCompany/domain"
	"strings"
)

type CompanyService struct {
	QuestionService
}

func (cs *CompanyService) GetFile(content []byte) (domain.Company, error) {
	lines := strings.Split(string(content), "\n")

	var company domain.Company
	company.Name = lines[0]
	questions, err := cs.GetQuestion(lines[1:])

	if err != nil {
		return company, err
	}
	company.Questions = questions

	return company, nil
}
