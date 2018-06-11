package test

import (
	"rankCompanies/domain"
	"rankCompanies/service"
	"strconv"
	"testing"
)

func TestCountAnswers(t *testing.T) {
	answers := []int32{1, 2, 4, 5, 1, 3}

	favorable, neutral, unfavorable, invalid := service.CountAnswers(answers)

	if favorable != 2 {
		t.Error("Miscalculation favorable" + strconv.Itoa(favorable))
	}
	if neutral != 1 {
		t.Error("Miscalculation neutral" + strconv.Itoa(neutral))
	}
	if unfavorable != 2 {
		t.Error("Miscalculation unfavorable" + strconv.Itoa(unfavorable))
	}

	if invalid != 1 {
		t.Error("Miscalculation invalid" + strconv.Itoa(invalid))
	}
}

func TestCalculatePercente(t *testing.T) {
	var question domain.Question
	question.Favorable = 3
	question.Neltral = 1
	question.Unfavorable = 4
	percente := service.CalculatePercente(question)

	if percente[0] != 37 {
		t.Error("Miscalculation fav:" + strconv.Itoa(percente[0]))
	}
	if percente[1] != 12 {
		t.Error("Miscalculation neltral:" + strconv.Itoa(percente[1]))
	}
	if percente[2] != 50 {
		t.Error("Miscalculation unfav:" + strconv.Itoa(percente[2]))
	}
}
