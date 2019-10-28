package mapLearn

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"testing"
)

type User struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func TestMapGet(t *testing.T) {
	myMap := make(map[string]User)
	user1 := User{Name: "A经理1", Type: "项目经理"}
	myMap["name"] = user1
	t.Log(myMap["name"])
	t.Log("Len ", len(myMap))
	value, ok := myMap["aa"]
	if ok {
		t.Log(value)
	} else {
		t.Log("Key not exist")
	}
}
func TestMapChar(t *testing.T) {
	value, err := excelize.ColumnNumberToName(1)
	if err != nil {
		t.Error(err)
	}
	t.Log(value)
	value, err = excelize.ColumnNumberToName(111)
	if err != nil {
		t.Error(err)
	}
	t.Log(value)
}
