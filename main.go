package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	controller "example.com/rankingSystem/controller/rankList"
	"example.com/rankingSystem/logger"
)

type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	LogLevel string `json:"log_level"`
}

func main() {
	// 初始化 app
	appConf := loadAppConf()
	logger.InitLogger(appConf.LogLevel)
	controller.Register()

	err := http.ListenAndServe(appConf.Host+":"+appConf.Port, nil)
	if err != nil {
		logger.Fatal("ListenAndServe: ", err)
	}
}

func loadAppConf() Config {
	file, err := ioutil.ReadFile("./config/app.json")
	if err != nil {
		panic(err)
	}
	conf := Config{}
	err = json.Unmarshal(file, &conf)
	if err != nil {
		panic(err)
	}
	return conf
}
