package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

var dbinserter inserter
var dbquerier querier
var dbvalidator validator
var conf config

func updatePic(i inserter, v validator) (err error) {
	tmpPURL, tmpDATE, err := getPictureInfo(-1, 1, "zh-CN")
	if err != nil {
		return
	}
	tmpPICs := rewriteURL(tmpPURL, tmpDATE)
	for _, tmpPIC := range tmpPICs {
		status, _ := v(tmpPIC.DATE)
		if status {
			i(tmpPIC.DATE, tmpPIC.HDURL, tmpPIC.UHDURL)
		}
	}
	return
}

func getDuration(hour int) time.Duration {
	now := time.Now()
	tmp := time.Duration(hour) * time.Hour
	tstr := now.Format("20060102")
	next, _ := time.ParseInLocation("20060102", tstr, time.Local)
	dur := next.Add(tmp + time.Second*10).Sub(now)
	if now.Hour() >= hour {
		return dur + time.Hour*24
	}
	return dur
}

func timeToUpdatePic() {
	timer := time.NewTimer(time.Second * 10)
	for {
		<-timer.C
		updatePic(dbinserter, dbvalidator)
		timer.Reset(getDuration(conf.UpdateTime))
	}
}

func init() {
	db, err := sql.Open("sqlite", "./picture.db")
	if err != nil {
		panic("Open Database Error")
	}
	dbinserter, err = newInserter(db)
	if err != nil {
		panic("Open Database Error")
	}
	dbquerier = newQuerier(db)
	dbvalidator = newValidator(db)
	readConf("./config.json", &conf)
	fmt.Println(conf)
	go timeToUpdatePic()
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/HDRES/", redirectToHD)
	http.HandleFunc("/UHDRES/", redirectToUHD)
	if conf.EnableTLS {
		http.ListenAndServeTLS("0.0.0.0:"+conf.Port,
			conf.CertPath, conf.KeyPath, nil)
	} else {
		http.ListenAndServe("0.0.0.0:"+conf.Port, nil)
	}
}
