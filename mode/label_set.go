package mode

import (
	"fmt"
)

type LabelSetDraw struct {
	Nothing
	View
	Tab fmt.Stringer
}

func (l *LabelSetDraw) DoChar(r rune) {
	if s := GetMoveValue(r); s != 0 {
		l.View.SetCursor(l.View.Cursor + s)
	}
}

func (l *LabelSetDraw) Draw() {
	py := 0

	py = PrintLine(py, " === "+l.Tab.String()+" Manager ===")

	l.View.PrintList(py)
}

func (l *LabelSetDraw) GetListLength() int {
	return l.View.Lister.GetListLength()
}

func (l *LabelSetDraw) Mode() Mode {
	return LABELSET
}
