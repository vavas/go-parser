package main

import (
	"bytes"
	"github.com/gorilla/rpc/v2/json2"
	"net/http"

	ers "github.com/vavas/go-parser/errors"
)

func (test *IntegrationTestSuite) TestParseSuccess() {
	// I am a http call to url service
	// I want to parse passed url
	// I expect to get parsed title and body

	args := []string{"https://github.com/vavas/go-parser"}

	json, err := json2.EncodeClientRequest("url.Parse", args)
	test.NoError(err)

	client := http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8000/rpc", bytes.NewReader(json))
	test.NoError(err)
	resp, err := client.Do(req)
	test.NoError(err)

	defer resp.Body.Close()
	reply := new(ParsedURL)
	err = json2.DecodeClientResponse(resp.Body, &reply)
	test.NoError(err)

	expected := ParsedURL{
		Title: 		 "GitHub - vavas/go-parser",
		Description: "Contribute to vavas/go-parser development by creating an account on GitHub.",
	}

	test.Equal(&expected, reply)

}

func (test *IntegrationTestSuite) TestParseInvalidUrl() {
	// I am a http call to url service
	// I expect to get get invalid URL error

	args := []string{"invalid url"}

	json, err := json2.EncodeClientRequest("url.Parse", args)
	test.NoError(err)

	client := http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8000/rpc", bytes.NewReader(json))
	test.NoError(err)
	resp, err := client.Do(req)
	test.NoError(err)

	defer resp.Body.Close()
	err = json2.DecodeClientResponse(resp.Body, new(string))
	test.Equal("Invalid arguments", ers.ErrInvalidArgument.Message)

}

func (test *IntegrationTestSuite) TestParseForbidden() {
	// I am a http call to url service
	// I expect to get forbidden url error

	args := []string{"https://httpstat.us/403"}

	json, err := json2.EncodeClientRequest("url.Parse", args)
	test.NoError(err)

	client := http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8000/rpc", bytes.NewReader(json))
	test.NoError(err)
	resp, err := client.Do(req)
	test.NoError(err)

	defer resp.Body.Close()
	err = json2.DecodeClientResponse(resp.Body, new(string))
	test.Equal("Forbidden URL", ers.ErrForbiddenURL.Message)

}


