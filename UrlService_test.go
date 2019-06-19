package main

import (
	"bytes"
	"net/http"

	"github.com/gorilla/rpc/v2/json2"
)

func (test *IntegrationTestSuite) TestParse() {
	// I am a http call to url service
	// I want to parse passed url
	// I expect to get parsed title and body

	args := []string{"vavas"}

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
		Title: 		 "title",
		Description: "description",
	}

	test.Equal(&expected, reply)

}
