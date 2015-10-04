package yalzo

import (
	"github.com/nsf/termbox-go"
	"os"
)

type Mode int

const (
	NORMAL Mode = iota
	CHANGE
	INPUT
	LABEL
)

const (
	INPUT_PREFIX = "> "
)

// input
type InputRcvr struct {
	Nothing
	*InputBox
}

func (i *InputRcvr) DoKeyArrowLeft() {
	i.MoveCursorOneRuneBackward()
}

func (i *InputRcvr) DoKeyCtrlB() {
	i.MoveCursorOneRuneBackward()
}

func (i *InputRcvr) DoKeyArrowRight() {
	i.MoveCursorOneRuneForward()
}

func (i *InputRcvr) DoKeyCtrlF() {
	i.MoveCursorOneRuneForward()
}

func (i *InputRcvr) DoKeyBackspace() {
	i.DeleteRuneBackward()
}

func (i *InputRcvr) DoKeyDelete() {
	i.DeleteRuneForward()
}

func (i *InputRcvr) DoKeySpace() {
	i.InsertRune(' ')
}

func (i *InputRcvr) DoChar(r rune) {
	i.InsertRune(r)
}

// label
type LabelRcvr struct {
	Nothing
	Labels []string
	Cursor int
	Index  int
}

func (l *LabelRcvr) DoEnter() {
}

func (l *LabelRcvr) DoChar(rune) {
}

type view struct {
	Nothing
	Width, Height int
	KeyRcvr       Operator
	TodoList      *TodoList
	Mode          Mode
	Input         *InputBox
	Tab           Tab
	List          []string
	Cursor        int
	ExCheck       int
	Selected      []int
}

func (v *view) DoKeyEsc() {
	v.Mode = NORMAL
}

func (v *view) DoKeyArrowLeft() {
	v.KeyRcvr.DoKeyArrowLeft()
}

func (v *view) DoKeyCtrlB() {
	v.KeyRcvr.DoKeyCtrlB()
}

func (v *view) DoKeyArrowRight() {
	v.KeyRcvr.DoKeyArrowRight()
}

func (v *view) DoKeyCtrlF() {
	v.KeyRcvr.DoKeyCtrlF()
}

func (v *view) DoKeyBackspace() {
	v.KeyRcvr.DoKeyBackspace()
}

func (v *view) DoKeyDelete() {
	v.KeyRcvr.DoKeyDelete()
}

func (v *view) DoKeyTab() {
	if v.Mode == NORMAL {
		v.Cursor = 0
		v.Selected = []int{}
		if v.Tab == TODO {
			v.Tab = ARCHIVE
		} else {
			v.Tab = TODO
		}
		v.List = v.TodoList.GetList(v.Width, v.Tab)
	}
}

func (v *view) DoKeyCtrlX() {
	if v.Mode == NORMAL {
		v.ExCheck = v.Cursor
		v.Mode = CHANGE
	}
}

func (v *view) DoKeyCtrlW() {
	if v.Mode == NORMAL {
		v.KeyRcvr = &InputRcvr{
			InputBox: v.Input,
		}
		v.Mode = INPUT
	}
}

func (v *view) DoKeyCtrlL() {
	if v.Mode == NORMAL {
		v.KeyRcvr = &LabelRcvr{
			Labels: v.TodoList.GetLabels(),
			Index:  v.Cursor,
		}
		v.Mode = LABEL
	}
}

func (v *view) DoKeyCtrlD() {
	v.KeyRcvr.DoKeyCtrlD()
}

func (v *view) DoKeyCtrlA() {
	v.KeyRcvr.DoKeyCtrlA()
}

func (v *view) DoKeyCtrlR() {
	v.KeyRcvr.DoKeyCtrlR()
}

func (v *view) DoKeySpace() {
	if v.Mode == NORMAL {
		if i, b := containsVal(v.Selected, v.Cursor); b {
			v.Selected = append(v.Selected[:i], v.Selected[i+1:]...)
		} else {
			v.Selected = append(v.Selected, v.Cursor)
		}
		return
	}

	v.KeyRcvr.DoKeySpace()
}

func (v *view) DoEnter() {
	switch v.Mode {
	case INPUT:
		v.TodoList.AddTodo(string(v.Input.input))
		v.Input.DeleteAll()
		v.List = v.TodoList.GetList(v.Width, v.Tab)
		v.Mode = NORMAL
	case CHANGE:
		v.TodoList.Exchange(v.ExCheck, v.Cursor, v.Tab)
		v.List = v.TodoList.GetList(v.Width, v.Tab)
		v.Mode = NORMAL
	}
}

func (v *view) DoChar(r rune) {
	switch v.Mode {
	case NORMAL:
		if r == 'j' && v.Cursor < len(v.List)-1 {
			v.Cursor += 1
		} else if r == 'k' && 0 < v.Cursor {
			v.Cursor -= 1
		}
	case CHANGE:
		if r == 'j' && v.Cursor < len(v.List)-1 {
			v.Cursor += 1
		} else if r == 'k' && 0 < v.Cursor {
			v.Cursor -= 1
		}
	default:
		v.KeyRcvr.DoChar(r)
	}
}

func containsVal(ary []int, val int) (int, bool) {
	for i, v := range ary {
		if v == val {
			return i, true
		}
	}

	return 0, false
}

func (v *view) Draw() {
	colorDef := termbox.ColorDefault
	termbox.Clear(colorDef, colorDef)

	py := 0

	// input
	py = PrintLine(py, v.Input.GetInputString())
	termbox.SetCursor(v.Input.prefixWidth+v.Input.cursorVOffset, 0)

	// tab
	pX := 0
	for _, t := range []Tab{TODO, ARCHIVE} {
		if t == v.Tab {
			pX += PrintText(pX, 1, termbox.ColorBlack, termbox.ColorWhite, " "+t.String()+" ")
		} else {
			pX += PrintText(pX, 1, colorDef, colorDef, " "+t.String()+" ")
		}
	}
	py += 1

	// list
	for i, e := range v.List {
		if i == v.Cursor {
			PrintText(0, py, colorDef, termbox.ColorCyan, e)
		} else if _, t := containsVal(v.Selected, i); t && v.Mode == NORMAL {
			PrintText(0, py, colorDef, termbox.ColorMagenta, e)
		} else if i == v.ExCheck && v.Mode == CHANGE {
			PrintText(0, py, colorDef, termbox.ColorGreen, e)
		} else {
			PrintText(0, py, colorDef, colorDef, e)
		}
		py += 1
	}

	termbox.Flush()
}

func NewView(fp *os.File, labels []string) view {
	w, h := termbox.Size()

	view := &view{
		Width:    w,
		Height:   h,
		TodoList: NewTodoList(fp, labels),
		Mode:     NORMAL,
		Input:    &InputBox{prefix: INPUT_PREFIX, prefixWidth: len(INPUT_PREFIX)},
		Tab:      TODO,
		List:     nil,
		Cursor:   0,
		Selected: make([]int, 0, 20),
	}

	view.List = view.TodoList.GetList(view.Width, view.Tab)

	return *view
}
