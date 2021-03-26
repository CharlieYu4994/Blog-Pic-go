package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	hdSuffix  string = "_1920x1080.jpg"
	uhdSuffix string = "_UHD.jpg"
	domain    string = "https://cn.bing.com"
)

type picture struct {
	HDURL, UHDURL, DATE string
}

func getPictureInfo(idx int, num int, mkt string) (pURL []string, date []string, err error) {
	gURL := fmt.Sprintf(
		"%s/HPImageArchive.aspx?format=js&idx=%d&n=%d&mkt=%s", domain, idx, num, mkt)

	response, err := http.Get(gURL)
	if err != nil {
		return nil, nil, err
	}
	msg, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}

	var tmp *struct {
		Images []struct {
			Urlbase string `json:"urlbase"`
			Enddate string `json:"enddate"`
		}
	}
	err = json.Unmarshal(msg, &tmp)
	if err != nil {
		return nil, nil, err
	}

	for _, data := range tmp.Images {
		pURL = append(pURL, data.Urlbase)
		date = append(date, data.Enddate)
	}
	return pURL, date, nil
}

func rewriteURL(pURL []string, date []string) (pic []*picture) {
	tmp := &picture{}
	for i := 0; i < len(pURL); i++ {
		st := domain + pURL[i]
		tmp.HDURL = st + hdSuffix
		tmp.UHDURL = st + uhdSuffix
		tmp.DATE = date[i]
		pic = append(pic, tmp)
	}
	return
}
