package main

import (
	"database/sql"
	"net/http"
	"time"
)

func updatePic() (err error) {
	tmpPURL, tmpDATE, err := getPictureInfo(-1, 1, "zh-CN")
	if err != nil {
		return
	}
	tmpPIC := rewriteURL(tmpPURL, tmpDATE)
	db, err := sql.Open("sqlite", "./picture.db")
	if err != nil {
		return
	}
	defer db.Close()
	err = insertDB(db, tmpPIC[0].DATE, tmpPIC[0].HDURL, tmpPIC[0].UHDURL)
	if err != nil {
		return
	}
	return
}

func timeToUpdatePic(dur time.Duration) {
	ticker := time.Tick(dur)
	for range ticker {
		updatePic()
	}
}

func init() {
	updatePic()
}

func main() {
	go timeToUpdatePic(time.Minute * 30)
	http.HandleFunc("/HDRES/", redirectToHD)
	http.HandleFunc("/UHDRES/", redirectToUHD)
	http.ListenAndServe(":9090", nil)
}
