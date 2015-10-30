package main

import (
	"github.com/BurntSushi/toml"
	"github.com/operando/golack"
)

const (
	KEY_UPDATE_POST = "update_post"
	KEY_ERROR_POST  = "error_post"
)

type Config struct {
	Slacks    map[string]golack.Slack `toml:"slacks"`
	Webhook   golack.Webhook          `toml:"webhook"`
	Log       string                  `toml:"log"`
	SleepTime int                     `toml:"sleeptime"`
	Android   Android                 `toml:"android"`
	Ios       Ios                     `toml:"ios"`
	ErrorPost bool                    `toml:"error_post"`
}

type Android struct {
	Package string `toml:"package"`
}

type Ios struct {
	Country string `toml:"country"`
	AppId   string `toml:"app_id"`
}

func LoadConfig(configPath string, config *Config) (*Config, error) {
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		return config, err
	}
	return config, nil
}
