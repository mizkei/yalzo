package main

import (
	"io/ioutil"
	"os"
	"path"
	"runtime"

	"github.com/mizkei/yalzo"
	"github.com/nsf/termbox-go"
)

const (
	YALZO_PATH       = ".yalzo"
	CONFIG_FILE_NAME = "config.json"
	DATA_FILE_NAME   = "todo.csv"
)

func loopDraw(fp *os.File, conf yalzo.Config) {
	view := yalzo.NewView(fp, conf.Labels)

	for {
		view.Draw()
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyCtrlQ:
				return
			case termbox.KeyEsc:
				view.DoKeyEsc()
			case termbox.KeyArrowLeft:
				view.DoKeyArrowLeft()
			case termbox.KeyCtrlB:
				view.DoKeyCtrlB()
			case termbox.KeyArrowRight:
				view.DoKeyArrowRight()
			case termbox.KeyCtrlF:
				view.DoKeyCtrlF()
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				view.DoKeyBackspace()
			case termbox.KeyDelete:
				view.DoKeyDelete()
			case termbox.KeyTab:
				view.DoKeyTab()
			case termbox.KeyCtrlX:
				view.DoKeyCtrlX()
			case termbox.KeyCtrlW:
				view.DoKeyCtrlW()
			case termbox.KeyCtrlL:
				view.DoKeyCtrlL()
			case termbox.KeyCtrlD:
				view.DoKeyCtrlD()
			case termbox.KeyCtrlA:
				view.DoKeyCtrlA()
			case termbox.KeyCtrlR:
				view.DoKeyCtrlR()
			case termbox.KeySpace:
				view.DoKeySpace()
			case termbox.KeyEnter:
				view.DoEnter()
			default:
				view.DoChar(ev.Ch)
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

func main() {
	// init termbox
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	// set yalzo path
	var yalzoPath string
	switch runtime.GOOS {
	case "linux":
		yalzoPath = path.Join(os.Getenv("HOME"), YALZO_PATH)
	case "dawwin":
		yalzoPath = path.Join(os.Getenv("HOME"), YALZO_PATH)
	case "windows":
		yalzoPath = path.Join(os.Getenv("HOMEPATH"), YALZO_PATH)
	}

	// read config
	cf, err := ioutil.ReadFile(path.Join(yalzoPath, CONFIG_FILE_NAME))
	if err != nil {
		panic(err)
	}
	conf := yalzo.LoadConf(cf)

	// open data file
	df, err := os.Open(path.Join(yalzoPath, DATA_FILE_NAME))
	if err != nil {
		panic(err)
	}

	loopDraw(df, conf)
}
