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

func (tl *TodoList) GetTodos(width int) []string {
	lines := make([]string, 0, 100)
	for i := 0; i < len(tl.todos); i++ {
		lines = append(lines, tl.todos[i].tolimitStr(width))
	}
	return lines
}

func (tl *TodoList) GetArchives(width int) []string {
	lines := make([]string, 0, 100)
	for i := 0; i < len(tl.archs); i++ {
		lines = append(lines, tl.archs[i].tolimitStr(width))
	}
	return lines
}

func (tl *TodoList) ChangeTitle(i int, t string, state Type) {
	if state == TODO {
		(*tl).todos[i].title = t
	} else {
		(*tl).archs[i].title = t
	}
}

func (tl *TodoList) ChangeLabelName(i int, l string, state Type) {
	if state == TODO {
		(*tl).todos[i].label = l
	} else {
		(*tl).archs[i].label = l
	}
}

func (tl *TodoList) Delete(n int) {
	for i := n; i < len((*tl).todos); i++ {
		(*tl).todos[i].setNumber(i)
	}
	tl.todos = append(tl.todos[:n], tl.todos[n+1:]...)
}

func (tl *TodoList) AddTodo(l string, t string) {
	tl.todos = append(tl.todos, Todo{
		no:    len(tl.todos),
		label: l,
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

func (t *Todo) tolimitStr(limit int) string {
	str := strconv.Itoa(t.no) + t.label + t.title
	length := len(str)
	if length > limit {
		return str[:limit]
	} else {
		return str
	}
}
