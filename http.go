package main

import (
	"math/rand"
	"net/http"
	"time"
)

func redirectToHD(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/HDRES/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	http.Redirect(w, r, picBuffer[0].HDURL, 302)
}

func redirectToUHD(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/UHDRES/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	http.Redirect(w, r, picBuffer[0].UHDURL, 302)
}

func redirectToRANDOM(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/RANDOM/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(picBuffer) - 1)
	http.Redirect(w, r, picBuffer[index].HDURL, 302)
}
