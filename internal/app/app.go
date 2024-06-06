package app

import "time"

type App struct {
	Cfg     *Config
	Started time.Time
}

var State *App

func Init(configFile string) error {
	cfg := &Config{}
	err := LoadConfig(configFile, cfg)
	if err != nil {
		return err
	}

	State = &App{
		Cfg:     cfg,
		Started: time.Now(),
	}

	return nil
}
