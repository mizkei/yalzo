package mode

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

const (
	INPUT_PREFIX = "> "
	colorDef     = termbox.ColorDefault
)

type Mode int

const (
	NORMAL Mode = iota
	INPUT
	LABELSET
	EXCHANGE
)

type Lister interface {
	GetListLength() int
	GetList(int) []string
	Add(string) int
	Remove(int) interface{}
	GetPresentName(int) string
	Rename(int, string)
	Exchange(int, int)
	ChangeLabel(int, string)
}

type View struct {
	Width    int
	Height   int
	Lister   Lister
	Input    *InputBox
	Cursor   int
	Check    int
	Selected []int
}

func (v *View) GetCursorIndex() int {
	return v.Cursor
}

func (v *View) GetSelectedIndex() []int {
	return v.Selected
}

func (v *View) SetCursor(i int) {
	if i < 0 {
		v.Cursor = 0
	} else if l := v.Lister.GetListLength() - 1; l < i {
		v.Cursor = l
	} else {
		v.Cursor = i
	}
}

func (v *View) Reset() {
	v.Cursor = 0
	v.Check = -1
	v.Selected = []int{}
	v.Input.DeleteAll()
}

func (v *View) SetLister(ls Lister) {
	v.Lister = ls
}

func (v *View) PrintList(y int) {
	var bgc termbox.Attribute

	for i, s := range v.Lister.GetList(v.Width) {
		bgc = termbox.ColorDefault
		if i == v.Cursor {
			bgc = termbox.ColorCyan
		} else if _, t := containsVal(v.Selected, i); t {
			bgc = termbox.ColorMagenta
		} else if i == v.Check {
			bgc = termbox.ColorGreen
		}

		PrintText(2, y, colorDef, bgc, s)
		y += 1
	}
}

func NewView(w, h int) *View {
	v := &View{
		Width:  w,
		Height: h,
		Input:  NewInputBox(INPUT_PREFIX),
	}
	v.Reset()

	return v
}

func NewViewWithCheck(w, h, i int) *View {
	v := NewView(w, h)
	v.Reset()
	v.Check = i

	return v
}

func PrintText(x, y int, fg, bg termbox.Attribute, text string) int {
	for _, c := range text {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}

	return x
}

func PrintLine(y int, str string) int {
	PrintText(0, y, colorDef, colorDef, str)

	return y + 1
}

func FillText(n, y int, fg, bg termbox.Attribute, c rune) {
	x := 0
	for i := 0; i < n; i += 1 {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
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

func ShiftIndex(ary *[]int, val int) {
	for i, v := range *ary {
		if v > val {
			(*ary)[i] -= 1
		}
	}
}

func GetMoveValue(r rune) int {
	switch r {
	case 'j':
		return 1
	case 'k':
		return -1
	case 'J':
		return 5
	case 'K':
		return -5
	case 'G':
		return 1000
	}

	return 0
}
