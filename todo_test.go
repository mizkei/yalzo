package yalzo

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

var init_archs = func() []Todo {
	todos := []Todo{
		Todo{
			label:      "Label3",
			title:      "Todo 3",
			isArchived: true,
			no:         1,
		},
		Todo{
			label:      "Label4",
			title:      "Todo 4",
			isArchived: true,
			no:         2,
		},
	}
	return todos
}

var init_todos = func() []Todo {
	todos := []Todo{
		Todo{
			label:      "Label1",
			title:      "Todo 1",
			isArchived: false,
			no:         1,
		},
		Todo{
			label:      "Label2",
			title:      "Todo 2",
			isArchived: false,
			no:         2,
		},
	}
	return todos
}

var init_labels = func() []string {
	labels := []string{
		"Label1",
		"Label2",
		"Label3",
		"Label4",
	}
	return labels
}

var init_todolist = func() *TodoList {
	fp, _ := ioutil.TempFile(os.TempDir(), "prefix")
	todolist := &TodoList{
		todos:  init_todos(),
		archs:  init_archs(),
		labels: init_labels(),
		file:   fp,
	}
	return todolist
}

func TestConstructorWithNewTodoList(t *testing.T) {
	f, _ := ioutil.TempFile(os.TempDir(), "tmpfile")
	f.WriteString("1,Label1,Todo 1,false\n")
	f.WriteString("2,Label2,Todo 2,false\n")

	ls := []string{"Label1", "Label2", "Label3", "Label4"}

	todolist := NewTodoList(f, ls)

	if todolist.labels[0] != "Label1" {
		t.Error("Cannot load labels.")
	} else {
		t.Log("Passed TestConstructorWithNewTodoList.")
	}
}

func TestGetList(t *testing.T) {
	todolist := init_todolist()
	tds := todolist.GetList(50, TODO)

	text := "  1 [      Label1      ] Todo 1                   "

	if tds[0] != text {
		t.Errorf("Not match text: %s", tds[0])
	} else {
		t.Log("Passed TestGetList.")
	}
}

func TestGetLabels(t *testing.T) {
	todolist := init_todolist()
	labels := init_labels()
	lbs := todolist.GetLabels()

	if !reflect.DeepEqual(labels, lbs) {
		t.Errorf("Not match labels: %s, %s", labels, lbs)
	} else {
		t.Log("Passed TestGetLabels.")
	}
}

func TestAddLabel(t *testing.T) {
	todolist := init_todolist()
	todolist.AddLabel("Label5")

	_, b := containsStr(todolist.labels, "Label5")

	if b {
		t.Log("Passed TestGetAddLabel.")
	} else {
		t.Errorf("Not match added label: %s", todolist.labels)
	}
}

func TestRemoveLabel(t *testing.T) {
	todolist := init_todolist()
	todolist.RemoveLavel("Label5")
	lbs := todolist.GetLabels()

	if lbs[len(lbs)-1] == "Label5" {
		t.Errorf("Not removed label: %s", lbs)
	} else {
		t.Log("Passed TestGetAddLabel.")
	}
}

func TestSave(t *testing.T) {
}

func TestChangeTitle(t *testing.T) {
	todolist := init_todolist()
	todolist.ChangeTitle(0, "Title Changed 1", TODO)

	if todolist.todos[0].title != "Title Changed 1" {
		t.Errorf("Not changed title: %s", todolist.todos[0].title)
	} else {
		t.Log("Passed TestChangeTitle.")
		todolist.ChangeTitle(0, "Todo 1", TODO)
	}

	todolist.ChangeTitle(0, "Title Changed 1", ARCHIVE)

	if todolist.archs[0].title != "Title Changed 1" {
		t.Errorf("Not changed title: %s", todolist.archs[0].title)
	} else {
		t.Log("Passed TestChangeTitle.")
		todolist.ChangeTitle(0, "Todo 3", ARCHIVE)
	}
}

