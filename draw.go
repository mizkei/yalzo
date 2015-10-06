package yalzo

import (
	"os"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

const (
	INPUT_PREFIX = "> "
	colorDef     = termbox.ColorDefault
)

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

func shiftIndex(ary *[]int, val int) {
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

type Drawer interface {
	Operator
	Draw()
}

// input
type InputDraw struct {
	Nothing
	view   *view
	action Mode
}

func (i *InputDraw) DoKeyArrowLeft() {
	i.view.Input.MoveCursorOneRuneBackward()
}

func (i *InputDraw) DoKeyCtrlB() {
	i.view.Input.MoveCursorOneRuneBackward()
}

func (i *InputDraw) DoKeyArrowRight() {
	i.view.Input.MoveCursorOneRuneForward()
}

func (i *InputDraw) DoKeyCtrlF() {
	i.view.Input.MoveCursorOneRuneForward()
}

func (i *InputDraw) DoKeyBackspace() {
	i.view.Input.DeleteRuneBackward()
}

func (i *InputDraw) DoKeyDelete() {
	i.view.Input.DeleteRuneForward()
}

func (i *InputDraw) DoKeySpace() {
	i.view.Input.InsertRune(' ')
}

func (i *InputDraw) DoChar(r rune) {
	i.view.Input.InsertRune(r)
}

func (i *InputDraw) DoEnter() {
	switch i.action {
	case NORMAL:
		no := i.view.TodoList.AddTodo(string(i.view.Input.input))
		i.view.Cursor = no
	case RENAME:
		i.view.TodoList.ChangeTitle(i.view.Cursor, string(i.view.Input.input), i.view.Tab)
	}
	i.view.Input.DeleteAll()
	i.view.List = i.view.TodoList.GetList(i.view.Width, i.view.Tab)
}

func (i *InputDraw) Draw() {
	switch i.action {
	case NORMAL:
		(&NormalDraw{view: i.view}).Draw()
	case RENAME:
		PrintLine(1, "(old) > "+i.view.List[i.view.Cursor])
	}

	FillText(i.view.Width, 0, colorDef, colorDef, ' ')
	PrintLine(0, i.view.Input.GetInputString())
	termbox.SetCursor(i.view.Input.prefixWidth+i.view.Input.cursorVOffset, 0)
}

//change
type ChangeDraw struct {
	Nothing
	view *view
}

func (c *ChangeDraw) DoEnter() {
	c.view.TodoList.Exchange(c.view.Cursor, c.view.ExCheck, c.view.Tab)
	c.view.List = c.view.TodoList.GetList(c.view.Width, c.view.Tab)
}

func (c *ChangeDraw) DoChar(r rune) {
	if s := GetMoveValue(r); s != 0 {
		c.view.SetCursor(c.view.Cursor + s)
	}
}

func (c *ChangeDraw) Draw() {
	py := 0

	// input
	py = PrintLine(py, " @ Change mode")

	// tab
	pX := 0
	for _, t := range []Tab{TODO, ARCHIVE} {
		if t == c.view.Tab {
			pX += PrintText(pX, 1, termbox.ColorBlack, termbox.ColorWhite, " "+t.String()+" ")
		} else {
			pX += PrintText(pX, 1, colorDef, colorDef, " "+t.String()+" ")
		}
	}
	py += 1

	// list
	for i, e := range c.view.List {
		if i == c.view.Cursor {
			PrintText(0, py, colorDef, termbox.ColorCyan, e)
		} else if i == c.view.ExCheck {
			PrintText(0, py, colorDef, termbox.ColorGreen, e)
		} else {
			PrintText(0, py, colorDef, colorDef, e)
		}
		py += 1
	}
}

// label
type LabelDraw struct {
	Nothing
	view   *view
	Labels []string
	Cursor int
}

func (l *LabelDraw) DoEnter() {
	l.view.TodoList.ChangeLabelName(l.view.Cursor, l.Labels[l.Cursor], l.view.Tab)
	l.view.List = l.view.TodoList.GetList(l.view.Width, l.view.Tab)
}

func (l *LabelDraw) DoChar(r rune) {
	if r == 'j' && l.Cursor < len(l.Labels)-1 {
		l.Cursor += 1
	} else if r == 'k' && 0 < l.Cursor {
		l.Cursor -= 1
	}
}

func (l *LabelDraw) Draw() {
	py := 0

	PrintLine(py, l.view.List[l.view.Cursor])
	py += 1

	for i, e := range l.Labels {
		PrintText(1, py, colorDef, colorDef, "*")
		if i == l.Cursor {
			PrintText(3, py, colorDef, termbox.ColorCyan, e)
		} else {
			PrintText(3, py, colorDef, colorDef, e)
		}
		py += 1
	}
}

// normal
type NormalDraw struct {
	Nothing
	view *view
}

func (n *NormalDraw) DoKeyTab() {
	n.view.Cursor = 0
	n.view.Selected = []int{}
	if n.view.Tab == TODO {
		n.view.Tab = ARCHIVE
	} else {
		n.view.Tab = TODO
	}
	n.view.List = n.view.TodoList.GetList(n.view.Width, n.view.Tab)
}

func (n *NormalDraw) DoKeyCtrlD() {
	for _, i := range n.view.Selected {
		n.view.TodoList.Delete(i)
		shiftIndex(&n.view.Selected, i)
	}
	n.view.Cursor = 0
	n.view.Selected = []int{}
	n.view.List = n.view.TodoList.GetList(n.view.Width, n.view.Tab)
}

func (n *NormalDraw) DoKeyCtrlA() {
	if len(n.view.Selected) == 0 {
		n.view.TodoList.MoveTodo(n.view.Cursor, n.view.Tab)
	} else {
		for _, i := range n.view.Selected {
			n.view.TodoList.MoveTodo(i, n.view.Tab)
			shiftIndex(&n.view.Selected, i)
		}
	}
	n.view.Cursor = 0
	n.view.Selected = []int{}
	n.view.List = n.view.TodoList.GetList(n.view.Width, n.view.Tab)
}

func (n *NormalDraw) DoKeySpace() {
	if i, b := containsVal(n.view.Selected, n.view.Cursor); b {
		n.view.Selected = append(n.view.Selected[:i], n.view.Selected[i+1:]...)
	} else {
		n.view.Selected = append(n.view.Selected, n.view.Cursor)
	}
}

func (n *NormalDraw) DoChar(r rune) {
	if s := GetMoveValue(r); s != 0 {
		n.view.SetCursor(n.view.Cursor + s)
	}
}

func (n *NormalDraw) Draw() {
	py := 0

	// input
	py = PrintLine(py, " === TODO Manager ===")

	// tab
	pX := 0
	for _, t := range []Tab{TODO, ARCHIVE} {
		if t == n.view.Tab {
			pX += PrintText(pX, 1, termbox.ColorBlack, termbox.ColorWhite, " "+t.String()+" ")
		} else {
			pX += PrintText(pX, 1, colorDef, colorDef, " "+t.String()+" ")
		}
	}
	py += 1

	// list
	for i, e := range n.view.List {
		if i == n.view.Cursor {
			PrintText(0, py, colorDef, termbox.ColorCyan, e)
		} else if _, t := containsVal(n.view.Selected, i); t {
			PrintText(0, py, colorDef, termbox.ColorMagenta, e)
		} else {
			PrintText(0, py, colorDef, colorDef, e)
		}
		py += 1
	}
}

type Draw struct {
	Drawer
	view *view
}

func (d *Draw) DoKeyEsc() {
	switch d.view.Mode {
	case INPUT, CHANGE, LABEL:
		d.view.Input.DeleteAll()
		d.view.Mode = NORMAL
		d.Drawer = &NormalDraw{view: d.view}
	}
}

func (d *Draw) DoKeyCtrlX() {
	switch d.view.Mode {
	case NORMAL:
		d.view.ExCheck = d.view.Cursor
		d.view.Mode = CHANGE
		d.Drawer = &ChangeDraw{view: d.view}
	}
}

func (d *Draw) DoKeyCtrlW() {
	switch d.view.Mode {
	case NORMAL:
		d.view.Mode = INPUT
		d.Drawer = &InputDraw{
			view:   d.view,
			action: NORMAL,
		}
	}
}

func (d *Draw) DoKeyCtrlL() {
	switch d.view.Mode {
	case NORMAL:
		d.view.Mode = LABEL
		d.Drawer = &LabelDraw{
			view:   d.view,
			Labels: d.view.TodoList.GetLabels(),
			Cursor: 0,
		}
	}
}

func (d *Draw) DoKeyCtrlR() {
	switch d.view.Mode {
	case NORMAL:
		d.view.Mode = INPUT
		d.Drawer = &InputDraw{
			view:   d.view,
			action: RENAME,
		}
	}
}

func (d *Draw) DoEnter() {
	switch d.view.Mode {
	case INPUT:
		d.Drawer.DoEnter()
		v, ok := d.Drawer.(*InputDraw)
		if !ok {
			break
		}
		switch v.action {
		case NORMAL:
			d.view.Mode = LABEL
			d.Drawer = &LabelDraw{
				view:   d.view,
				Labels: d.view.TodoList.GetLabels(),
				Cursor: 0,
			}
		case RENAME:
			d.view.Mode = NORMAL
			d.Drawer = &NormalDraw{view: d.view}
		}
		d.SaveTodo()
	case CHANGE, LABEL:
		d.Drawer.DoEnter()
		d.view.Mode = NORMAL
		d.Drawer = &NormalDraw{view: d.view}
		d.SaveTodo()
	}
}

func (d *Draw) Draw() {
	termbox.Clear(colorDef, colorDef)
	termbox.SetCursor(-1, -1)

	d.Drawer.Draw()

	termbox.Flush()
}

func (d *Draw) SaveTodo() {
	d.view.TodoList.Save()
}

func NewDraw(fp *os.File, labels []string) *Draw {
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

	return &Draw{
		Drawer: &NormalDraw{view: view},
		view:   view,
	}
}
