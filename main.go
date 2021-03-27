package main

import (
	"database/sql"
	"net/http"
	"time"
)

var dbinserter inserter
var dbquerier querier
var dbvalidator validator

var conf config
var bingHandler handler

func init() {
	err := readConf("./config.json", &conf)
	if err != nil {
		panic("Open Config Error")
	}

	db, err := sql.Open("sqlite3", conf.DataBase)
	if err != nil {
		panic("Open Database Error")
	}
	dbinserter, err = newInserter(db)
	if err != nil {
		panic("Open Database Error")
	}
	dbquerier = newQuerier(db)
	dbvalidator = newValidator(db)

	tmp := make([]picture, conf.PicNum)
	bingHandler.pic = tmp
	bingHandler.urlbase = "/bing"

	go bingHandler.timeToUpdatePic()
}

func main() {
	http.HandleFunc("/bing", bingHandler.redirectToPic)
	time.Sleep(time.Second)
	if conf.EnableTLS {
		http.ListenAndServeTLS("0.0.0.0:"+conf.Port,
			conf.CertPath, conf.KeyPath, nil)
	} else {
		http.ListenAndServe("0.0.0.0:"+conf.Port, nil)
	}
}