func TestChangeLabelName(t *testing.T) {
	todolist := init_todolist()
	todolist.ChangeLabelName(0, "Label5", TODO)

	if todolist.todos[0].label != "Label5" {
		t.Errorf("Not changed label name: %s", todolist.todos[0].label)
	} else {
		t.Log("Passed TestChangeLabelName.")
		todolist.ChangeLabelName(0, "Label1", TODO)
	}
	todolist.ChangeLabelName(0, "Label5", ARCHIVE)

	if todolist.archs[0].label != "Label5" {
		t.Errorf("Not changed label name: %s", todolist.archs[0].label)
	} else {
		t.Log("Passed TestChangeLabelName.")
		todolist.ChangeLabelName(0, "Label3", ARCHIVE)
	}
}

func TestDelete(t *testing.T) {
	todolist := init_todolist()
	todos_size := len(todolist.todos)
	archs_size := len(todolist.archs)
	cmp_todo := todolist.todos[todos_size-1]
	cmp_arch := todolist.archs[archs_size-1]

	todolist.Delete(todos_size-1, TODO)

	if containsTodo(cmp_todo, todolist.todos) || todos_size == len(todolist.todos) {
		t.Errorf("Not deleted todo: %+v", todolist.todos)
		t.Errorf("                : %+v", cmp_todo)
	} else {
		t.Log("Passed TestDelete: deleted todo of TODO tab.")
	}

	todolist.Delete(len(todolist.archs)-1, ARCHIVE)

	if containsTodo(cmp_arch, todolist.archs) || archs_size == len(todolist.archs) {
		t.Errorf("Not deleted todo: %+v", todolist.archs)
		t.Errorf("                : %+v", cmp_todo)
	} else {
		t.Log("Passed TestDelete: deleted todo of ARCHIVE tab.")
	}
}

func TestAddTodo(t *testing.T) {
	todolist := init_todolist()
	tmp_todo := Todo{
		label:      "",
		title:      "Todo 2",
		isArchived: false,
		no:         3,
	}
	todolist.AddTodo("Todo 2")

	if !reflect.DeepEqual(todolist.todos[len(todolist.todos)-1], tmp_todo) {
		t.Errorf("Not added todo: %+v", todolist.todos)
	} else {
		t.Log("Passed TestAddTodo.")
	}
}

func TestMoveTodo(t *testing.T) {
	todolist := init_todolist()
	todolist.MoveTodo(0, TODO)
	tmp_arch := Todo{
		label:      "Label1",
		title:      "Todo 1",
		isArchived: true,
		no:         len(todolist.archs),
	}

	if !reflect.DeepEqual(todolist.archs[len(todolist.archs)-1], tmp_arch) {
		t.Errorf("Not moved todo to archive list: %+v", todolist.archs)
	} else {
		t.Log("Passed TestMoveTodo (todo moved to archive list).")
	}

	todolist.MoveTodo(len(todolist.archs)-1, ARCHIVE)
	tmp_todo := Todo{
		label:      "Label1",
		title:      "Todo 1",
		isArchived: false,
		no:         len(todolist.todos),
	}

	if !reflect.DeepEqual(todolist.todos[len(todolist.todos)-1], tmp_todo) {
		t.Errorf("Not moved archive to todo: %+v", todolist.todos)
	} else {
		t.Log("Passed TestMoveTodo (archive moved to todo list).")
	}
}

func TestExchange(t *testing.T) {
	todolist := init_todolist()
	todos := init_todos()
	archs := init_archs()
	todolist.Exchange(0, 1, TODO)

	if reflect.DeepEqual(todolist.todos[0], todos[1]) {
		t.Errorf("Not exchaged: %+v", todolist.todos)
	} else {
		t.Log("Passed TestExchange (exchaged todo list).")
		todolist.Exchange(0, 1, TODO)
	}

	todolist.Exchange(0, 1, ARCHIVE)

	if reflect.DeepEqual(todolist.archs[0], archs[1]) {
		t.Errorf("Not exchaged: %+v", todolist.archs)
	} else {
		t.Log("Passed TestExchange (exchaged arhive list).")
		todolist.Exchange(0, 1, ARCHIVE)
	}
}

// private functions

func containsTodo(todo Todo, todos []Todo) bool {
	for _, v := range todos {
		if reflect.DeepEqual(v, todo) {
			return true
		}
	}
	return false
}
