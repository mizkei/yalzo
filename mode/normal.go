package mode

import (
	"fmt"
)

type NormalDraw struct {
	Nothing
	View
	Tab fmt.Stringer
}

func (n *NormalDraw) DoKeyCtrlD() {
	for _, i := range n.View.Selected {
		n.View.Lister.Remove(i)
		ShiftIndex(&n.View.Selected, i)
	}

	n.View.Reset()
}

func (n *NormalDraw) DoKeySpace() {
	if n.View.Lister.GetListLength() == 0 {
		return
	}

	if i, b := containsVal(n.View.Selected, n.View.Cursor); b {
		n.View.Selected = append(n.View.Selected[:i], n.View.Selected[i+1:]...)
	} else {
		n.View.Selected = append(n.View.Selected, n.View.Cursor)
	}
}

func (n *NormalDraw) DoChar(r rune) {
	if s := GetMoveValue(r); s != 0 {
		n.View.SetCursor(n.View.Cursor + s)
	}
}

func (n *NormalDraw) Draw() {
	py := 0

	py = PrintLine(py, " === "+n.Tab.String()+" Manager ===")

	n.View.PrintList(py)
}

func (n *NormalDraw) GetListLength() int {
	return n.View.Lister.GetListLength()
}

func (n *NormalDraw) Mode() Mode {
	return NORMAL
}

func (n *NormalDraw) ChangeLabel(i int, label string) {
	n.View.Lister.ChangeLabel(i, label)
}
