/**
 * Copyright (C) 2021 CharlieYu4994
 *
 * This file is part of Blog-Pic-go.
 *
 * Blog-Pic-go is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Blog-Pic-go is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with Blog-Pic-go.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type picture struct {
	Date    string `json:"enddate"`
	BaseURL string `json:"urlbase"`
}

func httpGet(url string) ([]byte, bool) {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return nil, false
	}
	defer resp.Body.Close()

	msg, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, false
	}
	return msg, true
}

func getBing(num int) ([]picture, bool) {
	url := bing + "/HPImageArchive.aspx?format=js&n=2&mkt=zh-CN"

	msg, ok := httpGet(url)
	if !ok {
		return nil, false
	}

	var tmp struct {
		Images []picture
	}
	err := json.Unmarshal(msg, &tmp)
	if err != nil {
		return nil, false
	}
	return tmp.Images, true
}

func getAPOD(num int) ([]picture, bool) {
	var url strings.Builder
	var ret []picture
	try := 0
	date := time.Now()
	matcher := regexp.MustCompile(`image/.*\.jpg`)

	for i := 0; i < 100; i++ {
		date = date.AddDate(0, 0, -1)

		url.Reset()
		url.Grow(40)
		url.WriteString(apod)
		url.WriteString("ap")
		url.WriteString(date.Format("060102"))
		url.WriteString(".html")

	try:
		msg, ok := httpGet(url.String())
		if !ok {
			if try < 5 {
				try += 1
				goto try
			}
			continue
		}

		baseURL := matcher.FindAll(msg, -1)
		if baseURL == nil {
			continue
		}
		ret = append(ret, picture{
			Date:    date.Format("20060102"),
			BaseURL: string(baseURL[1]),
		})
	}
	return ret, true
}
