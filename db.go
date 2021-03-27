package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type inserter func(date, BURL string) error
type querier func(n int) ([]picture, error)
type validator func(date string) (bool, error)

func newInserter(db *sql.DB) (inserter, error) {
	cmd, err := db.Prepare("INSERT INTO BING(DATE, BURL)  values(?, ?)")
	if err != nil {
		return nil, err
	}
	return func(date, BURL string) error {
		_, err = cmd.Exec(date, BURL)
		return err
	}, nil
}

func newQuerier(db *sql.DB) querier {
	return func(n int) ([]picture, error) {
		res, err := db.Query(
			"SELECT DATE,BURL FROM BING ORDER BY id DESC LIMIT ?", n)
		if err != nil {
			return nil, err
		}

		ret := make([]picture, 0, n)
		var tmp picture
		for res.Next() {
			err = res.Scan(&tmp.Date, &tmp.Burl)
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
			`SELECT IFNULL((SELECT Date FROM BING WHERE Date=?), "NULL")`, date)
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
