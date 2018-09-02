package test

import (
	"io/ioutil"
	"rankCompany/service"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReadFile(t *testing.T) {
	companyFile, _ := ioutil.ReadFile("./file/filetest.txt")
	companyService := new(service.CompanyService)

	Convey("Verify getFile", t, func() {
		_, err := companyService.GetFile(companyFile)
		Convey("Verify if file is read", func() {
			So(err, ShouldEqual, nil)

		})
	})
}

func TestReadNameCompany(t *testing.T) {
	companyFile, _ := ioutil.ReadFile("./file/filetest.txt")
	companyService := new(service.CompanyService)
	company, _ := companyService.GetFile(companyFile)
	Convey("Check if name that company is correct", t, func() {
		So(company.Name, ShouldEqual, "teste")
	})
}
