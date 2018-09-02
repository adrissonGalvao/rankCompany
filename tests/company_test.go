package test

import (
	"io/ioutil"
	"rankCompany/service"
	"testing"
)

func TestGetFile(t *testing.T) {
	content, _ := ioutil.ReadFile("./filetest.txt")
	var companyService = service.CompanyService{}
	company, err := companyService.GetFile(content)
	if err != nil {
		t.Error("error loading company")
	}
	if company.Name == "teste" {
		t.Log("right company:" + company.Name)
	}

}
