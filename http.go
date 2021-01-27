package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func getPictureURL(res int) (pURL string, err error) {
	db, err := sql.Open("sqlite", "./picture.db")
	if err != nil {
		return
	}
	defer db.Close()
	date := time.Now().Format("20060102")
	pic, err := queryDB(db, date)
	if err != nil {
		return
	}
	switch res {
	case 0:
		pURL = pic.HDURL
	case 1:
		pURL = pic.UHDURL
	default:
		err = errors.New("UnSupported Res")
	}
	return
}

func redirectToHD(w http.ResponseWriter, r *http.Request) {
	URL, err := getPictureURL(0)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	http.Redirect(w, r, URL, 301)
}

func redirectToUHD(w http.ResponseWriter, r *http.Request) {
	URL, err := getPictureURL(1)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	http.Redirect(w, r, URL, 301)
}
