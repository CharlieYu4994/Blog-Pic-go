package main

import (
	"fmt"
	"net/http"
)

func redirectToHD(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, picBuffer[0].HDURL, 302)
}

func redirectToUHD(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, picBuffer[0].UHDURL, 302)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,
		`Here is LassiCat's BingAPI.
1920x1080: https://api.lassi-cat.cn:60443/HDRES
UHD      : https://api.lassi-cat.cn:60443/UHDRES`)
}
