package yalzo

import (
	"encoding/json"
)

type Config struct {
	Labels []string `json:"labels"`
}

func LoadConf(b []byte) Config {
	var conf Config
	if err := json.Unmarshal(b, &conf); err != nil {
		panic(err)
	}

	return conf
}
