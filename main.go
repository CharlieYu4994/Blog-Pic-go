package main

import (
	"database/sql"
	"net/http"
	"time"
)

var dbinserter inserter
var dbquerier querier
var dbvalidator validator

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
		timer.Reset(getDuration(16))
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
	go timeToUpdatePic()
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/HDRES/", redirectToHD)
	http.HandleFunc("/UHDRES/", redirectToUHD)
	http.ListenAndServe("0.0.0.0:9090", nil)
}
