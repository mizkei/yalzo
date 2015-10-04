package yalzo

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

func PrintText(x, y int, fg, bg termbox.Attribute, text string) int {
	for _, c := range text {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}

	return x
}

func PrintLine(y int, str string) int {
	colorDef := termbox.ColorDefault
	PrintText(0, y, colorDef, colorDef, str)

	return y + 1
}
