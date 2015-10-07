package yalzo

type Mode int

const (
	NORMAL Mode = iota
	CHANGE
	INPUT
	LABEL
	LABEL_DEL
	RENAME
)

type view struct {
	Nothing
	Width, Height int
	TodoList      *TodoList
	Mode          Mode
	Input         *InputBox
	Tab           Tab
	List          []string
	Cursor        int
	ExCheck       int
	Selected      []int
}

func (v *view) SetCursor(i int) {
	if i < 0 {
		v.Cursor = 0
	} else if l := len(v.List) - 1; l < i {
		v.Cursor = l
	} else {
		v.Cursor = i
	}
}
