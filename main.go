package main

import (
	"database/sql"
	"net/http"
	"time"
)

func timeToUpdatePic(dur time.Duration) {
	ticker := time.Tick(dur)
	for range ticker {
		tmpPURL, tmpDATE, err := getPictureInfo(-1, 1, "zh-CN")
		if err != nil {
			continue
		}
		tmpPIC := rewriteURL(tmpPURL, tmpDATE)
		db, err := sql.Open("sqlite", "./picture.db")
		if err != nil {
			continue
		}
		defer db.Close()
		err = insertDB(db, tmpPIC[0].DATE, tmpPIC[0].HDURL, tmpPIC[0].UHDURL)
		if err != nil {
			continue
		}
	}
}

func main() {
	go timeToUpdatePic(time.Minute * 30)
	http.HandleFunc("/HDRES/", redirectToHD)
	http.HandleFunc("/UHDRES/", redirectToUHD)
	http.ListenAndServe(":9090", nil)
}
