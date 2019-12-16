package easydb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
)

const BaseUrl = "https://app.easydb.io"

func getDbURL(database string) (u *url.URL, err error) {
	if u, err = url.Parse(BaseUrl); err != nil {
		return
	}

	u.Path = fmt.Sprintf("database/%s", database)
	return
}

func getDbKeyURL(database, key string) (u *url.URL, err error) {
	if u, err = getDbURL(database); err != nil {
		return
	}

	u.Path = fmt.Sprintf("database/%s/%s", database, key)
	return
}

func getJSONBody(body io.Reader, res interface{}) (err error) {
	if err = json.NewDecoder(body).Decode(&res); err != nil {
		return
	}

	return
}
