package excel

import (
	"fmt"
	"ginExample/model"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
)

var DrawResultExcel = "./upload/drawResult.xlsx"

func ParseExcel(excelPath string) []model.Company {
	excelFile, err := excelize.OpenFile(excelPath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var companyList []model.Company
	for _, value := range excelFile.GetSheetMap() {
		rows, err := excelFile.GetRows(value)
		if err != nil {
			panic(nil)
		}
		company := model.Company{
			CompanyName: value,
		}
		for index, row := range rows {
			if index == 0 {
				continue
			}
			if len(row[0]) != 0 && len(row[1]) != 0 {
				user := model.User{
					Name: row[0],
					Type: row[1],
				}
				company.UserList = append(company.UserList, user)
			}
		}
		companyList = append(companyList, company)
	}
	return companyList
}
func CreateExcel(drawResult *model.DrawResult) {
	excelFile := excelize.NewFile()
	// Create a new sheet.
	// Create a new sheet.
	excelFile.SetSheetName(excelFile.GetSheetName(1), drawResult.CompanyName)
	sheetMap := excelFile.GetSheetMap()
	for key, value := range sheetMap {
		fmt.Println(key, value)
	}
	err := excelFile.SetCellValue(drawResult.CompanyName, "A1", "姓名")
	if err != nil {
		panic(err)
	}
	err = excelFile.SetCellValue(drawResult.CompanyName, "B1", "职位")
	if err != nil {
		panic(err)
	}
	for index, value := range drawResult.UserList {
		aAxis := "A" + strconv.Itoa(index+2)
		// Set value of a cell.
		err = excelFile.SetCellValue(drawResult.CompanyName, aAxis, value.Name)
		if err != nil {
			panic(err)
		}
		bAxis := "B" + strconv.Itoa(index+2)
		// Set value of a cell.
		err = excelFile.SetCellValue(drawResult.CompanyName, bAxis, value.Type)
		if err != nil {
			panic(err)
		}
	}
	// Save xlsx file by the given path.
	err = excelFile.SaveAs(DrawResultExcel)
	if err != nil {
		panic(err)
	}
}
