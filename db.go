package main

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type inserter func(date, hdURL, uhdURL string) error
type queryer func(date string) (*picture, error)

func newInserter(db *sql.DB) (inserter, error) {
	cmd, err := db.Prepare("INSERT INTO Pictures(Date, HDURL, UHDURL)  values(?, ?, ?)")
	if err != nil {
		return nil, err
	}
	return func(date, hdURL, uhdURL string) error {
		_, err = cmd.Exec(date, hdURL, uhdURL)
		return err
	}, nil
}

func newQueryer(db *sql.DB) queryer {
	return func(date string) (*picture, error) {
		res := db.QueryRow("SELECT Date,HDURL,UHDURL FROM Pictures WHERE Date = ?", date)
		tmp := &picture{}
		err := res.Scan(&tmp.DATE, &tmp.HDURL, &tmp.UHDURL)
		if err != nil {
			return nil, err
		}
		return tmp, err
	}
}
