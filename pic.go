package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type picture struct {
	Date string `json:"enddate"`
	Burl string `json:"urlbase"`
}

func getPicture(mkt string) (pics []picture, err error) {
	gURL := fmt.Sprintf("%s/HPImageArchive.aspx?format=js&idx=-1&n=9&mkt=%s", domain, mkt)

	response, err := http.Get(gURL)
	if err != nil {
		return nil, err
	}

	msg, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}

	var tmp struct {
		Images []picture
	}
	err = json.Unmarshal(msg, &tmp)
	if err != nil {
		return nil, err
	}
	return tmp.Images, nil
}
