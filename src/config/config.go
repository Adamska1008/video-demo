package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Net   Net   `json:"net"`
	Mysql Mysql `json:"mysql"`
}

type Net struct {
	Addr string
	Port int
}

type Mysql struct {
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
}

func Load(path string) (*Config, error) {
	dataBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	config := Config{}
	err = yaml.Unmarshal(dataBytes, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
