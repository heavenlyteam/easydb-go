package easydb

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

type Client struct {
	databaseName    string
	connectionToken string

	hc http.Client
}

func Open(databaseName, token string) (ed *Client, err error) {
	var easyDB Client

	if len(token) == 0 {
		err = ErrEmptyToken
		return
	}

	if len(databaseName) == 0 {
		err = ErrEmptyDB
		return
	}

	easyDB.databaseName = databaseName
	easyDB.connectionToken = token

	// Pointer binding
	ed = &easyDB
	return
}

func (c *Client) query(method string, u *url.URL, body []byte) (resp *http.Response, err error) {
	var req *http.Request
	if req, err = http.NewRequest(method, u.String(), bytes.NewReader(body)); err != nil {
		return
	}

	if method == http.MethodPost {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Add("token", c.connectionToken)

	if resp, err = c.hc.Do(req); err != nil {
		return
	}

	return
}

func (c *Client) Get(key string) (result interface{}, err error) {
	var (
		u    *url.URL
		resp *http.Response
	)

	if u, err = getDbKeyURL(c.databaseName, key); err != nil {
		return
	}

	if resp, err = c.query(http.MethodGet, u, nil); err != nil {
		return
	}
	defer resp.Body.Close()

	if err = getJSONBody(resp.Body, &result); err != nil {
		return
	}
	return
}

func (c *Client) List() (result []interface{}, err error) {
	var (
		u    *url.URL
		resp *http.Response
	)
	if u, err = getDbURL(c.databaseName); err != nil {
		return
	}

	if resp, err = c.query(http.MethodGet, u, nil); err != nil {
		return
	}
	defer resp.Body.Close()

	if err = getJSONBody(resp.Body, &result); err != nil {
		return
	}
	return
}

func (c *Client) Put(key string, value interface{}) (err error) {
	var (
		u    *url.URL
		resp *http.Response
		body []byte
	)

	if u, err = getDbKeyURL(c.databaseName, key); err != nil {
		return
	}

	if body, err = json.Marshal(map[string]interface{}{
		"value": value,
	}); err != nil {
		return
	}

	if resp, err = c.query(http.MethodPost, u, body); err != nil {
		return
	}
	defer resp.Body.Close()

	return
}

func (c *Client) Delete(key string) (err error) {
	var (
		u    *url.URL
		resp *http.Response
	)

	if u, err = getDbKeyURL(c.databaseName, key); err != nil {
		return
	}

	if resp, err = c.query(http.MethodDelete, u, nil); err != nil {
		return
	}
	defer resp.Body.Close()

	return
}
