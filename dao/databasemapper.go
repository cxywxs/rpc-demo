package dao

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Databases struct {
	Server   string `toml:"server"`
	Port     string `toml:"port"`
	Root     string `toml:"root"`
	Password string `toml:"password"`
	Database string `toml:"database"`
}

type Register struct {
	Server string `toml:"server"`
	Port   string `toml:"port"`
}

type ServerInfo struct {
	Server string `toml:"server"`
	Port   string `toml:"port"`
}

type Application struct {
	Databases  *Databases
	Register   *Register
	ServerInfo *ServerInfo
}

var application = new(Application)

func AnalysisApplication() *Application {
	_, err := toml.DecodeFile("application.toml", application)
	if err != nil {
		log.Println(err)
		log.Fatal(err)
	}
	return application
}
