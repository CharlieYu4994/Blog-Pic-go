package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type inserter func(date, hdURL, uhdURL string) error
type querier func(n int) ([]picture, error)
type validator func(date string) (bool, error)

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

func newQuerier(db *sql.DB) querier {
	return func(n int) ([]picture, error) {
		res, err := db.Query(
			"SELECT Date,HDURL,UHDURL FROM Pictures ORDER BY id DESC LIMIT ?", n)
		if err != nil {
			return nil, err
		}

		ret := make([]picture, 0, n)
		tmp := picture{}
		for res.Next() {
			err = res.Scan(&tmp.DATE, &tmp.HDURL, &tmp.UHDURL)
			if err != nil {
				return nil, err
			}
			ret = append(ret, tmp)
		}

		if n-len(ret) > 0 {
			for t := n - len(ret); t > 0; t-- {
				ret = append(ret, tmp)
			}
		}

		return ret, nil
	}
}

func newValidator(db *sql.DB) validator {
	return func(date string) (bool, error) {
		res := db.QueryRow(
			`SELECT IFNULL((SELECT Date FROM Pictures WHERE Date=?), "NULL")`, date)
		var tmp string
		err := res.Scan(&tmp)
		if err == nil {
			if tmp == "NULL" {
				return true, nil
			}
			return false, nil
		}
		return false, err
	}
}
