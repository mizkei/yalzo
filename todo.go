package yalzo

import (
	"bytes"
	"fmt"
	"github.com/mattn/go-runewidth"
	"os"
)

const (
	LABEL_TEXT_WIDTH = 16
)

type TodoList struct {
	todos  []Todo
	archs  []Todo
	labels []string
	file   *os.File
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
		file:   fp,
	}
}

func (tl *TodoList) GetList(width int, tab Tab) []string {
	lines := make([]string, 0, 100)
	lists := tl.getListInTab(tab)
	for i := 0; i < len(lists); i++ {
		lines = append(lines, lists[i].createTodoText(width))
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

func (tl *TodoList) Save() {
	SaveCSV(tl.todos, tl.archs, tl.file)
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

func (tl *TodoList) AddTodo(t string) int {
	no := len(tl.todos) + 1
	tl.todos = append(tl.todos, Todo{
		no:    no,
		label: "",
		title: t,
	})
	return no
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

func (tl *TodoList) Exchange(i1 int, i2 int, tab Tab) {
	switch tab {
	case TODO:
		tl.todos[i2].setNumber(i1 + 1)
		tl.todos[i1].setNumber(i2 + 1)
		tl.todos[i2], tl.todos[i1] = tl.todos[i1], tl.todos[i2]
	case ARCHIVE:
		tl.archs[i2].setNumber(i1 + 1)
		tl.archs[i1].setNumber(i2 + 1)
		tl.archs[i2], tl.archs[i1] = tl.archs[i1], tl.archs[i2]
	}
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

func (t *Todo) createTodoText(limit int) string {
	num_s := fmt.Sprintf("%3d", t.no)
	label_s := t.label
	label_len := runewidth.StringWidth(label_s)
	if label_len > LABEL_TEXT_WIDTH {
		label_s = runewidth.Truncate(label_s, LABEL_TEXT_WIDTH, "")
	} else {
		padding := repeatSpace((LABEL_TEXT_WIDTH - label_len) / 2)
		buf := bytes.NewBuffer(padding)
		buf.Write([]byte(label_s))
		buf.Write(padding)
		label_s = buf.String()

		// if label string lenght is odd and more than LABEL_TEXT_WIDTH
		if runewidth.StringWidth(label_s) == LABEL_TEXT_WIDTH+1 {
			label_s = runewidth.Truncate(label_s, LABEL_TEXT_WIDTH, "")
		}
	}

	str := num_s + " [ " + label_s + " ] " + t.title
	length := runewidth.StringWidth(str)
	if length > limit {
		return runewidth.Truncate(str, limit, "")
	} else {
		str_len := runewidth.StringWidth(str)
		buf := bytes.NewBufferString(str)
		for str_len < limit {
			str_len++
			buf.Write([]byte(" "))
		}
		return buf.String()
	}
}

func repeatSpace(cnt int) []byte {
	buf := bytes.NewBuffer(make([]byte, 0, cnt))
	for cnt-1 >= 0 {
		buf.Write([]byte(" "))
		cnt--
	}
	return buf.Bytes()
}
