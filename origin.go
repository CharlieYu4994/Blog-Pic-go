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
	"fmt"
	"io/ioutil"
	"net/http"
)

func getBing() (pics []picture, err error) {
	gURL := fmt.Sprintf("%s/HPImageArchive.aspx?format=js&idx=-1&n=9&mkt=zh-CN", domain)

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

func getAPOD() (pics []picture, err error) {
	return nil, nil
}
