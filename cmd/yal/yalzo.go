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

func loopDraw(todopath, confpath string) {
	// open data file
	fp, err := os.OpenFile(todopath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	// read config
	cf, err := ioutil.ReadFile(confpath)
	if err != nil {
		cf = []byte("{\"labels\": [\"完了\", \"未完了\"]}")
	}
	conf, err := yalzo.LoadConf(cf)
	if err != nil {
		panic(err)
	}

	dr := yalzo.NewDraw(fp, conf.Labels)
	defer dr.SaveTodo()
	defer func() {
		fp, err := os.OpenFile(confpath, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			return // TODO: 設定が保存できない場合の処理を追加
		}
		defer fp.Close()

		fp.Truncate(0)
		err = yalzo.SaveConf(fp, yalzo.Config{
			Labels: dr.GetLabels(),
		})
		if err != nil {
			return // TODO: 設定が保存できない場合の処理を追加
		}
	}()

	for {
		dr.Draw()
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
			case termbox.KeyCtrlV:
				dr.DoKeyCtrlV()
			case termbox.KeyCtrlC:
				dr.DoKeyCtrlC()
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
	os.Mkdir(yalzoPath, 0755)

	loopDraw(path.Join(yalzoPath, DATA_FILE_NAME), path.Join(yalzoPath, CONFIG_FILE_NAME))
}
