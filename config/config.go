package config

import (
	"golang_toDoList/utils"
	"log"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
	Static    string
}

var Config ConfigList

// メイン関数より前に実行
func init() {
	LoadConfig()
	utils.LoggingSetting(Config.LogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")

	if err != nil {
		log.Fatal(err)
	}

	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
		Static:    cfg.Section("web").Key("static").String(),
	}
}
