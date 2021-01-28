package main

import (
	"database/sql"
	"net/http"
	"time"
)

var dbinserter inserter
var dbqueryer queryer

func updatePic(i inserter) (err error) {
	tmpPURL, tmpDATE, err := getPictureInfo(-1, 1, "zh-CN")
	if err != nil {
		return
	}
	tmpPIC := rewriteURL(tmpPURL, tmpDATE)
	err = i(tmpPIC[0].DATE, tmpPIC[0].HDURL, tmpPIC[0].UHDURL)
	return
}

func timeToUpdatePic(dur time.Duration) {
	ticker := time.Tick(dur)
	for range ticker {
		updatePic(dbinserter)
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
	dbqueryer = newQueryer(db)
	updatePic(dbinserter)
}

func main() {
	go timeToUpdatePic(time.Minute * 10)
	http.HandleFunc("/", homePage)
	http.HandleFunc("/HDRES/", redirectToHD)
	http.HandleFunc("/UHDRES/", redirectToUHD)
	http.ListenAndServe("0.0.0.0:9090", nil)
}
