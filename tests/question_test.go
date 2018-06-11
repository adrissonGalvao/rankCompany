package test

import (
	"io/ioutil"
	"rankCompanies/service"
	"strings"
	"testing"
)

func TestGetQuestion(t *testing.T) {
	content, _ := ioutil.ReadFile("./file/filetest.txt")
	lines := strings.Split(string(content), "\n")
	var companyService = service.CompanyService{}
	questions, err := companyService.GetQuestion(lines[1:])
	if err != nil {
		t.Error("error loading questions")
	}

	if len(questions) != 3 {
		t.Error("error in number questions")
	}
}
