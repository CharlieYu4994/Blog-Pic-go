package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func redirectToBing(w http.ResponseWriter, r *http.Request) {
	var url string
	var urls picture

	if r.URL.Path != "/bing" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	parm := r.URL.Query()
	res, ok := parm["res"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	datT, ok := parm["dat"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	dat, err := strconv.Atoi(datT[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if dat <= len(picBuffer)-1 && dat >= 0 {
		urls = picBuffer[dat]
	} else if dat == -1 {
		rand.Seed(time.Now().Unix())
		i := rand.Intn(len(picBuffer) - 1)
		urls = picBuffer[i]
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch res[0] {
	case "hdres":
		url = urls.HDURL
	case "uhdres":
		url = urls.UHDURL
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}
