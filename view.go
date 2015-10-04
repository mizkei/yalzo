package yalzo

type Mode int

const (
	NORMAL Mode = iota
	CHANGE
	INPUT
	LABEL
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
