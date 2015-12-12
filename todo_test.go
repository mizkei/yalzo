package yalzo

import (
	"testing"
)

func initTodo(l, t string) Todo {
	return Todo{label: l, title: t}
}

func initTodoList() TodoList {
	todoList := []Todo{}

	for _, s := range []string{"1", "2", "3", "4"} {
		todoList = append(todoList, initTodo("l"+s, "t"+s))
	}

	return todoList
}

func TestTodoGetListLength(t *testing.T) {
	ls := initTodoList()

	if ls.GetListLength() != 4 {
		t.Error("label: get list length")
	}
}

func TestTodoGetList(t *testing.T) {
	initList := initTodoList()
	ls := initList.GetList(27)
	expected := []string{
		"  1 [        l1        ] t1",
		"  2 [        l2        ] t2",
		"  3 [        l3        ] t3",
		"  4 [        l4        ] t4",
	}

	for i := range ls {
		if ls[i] != expected[i] {
			t.Errorf("got \"%v\", expected \"%v\"", ls[i], expected[i])
		}
	}
}

func TestTodoIsInAryRange(t *testing.T) {
	ls := initTodoList()

	if !ls.IsInAryRange(3) {
		t.Error("in range error")
	}

	if ls.IsInAryRange(4) {
		t.Error("in range error")
	}
}

func TestTodoAdd(t *testing.T) {
	ls := initTodoList()

	i := ls.Add("t5")
	if i != 4 {
		t.Errorf("got \"%v\", expected \"%v\"", i, 4)
	}

	if ls[i].title != "t5" || ls[i].label != "" {
		t.Error("add error")
	}
}

func TestTodoRemove(t *testing.T) {
	ls := initTodoList()

	todo := ls.Remove(0).(*Todo)
	if todo.label != "l1" {
		t.Error("error remove")
	}

	if len(ls) != 3 {
		t.Error("error remove")
	}
}

func TestTodoGetPresentName(t *testing.T) {
	ls := initTodoList()

	name := ls.GetPresentName(0)
	if name != "t1" {
		t.Errorf("got \"%v\", expected \"%v\"", name, "t1")
	}
}

func TestTodoRename(t *testing.T) {
	ls := initTodoList()

	ls.Rename(0, "t10")
	if ls[0].title != "t10" {
		t.Errorf("got \"%v\", expected \"%v\"", ls[0].title, "t10")
	}
}

func TestTodoExchange(t *testing.T) {
	ls := initTodoList()

	ls.Exchange(0, 3)
	if ls[0].label != "l4" || ls[3].label != "l1" {
		t.Error("exchange")
	}
}

func TestTodoChangeLabel(t *testing.T) {
	ls := initTodoList()

	ls.ChangeLabel(0, "l10")
	if ls[0].label != "l10" {
		t.Errorf("got \"%v\", expected \"%v\"", ls[0].label, "l10")
	}
}
