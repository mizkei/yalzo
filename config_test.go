package yalzo

import (
	"bytes"
	"reflect"
	"testing"
)

func TestSaveConf(t *testing.T) {
	data := map[string]Config{
		"{\"labels\":[]}":                  Config{Labels: []string{}},
		"{\"labels\":[\"a\",\"b\",\"c\"]}": Config{Labels: []string{"a", "b", "c"}},
	}

	for k, v := range data {
		var res []byte
		resBuf := bytes.NewBuffer(res)

		if err := SaveConf(resBuf, v); err != nil {
			t.Error(err)
		}
		if !bytes.Equal(resBuf.Bytes(), []byte(k)) {
			t.Errorf("got %v, expected %v", res, v)
		}
	}
}

func TestLoadConf(t *testing.T) {
	data := map[string]Config{
		"{\"labels\":[]}":                  Config{Labels: []string{}},
		"{\"labels\":[\"a\",\"b\",\"c\"]}": Config{Labels: []string{"a", "b", "c"}},
	}

	for k, v := range data {
		conf, err := LoadConf([]byte(k))
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(v, *conf) {
			t.Errorf("got %v, expected %v", *conf, v)
		}
	}
}
