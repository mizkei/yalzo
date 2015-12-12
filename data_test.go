package yalzo

import (
	"testing"
)

func TestString(t *testing.T) {
	data := map[string]Tab{
		"Inprocess": INPROCESS,
		"Archive":   ARCHIVE,
		"Label":     LABEL,
	}

	for k, v := range data {
		if v.String() != k {
			t.Errorf("got %v, expected %v", v.String(), k)
		}
	}
}

func TestMoveTodo(t *testing.T) {
	todoList := &TodoList{}
	i := todoList.Add("test")
	todoList.ChangeLabel(i, "testl")
	data := Data{
		InProcess: todoList,
		Archive:   &TodoList{},
		Labels:    &LabelList{},
	}

	data.MoveTodo(0, INPROCESS)

	if data.InProcess.GetListLength() != 0 || data.Archive.GetListLength() != 1 {
		t.Errorf("failed move todo")
	}

	data.MoveTodo(0, INPROCESS)
	if data.InProcess.GetListLength() != 0 || data.Archive.GetListLength() != 1 {
		t.Errorf("failed move nothing")
	}
}
