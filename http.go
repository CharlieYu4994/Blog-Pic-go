package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

func getPictureURL(q queryer, res int) (string, error) {
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
	URL, err := getPictureURL(dbqueryer, 0)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	http.Redirect(w, r, URL, 301)
}

func redirectToUHD(w http.ResponseWriter, r *http.Request) {
	URL, err := getPictureURL(dbqueryer, 1)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	http.Redirect(w, r, URL, 301)
}
