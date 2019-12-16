package easydb

import (
	"bytes"
	"net/url"
	"testing"
)

var (
	testDataBaseName = "easydbtest"
	testKeyName      = "easykey"
)

func Test_getDbURL(t *testing.T) {
	var (
		u   *url.URL
		err error
	)

	if u, err = getDbURL(testDataBaseName); err != nil {
		t.Fatal(err)
	}

	if u.String() != "https://app.easydb.io/database/easydbtest" {
		t.Fatal("invalid database connection url")
	}
}

func Test_getDbKeyURL(t *testing.T) {
	var (
		u   *url.URL
		err error
	)

	if u, err = getDbKeyURL(testDataBaseName, testKeyName); err != nil {
		t.Fatal(err)
	}

	if u.String() != "https://app.easydb.io/database/easydbtest/easykey" {
		t.Fatal("invalid database key url")
	}
}

func Test_getJSONBody(t *testing.T) {
	var (
		err  error
		body = []byte(`{"status": "ok"}`)
	)

	var res map[string]string
	if err = getJSONBody(bytes.NewReader(body), &res); err != nil {
		t.Fatal(err)
	}

	if res["status"] != "ok" {
		t.Fatal("invalid decoded json payload")
	}
}
