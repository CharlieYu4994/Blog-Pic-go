package main

import (
	"database/sql"
	"flag"
	"net/http"
	"time"
)

var dbinserter inserter
var dbquerier querier
var dbvalidator validator

var conf config
var confpath string
var bingHandler handler

func init() {
	flag.StringVar(&confpath, "c", "./config.json", "Set the config path")

	flag.Parse()

	err := readConf(confpath, &conf)
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
	bingHandler.urlbase = "/" + conf.UrlBase

	go bingHandler.timeToUpdatePic(conf.UpdateTime, conf.PicNum)
}

func main() {
	http.HandleFunc("/"+conf.UrlBase, bingHandler.redirectToPic)
	time.Sleep(time.Second)
	if conf.EnableTLS {
		http.ListenAndServeTLS("0.0.0.0:"+conf.Port,
			conf.CertPath, conf.KeyPath, nil)
	} else {
		http.ListenAndServe("0.0.0.0:"+conf.Port, nil)
	}
}
