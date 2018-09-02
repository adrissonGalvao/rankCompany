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

		conpany, err := companyService.GetFile(content)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		for i, v := range conpany.Questions {
			conpany.Questions[i].Favorable, conpany.Questions[i].Neltral, conpany.Questions[i].Unfavorable, conpany.Questions[i].Invalid = service.CountAnswers(v.Answers)
		}
		companies = append(companies, conpany)
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
