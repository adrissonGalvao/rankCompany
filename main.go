package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"rankCompanies/domain"
	"rankCompanies/service"
)

var companyService = service.CompanyService{}

func main() {
	var companies []domain.Company
	for _, file := range os.Args[1:] {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		conpany, err := companyService.GetFile(content)

		if err != nil {
			fmt.Println(err.Error())
		}

		for i, v := range conpany.Questions {
			conpany.Questions[i].Favorable, conpany.Questions[i].Neltral, conpany.Questions[i].Unfavorable, conpany.Questions[i].Invalid = service.CountAnswers(v.Answers)
		}
		companies = append(companies, conpany)

	}

	fmt.Println("Summary by companies:\n")
	service.SummaryByCompanies(companies)
	fmt.Println("Fav answers by questions:\n")
	service.FavAnswersByQuestions(companies)
	fmt.Println("Valid answers:\n")
	service.ValidAnswersByCompany(companies)
	fmt.Println("Invalid answers:\n")
	service.InvalidAnswersByCompany(companies)
}
