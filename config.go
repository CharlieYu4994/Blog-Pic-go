package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type config struct {
	EnableTLS  bool   `json:"enabletls"`
	CertPath   string `json:"certpath"`
	KeyPath    string `json:"keypath"`
	Port       string `json:"port"`
	UpdateTime int    `json:"updatetime"`
}

func readConf(path string, conf *config) error {
	_, err := os.Stat(path)
	if err != nil || os.IsExist(err) {
		fmt.Println("Error No Config File")
		return errors.New("Error No Config File")
	}
	tmp, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, conf)
}
