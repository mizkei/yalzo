package yalzo

import (
	"os"
)

type Tab int

const (
	INPROCESS Tab = iota
	ARCHIVE
)

func (t Tab) String() string {
	switch t {
	case INPROCESS:
		return "Todo"
	case ARCHIVE:
		return "Archive"
	}
	return "Unknown"
}

type Data struct {
	InProcess *TodoList
	Archive   *TodoList
	Labels    *LabelList
	File      *os.File
}

func (d *Data) MoveTodo(i int, tabFrom Tab) {
	var from, to *TodoList

	switch tabFrom {
	case INPROCESS:
		from, to = d.InProcess, d.Archive
	case ARCHIVE:
		from, to = d.Archive, d.InProcess
	}

	if !from.IsInAryRange(i) {
		return
	}

	todo := from.Remove(i)
	no := to.Add(todo.title)
	to.ChangeLabel(no, todo.label)
}
