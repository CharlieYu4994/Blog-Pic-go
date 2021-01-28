package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

func getPictureURL(q querier, res int) (string, error) {
	pic, err := q(time.Now().Format("20060102"))
	if err != nil {
		return "", nil
	}
	switch res {
	case 0:
		return pic.HDURL, nil
	case 1:
		return pic.UHDURL, nil
	default:
		return "", errors.New("UnSupported Res")
	}
}

func redirectToHD(w http.ResponseWriter, r *http.Request) {
	URL, err := getPictureURL(dbquerier, 0)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	http.Redirect(w, r, URL, 301)
}

func redirectToUHD(w http.ResponseWriter, r *http.Request) {
	URL, err := getPictureURL(dbquerier, 1)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	http.Redirect(w, r, URL, 301)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,
		`Here is LassiCat's BingAPI.
1920x1080: https://api.lassi-cat.cn:60443/HDRES
UHD      : https://api.lassi-cat.cn:60443/UHDRES`)
}
