package main

import (
	"database/sql"
	"net/http"
	"time"
)

var dbinserter inserter
var dbquerier querier
var dbvalidator validator
var conf config
var picBuffer [7]picture

func updetePicBuffer(q querier, num int) error {
	tmp, err := q(num)
	if err != nil {
		return err
	}
	for index := 0; index < num; index++ {
		picBuffer[index] = tmp[index]
	}
	return nil
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

func timeToUpdatePic() {
	timer := time.NewTimer(0)
	for {
		<-timer.C
		updatePic(dbinserter, dbvalidator)
		updetePicBuffer(dbquerier, 7)
		timer.Reset(getDuration(conf.UpdateTime))
	}
}

func init() {
	err := readConf("./config.json", &conf)
	if err != nil {
		panic("Open Config Error")
	}

	db, err := sql.Open("sqlite", conf.DataBase)
	if err != nil {
		panic("Open Database Error")
	}
	dbinserter, err = newInserter(db)
	if err != nil {
		panic("Open Database Error")
	}
	dbquerier = newQuerier(db)
	dbvalidator = newValidator(db)

	go timeToUpdatePic()
}

func main() {
	http.HandleFunc("/bing", redirectToBing)
	time.Sleep(time.Second)
	if conf.EnableTLS {
		http.ListenAndServeTLS("0.0.0.0:"+conf.Port,
			conf.CertPath, conf.KeyPath, nil)
	} else {
		http.ListenAndServe("0.0.0.0:"+conf.Port, nil)
	}
}
