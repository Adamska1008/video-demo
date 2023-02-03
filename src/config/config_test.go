package config

import (
	"fmt"
	"testing"
)

func Test_Load(t *testing.T) {
	expect := Config{
		Net:   Net{Addr: "127.0.0.1", Port: 8080},
		Mysql: Mysql{Host: "127.0.0.1", Port: 3306, Username: "adam", Password: "1234", DbName: "video_demo"},
	}
	config, err := Load("..\\..\\config.yml")
	if err != nil {
		t.Error(err)
	}
	if expect != *config {
		t.Errorf("expect: %v, but get: %v", expect, config)
	}
	fmt.Println(config)
}
