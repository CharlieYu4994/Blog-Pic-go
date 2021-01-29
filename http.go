package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"text/template"
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
	tmpl, err := template.ParseFiles("./Web/template.html")
	if err != nil {
		fmt.Fprintln(w, err)
	}
	tmpl.Execute(w, conf.HTML)
}
