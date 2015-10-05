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
	UNIX_HOME        = "HOME"
	WIN_HOME         = "HOMEPATH"
)

func loopDraw(path string, conf yalzo.Config) {
	// open data file
	fp, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	dr := yalzo.NewDraw(fp, conf.Labels)
	defer dr.SaveTodo()

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
	var envHOME string
	switch runtime.GOOS {
	case "linux", "darwin":
		envHOME = UNIX_HOME
	case "windows":
		envHOME = WIN_HOME
	}

	home := os.Getenv(envHOME)
	if home == "" {
		panic("ENV " + envHOME + " does not exist")
	}

	yalzoPath := path.Join(home, YALZO_PATH)
	os.Mkdir(yalzoPath, 0666)

	// read config
	cf, err := ioutil.ReadFile(path.Join(yalzoPath, CONFIG_FILE_NAME))
	if err != nil {
		panic(err)
	}
	conf := yalzo.LoadConf(cf)

	loopDraw(path.Join(yalzoPath, DATA_FILE_NAME), conf)
}
