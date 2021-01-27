package main

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func insertDB(db *sql.DB, date, hdURL, uhdURL string) error {
	cmd, err := db.Prepare("INSERT INTO Pictures(Date, HDURL, UHDURL)  values(?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = cmd.Exec(date, hdURL, uhdURL)
	if err != nil {
		cmd.Close()
		return err
	}
	cmd.Close()
	return nil
}

func queryDB(db *sql.DB, date string) (*picture, error) {
	res := db.QueryRow("SELECT Date,HDURL,UHDURL FROM Pictures WHERE Date = ?", date)
	tmp := &picture{}
	err := res.Scan(&tmp.DATE, &tmp.HDURL, &tmp.UHDURL)
	if err != nil {
		return nil, err
	}
	return tmp, err
}
