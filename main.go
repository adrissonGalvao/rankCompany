package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"rankCompanies/service"
)

var companyService = service.CompanyService{}

func main() {
	for _, file := range os.Args[1:] {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		conpany, err := companyService.GetFile(content)

		fmt.Println(conpany.Name)

		// for _, v := range conpany.Questions {
		// 	fmt.Printf("Id:%d and len.Answers: %d\n", v.Id, len(v.Answers))
		// }

		/*teste*/
		for _, v := range conpany.Questions {
			favorable, neutral, unfavorable := service.PercentageCalculation(v.Answers)
			fmt.Println(v.Id, ":", favorable, "%fav,", neutral, "% neultral,", unfavorable, "% unfav")
		}
	}
}
