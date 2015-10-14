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

	fp.Seek(0, 0)

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

func containsStr(ary []string, val string) (int, bool) {
	for i, v := range ary {
		if v == val {
			return i, true
		}
	}

	return 0, false
}

func (tl *TodoList) GetLabels() []string {
	return tl.labels
}

func (tl *TodoList) AddLabel(label string) {
	if _, is := containsStr(tl.labels, label); !is {
		tl.labels = append(tl.labels, label)
	}
}

func (tl *TodoList) RemoveLavel(label string) {
	if i, is := containsStr(tl.labels, label); is {
		tl.labels = append(tl.labels[:i], tl.labels[i+1:]...)
	}
}

func (tl *TodoList) Save() {
	tl.file.Seek(0, 0)
	tl.file.Truncate(0)
	SaveCSV(append(tl.todos, tl.archs...), tl.file)
}

func (tl *TodoList) ChangeTitle(i int, t string, tab Tab) {
	ls := tl.getListInTab(tab)
	if len(ls)-1 < i {
		return
	}

	ls[i].title = t
}

func (tl *TodoList) GetTodoTitle(i int, tab Tab) string {
	ls := tl.getListInTab(tab)
	if len(ls)-1 < i {
		return ""
	}

	return ls[i].title
}

func (tl *TodoList) ChangeLabelName(i int, l string, tab Tab) {
	ls := tl.getListInTab(tab)
	if len(ls)-1 < i {
		return
	}

	ls[i].label = l
}

func (tl *TodoList) Delete(n int, tab Tab) {
	var ls *[]Todo
	switch tab {
	case TODO:
		ls = &tl.todos
	case ARCHIVE:
		ls = &tl.archs
	}
	if len(*ls)-1 < n {
		return
	}

	for i := n; i < len(*ls); i++ {
		(*ls)[i].setNumber(i)
	}
	*ls = append((*ls)[:n], (*ls)[n+1:]...)
}

func (tl *TodoList) AddTodo(t string) int {
	no := len(tl.todos) + 1
	tl.todos = append(tl.todos, Todo{
		no:    no,
		label: "",
		title: t,
	})
	return no - 1
}

func (tl *TodoList) MoveTodo(n int, t Tab) {
	from, to := &tl.todos, &tl.archs
	isArched := true

	switch t {
	case ARCHIVE:
		from, to = to, from
		isArched = !isArched
	}

	if len(*from)-1 < n {
		return
	}

	length := len(*to)
	(*from)[n].isArchived = isArched
	(*from)[n].setNumber(length + 1)
	*to = append(*to, (*from)[n])
	for i := n + 1; i < len(*from); i++ {
		(*from)[i].setNumber(i)
	}
	*from = append((*from)[:n], (*from)[n+1:]...)
}

func (tl *TodoList) Exchange(i1 int, i2 int, tab Tab) {
	ls := tl.getListInTab(tab)
	ls[i2].setNumber(i1 + 1)
	ls[i1].setNumber(i2 + 1)
	ls[i2], ls[i1] = ls[i1], ls[i2]
}

func (t *Todo) setNumber(n int) {
	(*t).no = n
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
