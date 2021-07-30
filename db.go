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
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type insertFunc func(date, baseUrl string) error
type queryFunc func(num int) ([]picture, error)
type checkFunc func(date string) (bool, error)

type dbOperator struct {
	insert insertFunc
	query  queryFunc
	check  checkFunc
}

func newDbOperator(db *sql.DB, table string) (*dbOperator, error) {
	insertCmd, err := db.Prepare(
		fmt.Sprintf("INSERT INTO %s(DATE, BURL)  values(?, ?)", table))
	if err != nil {
		return nil, err
	}

	queryCmd, err := db.Prepare(
		fmt.Sprintf("SELECT DATE,BURL FROM %s ORDER BY id DESC LIMIT ?", table))
	if err != nil {
		return nil, err
	}

	checkCmd, err := db.Prepare(
		fmt.Sprintf(`SELECT IFNULL((SELECT Date FROM %s WHERE Date=?), "NULL")`, table))
	if err != nil {
		return nil, err
	}

	return &dbOperator{
		insert: func(date, baseUrl string) error {
			_, err := insertCmd.Exec(date, baseUrl)
			return err
		},

		query: func(num int) ([]picture, error) {
			var pic picture
			tmp := make([]picture, 0, num)

			result, err := queryCmd.Query(num)
			if err != nil {
				return nil, err
			}

			for result.Next() {
				err := result.Scan(&pic.Date, &pic.BaseUrl)
				if err != nil {
					return nil, err
				}
				tmp = append(tmp, pic)
			}

			if num-len(tmp) > 0 {
				for i := 0; i < num-len(tmp); i++ {
					tmp = append(tmp, tmp[i])
				}
			}
			return tmp, nil
		},

		check: func(date string) (bool, error) {
			var tmp string

			result := checkCmd.QueryRow(date)
			err := result.Scan(&tmp)
			if err == nil {
				if tmp == "NULL" {
					return true, nil
				}
				return false, nil
			}
			return false, err
		},
	}, nil
}
