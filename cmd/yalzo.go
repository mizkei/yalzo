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
	dr := yalzo.NewDraw(fp, conf.Labels)

	for {
		dr.Drawer.Draw()
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyCtrlQ:
				return
			case termbox.KeyEsc:
				dr.DoKeyEsc()
			case termbox.KeyArrowLeft:
				dr.DoKeyArrowLeft()
			case termbox.KeyCtrlB:
				dr.DoKeyCtrlB()
			case termbox.KeyArrowRight:
				dr.DoKeyArrowRight()
			case termbox.KeyCtrlF:
				dr.DoKeyCtrlF()
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				dr.DoKeyBackspace()
			case termbox.KeyDelete:
				dr.DoKeyDelete()
			case termbox.KeyTab:
				dr.DoKeyTab()
			case termbox.KeyCtrlX:
				dr.DoKeyCtrlX()
			case termbox.KeyCtrlW:
				dr.DoKeyCtrlW()
			case termbox.KeyCtrlL:
				dr.DoKeyCtrlL()
			case termbox.KeyCtrlD:
				dr.DoKeyCtrlD()
			case termbox.KeyCtrlA:
				dr.DoKeyCtrlA()
			case termbox.KeyCtrlR:
				dr.DoKeyCtrlR()
			case termbox.KeySpace:
				dr.DoKeySpace()
			case termbox.KeyEnter:
				dr.DoEnter()
			default:
				dr.DoChar(ev.Ch)
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
	df, err := os.OpenFile(path.Join(yalzoPath, DATA_FILE_NAME), os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	loopDraw(df, conf)
}
