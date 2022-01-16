package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	controller "example.com/rankingSystem/controller/rankList"
)

type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func main() {
	controller.Register()
	appConf := loadAppConf()
	err := http.ListenAndServe(appConf.Host+":"+appConf.Port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
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
