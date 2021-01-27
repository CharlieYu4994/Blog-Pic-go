package main

import (
	"database/sql"
	"fmt"
)

func main() {
	pURL, date, _ := getPictureInfo(1, 1, "zh-CN")
	pic := rewriteURL(pURL, date)
	db, err := sql.Open("sqlite", "./picture.db")
	if err != nil {
		return
	}
	defer db.Close()

	for _, tmp := range pic {
		fmt.Println(insertDB(db, tmp.DATE, tmp.HDURL, tmp.UHDURL))
	}
	fmt.Println(pic)
	fmt.Println(queryDB(db, "20210126"))
}
