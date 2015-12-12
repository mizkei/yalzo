package yalzo

import (
	"bytes"
	"testing"
)

func TestReadCSV(t *testing.T) {
	data := bytes.NewBuffer([]byte("l1,t1,true\nl2,t2,false\n"))

	t1, t2, err := ReadCSV(data)
	if err != nil {
		t.Error(err)
	}

	if t1.GetListLength() != 1 || t2.GetListLength() != 1 {
		t.Error("failed read csv")
	}
}

func TestSaveCSV(t *testing.T) {
	expected := "l1,t1,true\n"
	var res []byte
	resBuf := bytes.NewBuffer(res)

	SaveCSV([]Todo{Todo{label: "l1", title: "t1"}}, true, resBuf)
	if resBuf.String() != expected {
		t.Errorf("got %v, expected %v", resBuf.String(), expected)
	}
}
