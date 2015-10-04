package yalzo

import (
	"os"
	"strconv"
)

type TodoList struct {
	todos  []Todo
	archs  []Todo
	labels []string
}

type Todo struct {
	label      string
	title      string
	isArchived bool
	no         int
}

type Tab int

const (
	TODO Tab = iota
	ARCHIVE
)

func (t Tab) String() string {
	switch t {
	case TODO:
		return "Todo"
	case ARCHIVE:
		return "Archive"
	}
	return "Unknown"
}

func NewTodoList(fp *os.File, ls []string) *TodoList {
	l, as, err := ReadCSV(fp)
	if err != nil {
		panic(err)
	}

	return &TodoList{
		todos:  l,
		archs:  as,
		labels: ls,
	}
}

func (tl *TodoList) GetList(width int, tab Tab) []string {
	lines := make([]string, 0, 100)
	lists := tl.getListInTab(tab)
	for i := 0; i < len(lists); i++ {
		lines = append(lines, lists[i].tolimitStr(width))
	}
	return lines
}

func (tl *TodoList) GetLabels() []string {
	for i := 0; i < len(tl.todos); i++ {
		label := tl.todos[i].label
		if !tl.existLabel(label) && label != "" {
			tl.labels = append(tl.labels, label)
		}
	}

	for i := 0; i < len(tl.archs); i++ {
		label := tl.archs[i].label
		if !tl.existLabel(label) && label != "" {
			tl.labels = append(tl.labels, label)
		}
	}

	return tl.labels
}

func (tl *TodoList) ChangeTitle(i int, t string, tab Tab) {
	switch tab {
	case TODO:
		(*tl).todos[i].title = t
	case ARCHIVE:
		(*tl).archs[i].title = t
	}
}

func (tl *TodoList) ChangeLabelName(i int, l string, tab Tab) {
	switch tab {
	case TODO:
		(*tl).todos[i].label = l
	case ARCHIVE:
		(*tl).archs[i].label = l
	}
}

func (tl *TodoList) Delete(n int) {
	for i := n; i < len((*tl).todos); i++ {
		(*tl).todos[i].setNumber(i)
	}
	tl.todos = append(tl.todos[:n], tl.todos[n+1:]...)
}

func (tl *TodoList) AddTodo(t string) {
	tl.todos = append(tl.todos, Todo{
		no:    len(tl.todos),
		label: "",
		title: t,
	})
}

func (tl *TodoList) MoveArchive(n int) {
	length := len(tl.todos)
	tl.todos[n].isArchived = true
	tl.todos[n].setNumber(length)
	tl.archs = append(tl.archs, tl.todos[n])
	for i := n + 1; i < len(tl.todos); i++ {
		tl.todos[i].setNumber(i - 1)
	}
	tl.todos = append(tl.todos[:n], tl.todos[n+1:]...)
}

func (tl *TodoList) MoveTodo(n int) {
	length := len(tl.todos)
	tl.archs[n].isArchived = false
	tl.archs[n].setNumber(length)
	tl.todos = append(tl.todos, tl.archs[n])
	for i := n + 1; i < length; i++ {
		tl.todos[i].setNumber(i - 1)
	}
	tl.archs = append(tl.archs[:n], tl.archs[n+1:]...)
}

func (tl *TodoList) Exchange(i1 int, i2 int) {
	(*tl).todos[i2].setNumber(i1 + 1)
	(*tl).todos[i1].setNumber(i2 + 1)

	(*tl).todos[i2], (*tl).todos[i1] = (*tl).todos[i1], (*tl).todos[i2]
}

func (t *Todo) setNumber(n int) {
	(*t).no = n
}

func (tl *TodoList) existLabel(l string) bool {
	for i := 0; i < len(tl.labels); i++ {
		if tl.labels[i] == l {
			return true
		}
	}
	return false
}

func (tl *TodoList) getListInTab(tab Tab) []Todo {
	switch tab {
	case ARCHIVE:
		return tl.archs
	case TODO:
		return tl.todos
	}
	return []Todo{}
}

func (t *Todo) tolimitStr(limit int) string {
	str := strconv.Itoa(t.no) + t.label + t.title
	length := len(str)
	if length > limit {
		return str[:limit]
	} else {
		return str
	}
}
