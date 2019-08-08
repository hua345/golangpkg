package excel

import (
	"ginExample/model"
	"testing"
)

func TestParseExcel(t *testing.T) {
	t.Log(ParseExcel("./test.xlsx"))
}
func TestCreateExcel(t *testing.T) {
	drawResult := &model.DrawResult{}
	drawResult.CompanyName = "A公司"
	drawResult.UserList = append(drawResult.UserList, model.User{Name: "A经理1", Type: "项目经理"})
	drawResult.UserList = append(drawResult.UserList, model.User{Name: "A前端5", Type: "前端"})
	drawResult.UserList = append(drawResult.UserList, model.User{Name: "A前端6", Type: "前端"})
	drawResult.UserList = append(drawResult.UserList, model.User{Name: "A后端1", Type: "后端"})
	drawResult.UserList = append(drawResult.UserList, model.User{Name: "A后端3", Type: "后端"})
	drawResult.UserList = append(drawResult.UserList, model.User{Name: "A后端5", Type: "后端"})

	CreateExcel(drawResult)
}
