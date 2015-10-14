package yalzo

import (
	"bytes"
	"fmt"

	"github.com/mattn/go-runewidth"
)

const (
	LABEL_TEXT_WIDTH = 16
)

func makeCenterText(text string, width int) string {
	ln := runewidth.StringWidth(text)
	if ln > width {
		return runewidth.Truncate(label_s, width, "")
	}

	padding := repeatSpace((width - ln) / 2)
	buf := bytes.NewBuffer(padding)
	buf.Write([]byte(text))
	buf.Write(padding)

	ctext := buf.String()

	// if label string length is odd and more than width
	if runewidth.StringWidth(ctext) == width+1 {
		return runewidth.Truncate(ctext, width, "")
	}

	return ctext
}

func repeatSpace(cnt int) []byte {
	buf := bytes.NewBuffer(make([]byte, 0, cnt))
	for 0 < cnt {
		buf.Write([]byte(" "))
		cnt--
	}

	return buf.Bytes()
}

func fillSpace(text string, limit int) string {
	ln := runewidth.StringWidth(text)
	if ln > limit {
		return runewidth.Truncate(text, limit, "")
	}

	buf := bytes.NewBufferString(text)
	for ; ln < limit; ln += 1 {
		buf.Write(' ')
	}
	return buf.String()
}

type Todo struct {
	label string
	title string
}

func (t *Todo) createTodoText(no, limit int) string {
	no_s := fmt.Sprintf("%3d", no)
	label_s := makeCenterText(t.label, LABEL_TEXT_WIDTH)

	str := no_s + " [ " + label_s + " ] " + t.title

	return fillSpace(str, limit)
}

type TodoList []Todo

func (tl *TodoList) IsInAryRange(i int) bool {
	if i < 0 || len(*tl)-1 < i {
		return false
	}

	return true
}

func (tl *TodoList) GetList(width int) []string {
	lines := make([]string, 0, 100)
	for i, v := range *tl {
		lines = append(lines, tl.createTodoText(i+1, width))
	}

	return lines
}

func (tl *TodoList) Rename(i int, title string) {
	if !tl.IsInAryRange(i) {
		return
	}

	tl[i].title = title
}

func (tl *TodoList) GetPresentName(i int) string {
	if !tl.IsInAryRange(i) {
		return ""
	}

	return tl[i].title
}

func (tl *TodoList) ChangeLabel(i int, name string) {
	if !tl.IsInAryRange(i) {
		return ""
	}

	tl[i].label = name
}

func (tl *TodoList) Remove(i int) Todo {
	if !tl.IsInAryRange(i) {
		return ""
	}

	delValue := (*tl)[i]
	*tl = append((*tl)[:i], (*tl)[i+1:]...)

	return delValue
}

func (tl *TodoList) Add(title string) int {
	*tl = append(*tl, Todo{
		label: "",
		title: title,
	})

	return len(*tl) - 1
}

func (tl *TodoList) Exchange(i, j int) {
	(*tl)[j], (*tl)[i] = (*tl)[i], (*tl)[j]
}
