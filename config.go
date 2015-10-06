package yalzo

import (
	"encoding/json"
	"io"
)

type Config struct {
	Labels []string `json:"labels"`
}

func SaveConf(w io.Writer, conf Config) error {
	b, err := json.Marshal(conf)
	if err != nil {
		return err
	}

	if _, err := w.Write(b); err != nil {
		return err
	}

	return nil
}

func LoadConf(b []byte) (*Config, error) {
	var conf Config
	if err := json.Unmarshal(b, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
