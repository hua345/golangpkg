package excel

import (
	"fmt"
	"ginExample/model"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
)

var DrawResultExcelV2 = "./upload/drawResultV2.xlsx"

func ParseExcelV2(excelPath string) []model.Company {
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
		if len(rows) <= 2 {
			continue
		}
		company := model.Company{}
		var typeRow []string
		for index, row := range rows {
			if index == 0 {
				if len(row[0]) != 0 {
					company.CompanyName = row[0]
				}
				continue
			}
			if index == 1 {
				typeRow = []string{}
				typeRow = append(typeRow, row[:]...)
				continue
			}
			for colIndex, colCell := range row {
				typeValue := typeRow[colIndex]
				if len(colCell) != 0 && len(typeValue) != 0 {
					user := model.User{
						Name:      colCell,
						Type:      typeValue,
						TypeIndex: colIndex,
					}
					company.UserList = append(company.UserList, user)
				}
			}
		}
		companyList = append(companyList, company)
	}
	return companyList
}
func CreateExcelV2(drawResult *model.DrawResult) {
	excelFile := excelize.NewFile()
	// Create a new sheet.
	excelFile.SetSheetName(excelFile.GetSheetName(1), drawResult.CompanyName)
	err := excelFile.SetCellValue(drawResult.CompanyName, "A1", drawResult.CompanyName)
	if err != nil {
		panic(err)
	}
	axisIndex := 1
	userTypeMap := drawResult.GetTypeList()
	axisStartName, err := excelize.ColumnNumberToName(axisIndex)
	if err != nil {
		panic(err)
	}
	axisEndName, err := excelize.ColumnNumberToName(len(userTypeMap))
	if err != nil {
		panic(err)
	}
	err = excelFile.MergeCell(drawResult.CompanyName, axisStartName+strconv.Itoa(1), axisEndName+strconv.Itoa(1))
	if err != nil {
		panic(err)
	}
	style, err := excelFile.NewStyle(`{"alignment":{"horizontal":"center"},"font":{"bold":true,"family":"宋体","size":16,"color":"#ecc63e"}}`)
	if err != nil {
		fmt.Println(err)
	}
	err = excelFile.SetCellStyle(drawResult.CompanyName, axisStartName+strconv.Itoa(1), axisEndName+strconv.Itoa(1), style)
	if err != nil {
		fmt.Println(err)
	}
	for userType, userList := range userTypeMap {
		fmt.Println(userType, userList)
		if len(userList) == 0 {
			continue
		}
		originIndex := userList[0].TypeIndex + 1
		axisName, err := excelize.ColumnNumberToName(originIndex)
		if err != nil {
			panic(err)
		}
		aAxis := axisName + strconv.Itoa(2)
		// Set value of a cell.
		err = excelFile.SetCellValue(drawResult.CompanyName, aAxis, userType)
		if err != nil {
			panic(err)
		}
		for rowIndex, userInfo := range userList {
			aAxis := axisName + strconv.Itoa(rowIndex+3)
			// Set value of a cell.
			err = excelFile.SetCellValue(drawResult.CompanyName, aAxis, userInfo.Name)
			if err != nil {
				panic(err)
			}
		}
	}
	// Save xlsx file by the given path.
	err = excelFile.SaveAs(DrawResultExcelV2)
	if err != nil {
		panic(err)
	}
}
