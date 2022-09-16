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
	"os"
)

type config struct {
	EnableTLS  bool   `json:"enabletls"`
	UpdateTime int    `json:"updatetime"`
	PicNum     int    `json:"picnum"`
	CertPath   string `json:"certpath"`
	KeyPath    string `json:"keypath"`
	Port       string `json:"port"`
	DataBase   string `json:"database"`
}

func readConf(path string, conf *config) error {
	_, err := os.Stat(path)
	if err != nil || os.IsExist(err) {
		return err
	}
	tmp, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, conf)
}
