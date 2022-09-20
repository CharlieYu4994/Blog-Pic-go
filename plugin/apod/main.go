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
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/dlclark/regexp2"
)

type info struct {
	PluginName string
	BaseUrl    string
	Version    string
	ResAdj     bool
}

var PluginInfo info = info{
	PluginName: "apod",
	BaseUrl:    "https://apod.nasa.gov/apod/",
	ResAdj:     false,
	Version:    "1.0.0",
}

type picture struct {
	Date    string
	BaseUrl string
}

func GetPlugInfo() info {
	return PluginInfo
}

func Update(num int) ([]picture, bool) {
	var url strings.Builder
	var ret []picture
	var tmp []byte
	date := time.Now().AddDate(0, 0, 1)
	matcher := regexp2.MustCompile(`(?<=\<a href=\")(image\/\d{4}\/.*\.\w{3,4})(?=[\s\S]*<IMG)`, 0)

	for i := 0; i < num; i++ {
		date = date.AddDate(0, 0, -1)
		url.Reset()
		url.Grow(40)
		url.WriteString(PluginInfo.BaseUrl)
		url.WriteString("ap")
		url.WriteString(date.Format("060102"))
		url.WriteString(".html")

		for t := 0; t <= 5; t++ {
			resp, err := http.Get(url.String())
			if err != nil || resp.StatusCode != 200 {
				continue
			}

			tmp, err = io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				continue
			}
		}

		if tmp == nil {
			continue
		}

		baseUrl, err := matcher.FindStringMatch(string(tmp))
		if err != nil || baseUrl == nil {
			continue
		}
		ret = append(ret, picture{
			Date:    date.Format("20060102"),
			BaseUrl: string(baseUrl.String()),
		})
	}
	return ret, true
}
