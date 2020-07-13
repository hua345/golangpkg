package data_structures

import (
	"container/list"
	"testing"
)

func TestLinkedList(t *testing.T) {
	myList := list.New()
	e4 := myList.PushBack(4)
	e1 := myList.PushFront(1)
	myList.InsertBefore(3, e4)
	myList.InsertAfter(2, e1)
	for element := myList.Front(); element != nil; element = element.Next() {
		t.Log(element.Value)
	}
}
