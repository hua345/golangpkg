package excel

import (
	"ginExample/model"
	"testing"
)

func TestParseExcelV2(t *testing.T) {
	t.Log(ParseExcelV2("./testData.xlsx"))
}
func TestCreateExcelV2(t *testing.T) {
	drawResult := &model.DrawResult{}
	drawResult.CompanyName = "A公司"
	drawResult.UserList = append(drawResult.UserList, model.User{Name: "A经理1", Type: "项目经理", TypeIndex: 0})
	drawResult.UserList = append(drawResult.UserList, model.User{Name: "A前端5", Type: "前端", TypeIndex: 2})
	drawResult.UserList = append(drawResult.UserList, model.User{Name: "A前端6", Type: "前端", TypeIndex: 2})
	drawResult.UserList = append(drawResult.UserList, model.User{Name: "A后端1", Type: "后端", TypeIndex: 1})
	drawResult.UserList = append(drawResult.UserList, model.User{Name: "A后端3", Type: "后端", TypeIndex: 1})
	drawResult.UserList = append(drawResult.UserList, model.User{Name: "A后端5", Type: "后端", TypeIndex: 1})

	CreateExcelV2(drawResult)
}
