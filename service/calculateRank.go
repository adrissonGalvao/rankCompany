package service

import (
	"fmt"
	"rankCompany/domain"
	"strconv"
)

func CountAnswers(answers []int32) (int, int, int, int) {

	favorable := 0
	neutral := 0
	unfavorable := 0
	invalid := 0
	for _, answer := range answers {
		if answer == 0 || answer == 1 {
			favorable++
		} else if answer == 2 {
			neutral++
		} else if answer == 3 || answer == 4 {
			unfavorable++
		} else {
			invalid++
		}
	}

	return favorable, neutral, unfavorable, invalid

}

func SummaryByCompanies(companies []domain.Company) {
	for _, company := range companies {
		fmt.Println(company.Name)
		for _, question := range company.Questions {
			percente := CalculatePercente(question)
			fmt.Printf("%d: %d%% fav, %d%% neutral, %d%% unfav\n", question.Id, percente[0], percente[1], percente[2])
		}
	}
}

func CalculatePercente(question domain.Question) []int {

	total := float32(question.Favorable + question.Neltral + question.Unfavorable)
	var percFavorable float32
	var percNeltral float32
	var percUnfavorable float32
	if total != 0 {
		percFavorable = (float32(question.Favorable) / total) * 100
		percNeltral = (float32(question.Neltral) / total) * 100
		percUnfavorable = (float32(question.Unfavorable) / total) * 100
	} else {
		percFavorable = 0
		percNeltral = 0
		percUnfavorable = 0
	}
	percente := make([]int, 3)
	percente[0] = int(percFavorable)
	percente[1] = int(percNeltral)
	percente[2] = int(percUnfavorable)

	return percente
}

func FavAnswersByQuestions(companies []domain.Company) {
	favAnswers := make(map[int64]string)
	for _, v := range companies {
		for _, x := range v.Questions {
			percente := (CalculatePercente(x))[0]
			favAnswers[x.Id] += v.Name + " " + strconv.Itoa(percente) + "% fav"
		}
	}
	for i, v := range favAnswers {
		fmt.Printf("%d %s \n", i, v)
	}

}

func ValidAnswersByCompany(companies []domain.Company) {
	answersValid := make(map[string]int)
	for _, company := range companies {

		for _, question := range company.Questions {
			answersValid[company.Name] += question.Favorable + question.Neltral + question.Unfavorable
		}
	}
	for company, totalValid := range answersValid {
		fmt.Printf("%s: %d \n", company, totalValid)
	}
}

func InvalidAnswersByCompany(companies []domain.Company) {
	answersInvalid := make(map[string]int)
	for _, company := range companies {
		for _, question := range company.Questions {
			answersInvalid[company.Name] += question.Invalid
		}
	}

	for i, v := range answersInvalid {
		fmt.Printf("%s: %d \n", i, v)
	}
}
