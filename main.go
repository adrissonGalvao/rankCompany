package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"rankCompany/domain"
	"rankCompany/service"
)

func main() {

	var companies []domain.Company
	var companyService = service.CompanyService{}

	for _, file := range os.Args[1:] {

		content, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		company, err := companyService.GetFile(content)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		for i, v := range company.Questions {
			company.Questions[i].Favorable, company.Questions[i].Neltral, company.Questions[i].Unfavorable, company.Questions[i].Invalid = service.CountAnswers(v.Answers)
		}
		companies = append(companies, company)
	}

	fmt.Println("Summary by companies:")
	service.SummaryByCompanies(companies)
	fmt.Println("Fav answers by questions:")
	service.FavAnswersByQuestions(companies)
	fmt.Println("Valid answers:")
	service.ValidAnswersByCompany(companies)
	fmt.Println("Invalid answers:")
	service.InvalidAnswersByCompany(companies)
}
