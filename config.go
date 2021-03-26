package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type config struct {
	EnableTLS  bool   `json:"enabletls"`
	UpdateTime int    `json:"updatetime"`
	PicNum     int    `json:"picnum"`
	CertPath   string `json:"certpath"`
	KeyPath    string `json:"keypath"`
	Port       string `json:"port"`
	DataBase   string `json:"database"`
}

func readConf(path string, conf *config) error {
	_, err := os.Stat(path)
	if err != nil || os.IsExist(err) {
		return errors.New("ErrorNoConfigFile")
	}
	tmp, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, conf)
}
