package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

// UrlService service to manage url content
type UrlService int

//HttpResponse details
type HttpResponse struct {
	Url    	 string
	Response *http.Response
	Err      error
	Status   string
}

// Parse passed url
func (u *UrlService) Parse(r *http.Request, args *string, reply *ParsedURL) error {


	if !isValidUrl(*args) {
		// TODO implement proper GOPATH
	}


	log.Println(*args)
	//ch := make(chan *HttpResponse)
	result := new(ParsedURL)
	result.Title = "title"
	result.Description = "description"

	*reply = *result

	return nil
}

//asyncHttpGet
func asyncHttpGet(url string, ch chan *HttpResponse) {

	timeout := time.Duration(500 * time.Millisecond)
	client := http.Client{
		Timeout: timeout,
	}

	log.Printf("Fetching %s \n", url)
	resp, err := client.Get(url)
	fmt.Printf("%+v\n",resp.Header)
	fmt.Printf("%+v\n",resp.Body)

	u := &HttpResponse{
		Url: 		url,
		Response: 	resp,
		Err:		err,
		Status: 	"fetched",
	}
	ch <- u
	log.Println("sent to chan")

}

// isValidUrl tests a string to determine if it is a url or not.
func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	} else {
		return true
	}
}