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
	"database/sql"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type updateFunc func(num int) ([]picture, bool)

type handler struct {
	name    string
	base    string
	picNum  int
	pic     []picture
	db      *dbOperator
	update  updateFunc
	withres bool
}

func getDuration(t int) time.Duration {
	dur := time.Duration(t) * time.Hour
	now := time.Now()
	tmp := now.Format("20060102")
	today, _ := time.ParseInLocation("20060102", tmp, time.Local)
	ret := today.Add(dur + time.Second*10).Sub(now)
	if now.Hour() >= t {
		return ret + time.Hour*24
	}
	return ret
}

func newHandler(name, base string, num int, res bool, u updateFunc, db *sql.DB) (*handler, error) {
	tmp, err := newDbOperator(db, name)
	if err != nil {
		return nil, err
	}

	return &handler{
		name:    name,
		pic:     make([]picture, 0, num),
		db:      tmp,
		update:  u,
		withres: res,
	}, nil
}

func (h *handler) updatePics() bool {
	pics, ok := h.update(7)
	if !ok {
		return false
	}

	for i := range pics {
		if ok, _ := h.db.check(pics[i].Date); ok {
			h.db.insert(pics[i].Date, pics[i].BaseURL)
		}
	}
	return true
}

func (h *handler) updateBuff() bool {
	tmp, err := h.db.query(h.picNum)
	if err != nil {
		return false
	}

	for i := range tmp {
		h.pic[i] = tmp[i]
	}
	return true
}

func (h *handler) cronTask(dur int) {
	timer := time.NewTimer(0)
	for {
		<-timer.C
		h.updatePics()
		h.updateBuff()
		timer.Reset(getDuration(dur))
	}
}

func (h *handler) redirect(w http.ResponseWriter, r *http.Request) {
	var url string
	var pic picture

	args := r.URL.Query()
	dat, ok := args["dat"]
	if !ok {
		dat = append(dat, "0")
	}
	offset, err := strconv.Atoi(dat[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if offset < h.picNum && offset > -1 {
		pic = h.pic[offset]
	} else if offset == -1 {
		rand.Seed(time.Now().Unix())
		i := rand.Intn(len(h.pic) - 1)
		pic = h.pic[i]
	} else {
		pic = h.pic[h.picNum-1]
	}

	if h.withres {
		res, ok := args["res"]
		if !ok {
			res = append(res, "1920x1080")
		}

		switch res[0] {
		case "uhdres":
			url = h.base + pic.BaseURL + "_UHD.jpg"
		default:
			url = h.base + pic.BaseURL + "_" + res[0] + ".jpg"
		}
	}
	http.Redirect(w, r, url, http.StatusFound)
}
