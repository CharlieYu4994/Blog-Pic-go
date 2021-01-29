package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func redirectToHD(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, picBuffer[0].HDURL, 302)
}

func redirectToUHD(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, picBuffer[0].UHDURL, 302)
}

func redirectToRANDOM(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(picBuffer) - 1)
	http.Redirect(w, r, picBuffer[index].HDURL, 302)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,
		`Here is LassiCat's BingAPI.
1920x1080: https://api.lassi-cat.cn:60443/HDRES
UHD      : https://api.lassi-cat.cn:60443/UHDRES`)
}
