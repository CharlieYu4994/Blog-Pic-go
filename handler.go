package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const (
	hdSuffix  string = "_1920x1080.jpg"
	uhdSuffix string = "_UHD.jpg"
	domain    string = "https://cn.bing.com"
)

func getDuration(t int) time.Duration {
	now := time.Now()
	tmp := time.Duration(t) * time.Hour
	tstr := now.Format("20060102")
	next, _ := time.ParseInLocation("20060102", tstr, time.Local)
	dur := next.Add(tmp + time.Second*10).Sub(now)
	if now.Hour() >= t {
		return dur + time.Hour*24
	}
	return dur
}

func updatePic(i inserter, v validator) error {
	pics, err := getPicture("zh-CN")
	if err != nil {
		return err
	}

	for j := len(pics) - 1; j >= 0; j-- {
		ok, _ := v(pics[j].Date)
		if ok {
			i(pics[j].Date, pics[j].Burl)
		}
	}
	return nil
}

type handler struct {
	urlbase string
	pic     []picture
}

func (h *handler) redirectToPic(w http.ResponseWriter, r *http.Request) {
	var url string
	var urls picture

	if r.URL.Path != h.urlbase {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	parm := r.URL.Query()
	res, ok := parm["res"]
	if !ok {
		res = append(res, "hdres")
	}
	datT, ok := parm["dat"]
	if !ok {
		datT = append(datT, "0")
	}
	dat, err := strconv.Atoi(datT[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if dat < len(h.pic) && dat >= 0 {
		urls = h.pic[dat]
	} else if dat == -1 {
		rand.Seed(time.Now().Unix())
		i := rand.Intn(len(h.pic) - 1)
		urls = h.pic[i]
	} else {
		urls = h.pic[len(h.pic)-1]
	}

	switch res[0] {
	case "hdres":
		url = domain + urls.Burl + hdSuffix
	case "uhdres":
		url = domain + urls.Burl + uhdSuffix
	default:
		url = domain + urls.Burl + "_" + res[0] + ".jpg"
	}

	http.Redirect(w, r, url, http.StatusFound)
}

func (h *handler) updetePicBuffer(q querier, n int) error {
	tmp, err := q(n)
	if err != nil {
		return err
	}

	for i := 0; i < n; i++ {
		h.pic[i] = tmp[i]
	}
	return nil
}

func (h *handler) timeToUpdatePic() {
	timer := time.NewTimer(0)
	for {
		<-timer.C
		updatePic(dbinserter, dbvalidator)
		h.updetePicBuffer(dbquerier, conf.PicNum)
		timer.Reset(getDuration(conf.UpdateTime))
	}
}
