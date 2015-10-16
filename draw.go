package yalzo

import (
	"os"

	"github.com/mizkei/yalzo/mode"
	"github.com/nsf/termbox-go"
)

const (
	colorDef = termbox.ColorDefault
)

type Drawer interface {
	mode.Operator
	GetCursorIndex() int
	GetSelectedIndex() []int
	SetLister(mode.Lister)
	GetListLength() int
	Mode() mode.Mode
	Reset()
	Draw()
}

type Target struct {
	Index int
	Tab   Tab
}

type Draw struct {
	Drawer
	data   *Data
	tab    Tab
	target *Target
	file   *os.File
}

func (d *Draw) DoKeyEsc() {
	w, h := termbox.Size()

	switch d.Drawer.Mode() {
	case mode.INPUT, mode.EXCHANGE:
		view := mode.NewView(w, h)
		d.Drawer = &mode.NormalDraw{
			View: *view,
			Tab:  d.tab,
		}
		d.SetLister(d.tab)
	case mode.NORMAL:
		d.Reset()
	}
}

func (d *Draw) DoKeyTab() {
	if d.tab == LABEL {
		return
	}

	w, h := termbox.Size()

	switch d.Drawer.Mode() {
	case mode.NORMAL:
		if d.tab == INPROCESS {
			d.tab = ARCHIVE
		} else if d.tab == ARCHIVE {
			d.tab = INPROCESS
		}

		view := mode.NewView(w, h)
		d.Drawer = &mode.NormalDraw{
			View: *view,
			Tab:  d.tab,
		}
		d.SetLister(d.tab)
	}
}

func (d *Draw) DoKeyCtrlA() {
	if d.tab == LABEL {
		return
	}

	w, h := termbox.Size()

	switch d.Drawer.Mode() {
	case mode.NORMAL:
		selected := d.GetSelectedIndex()
		if len(selected) == 0 {
			d.data.MoveTodo(d.GetCursorIndex(), d.tab)
		} else {
			for _, i := range selected {
				d.data.MoveTodo(i, d.tab)
				mode.ShiftIndex(&selected, i)
			}
		}

		view := mode.NewView(w, h)
		d.Drawer = &mode.NormalDraw{
			View: *view,
			Tab:  d.tab,
		}
		d.SetLister(d.tab)
	}
}

func (d *Draw) DoKeyCtrlW() {
	w, h := termbox.Size()

	switch d.Drawer.Mode() {
	case mode.NORMAL:
		view := mode.NewView(w, h)
		d.Drawer = &mode.InputDraw{
			View: *view,
			Act:  mode.INSERT,
		}
		d.SetLister(d.tab)
	default:
		d.Drawer.DoKeyCtrlW()
	}
}

func (d *Draw) DoKeyCtrlL() {
	w, h := termbox.Size()

	switch d.Drawer.Mode() {
	case mode.NORMAL:
		if d.tab == LABEL || d.GetListLength() == 0 {
			return
		}

		view := mode.NewView(w, h)
		d.target = &Target{
			Index: d.GetCursorIndex(),
			Tab:   d.tab,
		}
		d.tab = LABEL
		d.Drawer = &mode.LabelSetDraw{
			View: *view,
			Tab:  d.tab,
		}
		d.SetLister(d.tab)
	default:
		d.Drawer.DoKeyCtrlL()
	}
}

func (d *Draw) DoKeyCtrlV() {
	w, h := termbox.Size()

	switch d.Drawer.Mode() {
	case mode.NORMAL:
		if d.tab == LABEL {
			d.tab = INPROCESS
		} else {
			d.tab = LABEL
		}

		view := mode.NewView(w, h)
		d.Drawer = &mode.NormalDraw{
			View: *view,
			Tab:  d.tab,
		}
		d.SetLister(d.tab)
	default:
		d.Drawer.DoKeyCtrlV()
	}
}

func (d *Draw) DoKeyCtrlX() {
	w, h := termbox.Size()

	switch d.Drawer.Mode() {
	case mode.NORMAL:
		view := mode.NewViewWithCheck(w, h, d.GetCursorIndex())
		d.Drawer = &mode.ExchangeDraw{
			View: *view,
			Tab:  d.tab,
		}
		d.SetLister(d.tab)
	default:
		d.Drawer.DoKeyCtrlV()
	}
}

func (d *Draw) DoKeyCtrlR() {
	w, h := termbox.Size()

	switch d.Drawer.Mode() {
	case mode.NORMAL:
		view := mode.NewView(w, h)
		view.Cursor = d.GetCursorIndex()
		d.Drawer = &mode.InputDraw{
			View: *view,
			Act:  mode.RENAME,
		}
		d.SetLister(d.tab)
	default:
		d.Drawer.DoKeyCtrlR()
	}
}

func (d *Draw) DoEnter() {
	w, h := termbox.Size()
	d.Drawer.DoEnter()

	switch d.Drawer.Mode() {
	case mode.INPUT:
		iDraw := d.Drawer.(*mode.InputDraw)
		view := mode.NewView(w, h)
		if iDraw.Act == mode.RENAME || d.tab == LABEL {
			d.Drawer = &mode.NormalDraw{
				View: *view,
				Tab:  d.tab,
			}
		} else {
			d.target = &Target{
				Index: d.GetListLength() - 1,
				Tab:   d.tab,
			}
			d.tab = LABEL
			d.Drawer = &mode.LabelSetDraw{
				View: *view,
				Tab:  d.tab,
			}
		}
		d.SetLister(d.tab)
	case mode.EXCHANGE:
		view := mode.NewView(w, h)
		d.Drawer = &mode.NormalDraw{
			View: *view,
			Tab:  d.tab,
		}
		d.SetLister(d.tab)
	case mode.LABELSET:
		view := mode.NewView(w, h)
		d.tab = d.target.Tab
		nd := &mode.NormalDraw{
			View: *view,
			Tab:  d.tab,
		}
		i := d.data.Labels.GetPresentName(d.GetCursorIndex())
		d.Drawer = nd
		d.SetLister(d.tab)
		nd.ChangeLabel(d.target.Index, i)
	}
}

func (d *Draw) Draw() {
	termbox.Clear(colorDef, colorDef)
	termbox.SetCursor(-1, -1)

	d.Drawer.Draw()

	termbox.Flush()
}

func (d *Draw) SetLister(tab Tab) {
	switch tab {
	case INPROCESS:
		d.Drawer.SetLister(d.data.InProcess)
	case ARCHIVE:
		d.Drawer.SetLister(d.data.Archive)
	case LABEL:
		d.Drawer.SetLister(d.data.Labels)
	}
}

func (d *Draw) SaveData() {
	d.file.Seek(0, 0)
	d.file.Truncate(0)
	SaveCSV(*d.data.InProcess, false, d.file)
	SaveCSV(*d.data.Archive, true, d.file)
}

func (d *Draw) GetLabels() []string {
	return *d.data.Labels
}

func NewDraw(fp *os.File, labels []string) *Draw {
	w, h := termbox.Size()
	tab := INPROCESS

	inproc, arch, err := ReadCSV(fp)
	if err != nil {
		panic(err)
	}
	fp.Seek(0, 0)

	data := &Data{Labels: &LabelList{}}
	data.InProcess = inproc
	data.Archive = arch
	*data.Labels = labels

	view := mode.NewView(w, h)
	view.Reset()
	drawer := &mode.NormalDraw{
		View: *view,
		Tab:  tab,
	}
	drawer.SetLister(data.InProcess)

	return &Draw{
		Drawer: drawer,
		data:   data,
		tab:    tab,
		file:   fp,
	}
}
