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
	"flag"
	"net/http"
	"sync"
)

const bing string = "https://cn.bing.com"
const apod string = "https://apod.nasa.gov/apod/"

var conf config
var confpath string
var wg sync.WaitGroup

var bingHandler *handler
var apodHandler *handler

func init() {
	flag.StringVar(&confpath, "c", "./config.json", "Set the config path")

	flag.Parse()

	err := readConf(confpath, &conf)
	if err != nil {
		panic("OpenConfigError")
	}

	db, err := sql.Open("sqlite3", conf.DataBase)
	if err != nil {
		panic("OpenDatabaseError")
	}

	bingHandler, err = newHandler("bing", bing, conf.PicNum, true, getBing, db)
	if err != nil {
		panic("CreateHandlerError")
	}

	apodHandler, err = newHandler("apod", apod, conf.PicNum, false, getAPOD, db)
	if err != nil {
		panic("CreateHandlerError")
	}

	wg.Add(2)
	go bingHandler.updateTask(conf.UpdateTime, &wg)
	go apodHandler.updateTask(conf.UpdateTime, &wg)
}

func main() {
	http.HandleFunc(bingHandler.path, bingHandler.redirect)
	http.HandleFunc(apodHandler.path, apodHandler.redirect)

	wg.Wait()

	if conf.EnableTLS {
		http.ListenAndServeTLS("0.0.0.0:"+conf.Port,
			conf.CertPath, conf.KeyPath, nil)
	} else {
		http.ListenAndServe("0.0.0.0:"+conf.Port, nil)
	}
}
