package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"text/template"
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
		pageNotFound(w)
		return
	}
	http.Redirect(w, r, picBuffer[0].UHDURL, 302)
}

func redirectToRANDOM(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/RANDOM/" {
		w.WriteHeader(http.StatusNotFound)
		pageNotFound(w)
		return
	}
	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(picBuffer) - 1)
	http.Redirect(w, r, picBuffer[index].HDURL, 302)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		pageNotFound(w)
		return
	}
	tmpl, err := template.ParseFiles("./Web/template.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
	}
	tmpl.Execute(w, conf.HTML)
}

func pageNotFound(w http.ResponseWriter) {
	content, err := ioutil.ReadFile("./Web/404.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprint(w, string(content))
}
