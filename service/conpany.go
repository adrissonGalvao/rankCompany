package service

import (
	"rankCompanies/domain"
	"strings"
)

type CompanyService struct{}

func (cs *CompanyService) GetFile(content []byte) (domain.Company, error) {
	lines := strings.Split(string(content), "\n")
	var company domain.Company
	company.Name = lines[0]

	return company, nil
}
