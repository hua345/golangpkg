package data_structures

import (
	"testing"
)

type User struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func TestMapGet(t *testing.T) {
	// 内置的make函数可以创建一个map
	myMap := make(map[string]int)
	// 也可以用map字面值的语法创建map，同时还可以指定一些最初的key/value
	ages := map[string]int{
		"fang": 26,
		"hua":  28,
	}
	t.Log(ages)
	myMap["fang"] = 26
	myMap["hua"] = 28
	t.Logf("fang:%d", myMap["fang"])
	t.Log("Len ", len(myMap))
	value, ok := myMap["aa"]
	if ok {
		t.Log(value)
	} else {
		t.Log("Key not exist")
	}
}
