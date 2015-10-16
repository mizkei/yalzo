package yalzo

type Tab int

const (
	INPROCESS Tab = iota
	ARCHIVE
	LABEL
)

func (t Tab) String() string {
	switch t {
	case INPROCESS:
		return "Inprocess"
	case ARCHIVE:
		return "Archive"
	case LABEL:
		return "Label"
	}

	return "Unknown"
}

type Data struct {
	InProcess *TodoList
	Archive   *TodoList
	Labels    *LabelList
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

	todo := from.Remove(i).(*Todo)
	no := to.Add(todo.title)
	to.ChangeLabel(no, todo.label)
}
