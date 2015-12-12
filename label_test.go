package yalzo

import (
	"testing"
)

func initLabelList() LabelList {
	return []string{"l1", "l2", "l3", "l4"}
}

func TestGetListLength(t *testing.T) {
	ls := initLabelList()

	if ls.GetListLength() != 4 {
		t.Error("label: get list length")
	}
}

func TestGetList(t *testing.T) {
	initList := initLabelList()
	ls := initList.GetList(10)
	expected := []string{"l1", "l2", "l3", "l4"}

	for i := range ls {
		if ls[i] != expected[i] {
			t.Errorf("got %v, expected %v", ls[i], expected[i])
		}
	}
}

func TestIsInAryRange(t *testing.T) {
	ls := initLabelList()

	if !ls.IsInAryRange(3) {
		t.Error("in range error")
	}

	if ls.IsInAryRange(4) {
		t.Error("in range error")
	}
}

func TestContains(t *testing.T) {
	ls := initLabelList()

	if i, ok := ls.Contains("l2"); i != 1 || !ok {
		t.Error("error contains")
	}

	if i, ok := ls.Contains("l100"); i != 0 || ok {
		t.Error("error contains")
	}
}

func TestAdd(t *testing.T) {
	ls := initLabelList()

	i := ls.Add("l5")
	if i != 5 {
		t.Error("add error")
	}

	if ls[i-1] != "l5" {
		t.Error("add error")
	}
}

func TestRemove(t *testing.T) {
	ls := initLabelList()

	label := ls.Remove(0).(string)
	if label != "l1" {
		t.Error("error remove")
	}

	if len(ls) != 3 {
		t.Error("error remove")
	}
}

func TestGetPresentName(t *testing.T) {
	ls := initLabelList()

	name := ls.GetPresentName(0)
	if name != "l1" {
		t.Error("get name")
	}
}

func TestRename(t *testing.T) {
	ls := initLabelList()

	ls.Rename(0, "l10")
	if ls[0] != "l10" {
		t.Error("rename")
	}
}

func TestExchange(t *testing.T) {
	ls := initLabelList()

	ls.Exchange(0, 3)
	if ls[0] != "l4" || ls[3] != "l1" {
		t.Error("exchange")
	}
}

func TestChangeLabel(t *testing.T) {
	ls := initLabelList()

	ls.ChangeLabel(0, "l10")
	if ls[0] != "l10" {
		t.Error("rename")
	}
}
