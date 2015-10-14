package yalzo

import (
	"unicode/utf8"

	"github.com/mattn/go-runewidth"
)

func getVoffset(text []byte, boffset int) int {
	var voffset int

	text = text[:boffset]
	for len(text) > 0 {
		r, size := utf8.DecodeRune(text)
		text = text[size:]
		voffset += runewidth.RuneWidth(r)
	}

	return voffset
}

func makeCap(s []byte, desiredCap int) []byte {
	if cap(s) < desiredCap {
		ns := make([]byte, len(s), desiredCap)
		copy(ns, s)
		return ns
	}
	return s
}

func removeByte(text []byte, from, to int) []byte {
	size := to - from
	copy(text[from:], text[to:])
	text = text[:len(text)-size]
	return text
}

func insertByte(text []byte, offset int, what []byte) []byte {
	n := len(text) + len(what)
	text = makeCap(text, n)
	text = text[:n]
	copy(text[offset+len(what):], text[offset:])
	copy(text[offset:], what)
	return text
}

type InputBox struct {
	input         []byte
	prefix        string
	prefixWidth   int
	cursorBOffset int
	cursorVOffset int
}

func (ib *InputBox) GetInputString() string {
	return ib.prefix + string(ib.input)
}

func (ib *InputBox) MoveCursorTo(bOffset int) {
	ib.cursorBOffset = bOffset
	ib.cursorVOffset = getVoffset(ib.input, bOffset)
}

func (ib *InputBox) RuneUnderCursor() (rune, int) {
	return utf8.DecodeRune(ib.input[ib.cursorBOffset:])
}

func (ib *InputBox) RuneBeforeCursor() (rune, int) {
	return utf8.DecodeLastRune(ib.input[:ib.cursorBOffset])
}

func (ib *InputBox) MoveCursorOneRuneBackward() {
	if ib.cursorBOffset == 0 {
		return
	}
	_, size := ib.RuneBeforeCursor()
	ib.MoveCursorTo(ib.cursorBOffset - size)
}

func (ib *InputBox) MoveCursorOneRuneForward() {
	if ib.cursorBOffset == len(ib.input) {
		return
	}
	_, size := ib.RuneUnderCursor()
	ib.MoveCursorTo(ib.cursorBOffset + size)
}

func (ib *InputBox) DeleteRuneBackward() {
	if ib.cursorBOffset == 0 {
		return
	}

	ib.MoveCursorOneRuneBackward()
	_, size := ib.RuneUnderCursor()
	ib.input = removeByte(ib.input, ib.cursorBOffset, ib.cursorBOffset+size)
}

func (ib *InputBox) DeleteRuneForward() {
	if ib.cursorBOffset == len(ib.input) {
		return
	}
	_, size := ib.RuneUnderCursor()
	ib.input = removeByte(ib.input, ib.cursorBOffset, ib.cursorBOffset+size)
}

func (ib *InputBox) DeleteAll() {
	ib.MoveCursorTo(0)
	ib.input = []byte{}
}

func (ib *InputBox) InsertRune(r rune) {
	var buf [utf8.UTFMax]byte
	n := utf8.EncodeRune(buf[:], r)
	ib.input = insertByte(ib.input, ib.cursorBOffset, buf[:n])
	ib.MoveCursorOneRuneForward()
}

func (ib *InputBox) InsertStr(s string) {
	for _, r := range s {
		ib.InsertRune(r)
	}
}

func (ib *InputBox) GetCursorPos() int {
	return ib.cursorVOffset + ib.prefixWidth
}
