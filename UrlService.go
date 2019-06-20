package main

import (
	"github.com/PuerkitoBio/goquery"
	ers "github.com/vavas/go-parser/errors"
	"log"
	"net/http"
	"net/url"
	"time"
)

// UrlService service to manage url content
type UrlService int

//HttpResponse details
type HttpResponse struct {
	Url    	 		string
	Title    		string
	Description    	string
	Response 		*http.Response
	Err      		error
	Status   		string
}

// Parse passed url
func (u *UrlService) Parse(r *http.Request, args *string, reply *ParsedURL) error {

	if !isValidUrl(*args) {
		return ers.ErrInvalidArgument
	}

	// async URL processing
	ch := make(chan *HttpResponse)
	go asyncHttpGet(*args, ch)
	processedURL := <-ch

	result := new(ParsedURL)
	result.Title = processedURL.Title
	result.Description = processedURL.Description

	*reply = *result

	return nil
}

//asyncHttpGet
func asyncHttpGet(url string, ch chan *HttpResponse) {

	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	// Make HTTP GET request
	resp, err := client.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var title string
	var description string

	// Find the review items
	title = doc.Find("title").Contents().Text()
	doc.Find("meta").Each(func(index int, item *goquery.Selection) {
		if item.AttrOr("name","") == "description" {
			description = item.AttrOr("content", "")
		}
	})

	processedURL := &HttpResponse{
		Url: 			url,
		Title: 			title,
		Description: 	description,
		Response: 		resp,
		Err:			err,
		Status: 		"fetched",
	}
	ch <- processedURL
	log.Println("processedURL has been sent to the chan")

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