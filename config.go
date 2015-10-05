package yalzo

import (
	"encoding/json"
)

type Config struct {
	Labels []string `json:"labels"`
}

func LoadConf(b []byte) (*Config, error) {
	var conf Config
	if err := json.Unmarshal(b, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
