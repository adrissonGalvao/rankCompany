package main

import (
	"fmt"
	"io/ioutil"
	"rankCompanies/service"
)

var companyService = service.CompanyService{}

func main() {

	content, err := ioutil.ReadFile("./files/hall_a.i._ltd._iCiSXO.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	teste, err := companyService.GetFile(content)

	fmt.Println(teste.Name)

}
