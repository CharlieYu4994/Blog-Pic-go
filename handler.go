package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"
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
	tmpPURL, tmpDATE, err := getPictureInfo(-1, 9, "zh-CN")
	if err != nil {
		return err
	}

	tmpPICs := rewriteURL(tmpPURL, tmpDATE)
	for index := len(tmpPICs) - 1; index >= 0; index-- {
		tmpPIC := tmpPICs[index]
		status, _ := v(tmpPIC.DATE)
		if status {
			i(tmpPIC.DATE, tmpPIC.HDURL, tmpPIC.UHDURL)
		}
	}
	return nil
}

type handler struct {
	pic []*picture
}

func (h *handler) redirectToPic(w http.ResponseWriter, r *http.Request) {
	var url string
	var urls *picture

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

	if dat <= len(h.pic)-1 && dat >= 0 {
		urls = h.pic[dat]
	} else if dat == -1 {
		rand.Seed(time.Now().Unix())
		i := rand.Intn(len(h.pic) - 1)
		urls = h.pic[i]
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

func (h *handler) updetePicBuffer(q querier, num int) error {
	tmp, err := q(num)
	if err != nil {
		return err
	}
	for index := 0; index < num; index++ {
		h.pic[index] = tmp[index]
	}
	return nil
}

func (h *handler) timeToUpdatePic() {
	timer := time.NewTimer(0)
	for {
		<-timer.C
		updatePic(dbinserter, dbvalidator)
		h.updetePicBuffer(dbquerier, 7)
		timer.Reset(getDuration(conf.UpdateTime))
	}
}
