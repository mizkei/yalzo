package mode

import (
	"github.com/nsf/termbox-go"
)

type Action int

const (
	RENAME Action = iota
	INSERT
)

type InputDraw struct {
	Nothing
	View
	Act Action
}

func (i *InputDraw) DoKeyArrowLeft() {
	i.View.Input.MoveCursorOneRuneBackward()
}

func (i *InputDraw) DoKeyCtrlB() {
	i.View.Input.MoveCursorOneRuneBackward()
}

func (i *InputDraw) DoKeyArrowRight() {
	i.View.Input.MoveCursorOneRuneForward()
}

func (i *InputDraw) DoKeyCtrlF() {
	i.View.Input.MoveCursorOneRuneForward()
}

func (i *InputDraw) DoKeyBackspace() {
	i.View.Input.DeleteRuneBackward()
}

func (i *InputDraw) DoKeyDelete() {
	i.View.Input.DeleteRuneForward()
}

func (i *InputDraw) DoKeySpace() {
	i.View.Input.InsertRune(' ')
}

func (i *InputDraw) DoChar(r rune) {
	i.View.Input.InsertRune(r)
}

func (i *InputDraw) DoEnter() {
	switch i.Act {
	case INSERT:
		no := i.View.Lister.Add(string(i.View.Input.input))
		i.View.Cursor = no
	case RENAME:
		i.View.Lister.Rename(i.View.Cursor, string(i.View.Input.input))
	}
}

func (i *InputDraw) Draw() {
	switch i.Act {
	case INSERT:
		i.View.PrintList(2)
	case RENAME:
		PrintLine(1, "(old) > "+i.View.Lister.GetPresentName(i.View.Cursor))
	}

	PrintLine(0, i.View.Input.GetInputString())
	termbox.SetCursor(i.View.Input.prefixWidth+i.View.Input.cursorVOffset, 0)
}

func (i *InputDraw) GetListLength() int {
	return i.View.Lister.GetListLength()
}

func (i *InputDraw) Mode() Mode {
	return INPUT
}
