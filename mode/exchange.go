package mode

import (
	"fmt"
)

type ExchangeDraw struct {
	Nothing
	View
	Tab fmt.Stringer
}

func (e *ExchangeDraw) DoChar(r rune) {
	if s := GetMoveValue(r); s != 0 {
		e.View.SetCursor(e.View.Cursor + s)
	}
}

func (e *ExchangeDraw) Draw() {
	py := 0

	py = PrintLine(py, " === "+e.Tab.String()+" Manager ===")

	e.View.PrintList(py)
}

func (e *ExchangeDraw) DoEnter() {
	e.View.Lister.Exchange(e.View.Cursor, e.View.Check)
}

func (e *ExchangeDraw) GetListLength() int {
	return e.View.Lister.GetListLength()
}

func (e *ExchangeDraw) Mode() Mode {
	return EXCHANGE
}
