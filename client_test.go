package easydb

import (
	"errors"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	Status string `json:"status"`
}

func TestOpen(t *testing.T) {
	var (
		err error
	)

	if _, err = Open("", ""); err == nil {
		t.Fatal("an error occurred. Provided empty credentials. Should be an error")
	}

	if _, err = Open("testDatabase", "testToken"); err != nil {
		t.Fatal("an error occurred. Provided non empty credentials. Should not be an error")
	}
}

func TestEasyDB_Get(t *testing.T) {
	var (
		err        error
		db         *Client
		assertions = assert.New(t)
	)

	if db, err = getInstance(); err != nil {
		t.Fatal(err)
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		description   string
		responder     httpmock.Responder
		expected      interface{}
		expectedError error
	}{
		{
			description:   "Success result on 'Get'",
			responder:     httpmock.NewStringResponder(http.StatusOK, `{"status": "ok"}`),
			expected:      map[string]interface{}{"status": "ok"},
			expectedError: nil,
		},
		{
			description:   "Error response on 'Get'",
			responder:     httpmock.NewStringResponder(http.StatusOK, ``),
			expected:      nil,
			expectedError: errors.New("EOF"),
		},
	}

	for _, tc := range tests {
		httpmock.RegisterResponder(http.MethodGet, "https://app.easydb.io/database/test/test", tc.responder)
		r, err := db.Get("test")

		assertions.Equal(r, tc.expected, tc.description)
		assertions.Equal(err, tc.expectedError, tc.description)
	}
}

func TestEasyDB_Put(t *testing.T) {
	var (
		err        error
		db         *Client
		assertions = assert.New(t)
	)

	if db, err = getInstance(); err != nil {
		t.Fatal(err)
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		description   string
		responder     httpmock.Responder
		payload       interface{}
		expectedError error
	}{
		{
			description:   "Success result on 'Put'",
			responder:     httpmock.NewStringResponder(http.StatusOK, ``),
			payload:       "test",
			expectedError: nil,
		},
		{
			description: "Error response on 'Get'",
			responder:   httpmock.NewStringResponder(http.StatusOK, ``),
			payload: struct {
				Status string
			}{
				Status: "success",
			},
			expectedError: nil,
		},
	}

	for _, tc := range tests {
		httpmock.RegisterResponder(http.MethodPost, "https://app.easydb.io/database/test/test", tc.responder)
		err := db.Put("test", tc.payload)

		assertions.Equal(err, tc.expectedError, tc.description)
	}
}

func TestEasyDB_Delete(t *testing.T) {
	var (
		err        error
		db         *Client
		assertions = assert.New(t)
	)

	if db, err = getInstance(); err != nil {
		t.Fatal(err)
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		description   string
		responder     httpmock.Responder
		expectedError error
	}{
		{
			description:   "Success result on 'Delete'",
			responder:     httpmock.NewStringResponder(http.StatusOK, ``),
			expectedError: nil,
		},
	}

	for _, tc := range tests {
		httpmock.RegisterResponder(http.MethodDelete, "https://app.easydb.io/database/test/test", tc.responder)
		err := db.Delete("test")

		assertions.Equal(err, tc.expectedError, tc.description)
	}
}

func TestEasyDB_List(t *testing.T) {
	var (
		err        error
		db         *Client
		assertions = assert.New(t)
	)

	if db, err = getInstance(); err != nil {
		t.Fatal(err)
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		description   string
		responder     httpmock.Responder
		expectedError error
		expected      interface{}
	}{
		{
			description:   "Success result on 'List'",
			responder:     httpmock.NewStringResponder(http.StatusOK, `[{"Status": "ok"}]`),
			expectedError: nil,
			expected: testStruct{
				Status: "ok",
			},
		},
		{
			description:   "Error response on 'List'",
			responder:     httpmock.NewStringResponder(http.StatusOK, ``),
			expectedError: errors.New("EOF"),
			expected:      nil,
		},
	}

	for _, tc := range tests {
		httpmock.RegisterResponder(http.MethodGet, "https://app.easydb.io/database/test", tc.responder)
		r, err := db.List()

		if len(r) > 0 {
			var res testStruct
			err = mapstructure.Decode(r[0], &res)
			assertions.Equal(res, tc.expected, tc.description)
		}

		assertions.Equal(err, tc.expectedError, tc.description)
	}
}

func getInstance() (db *Client, err error) {
	return Open("test", "test")
}
