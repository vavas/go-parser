package main

import (
	"net/http"
)

// UrlService service to manage url content
type UrlService int

// Parse passed url
func (u *UrlService) Parse(r *http.Request, args *string, reply *ParsedURL) error {

	result := new(ParsedURL)
	result.Title = "title"
	result.Description = "description"

	*reply = *result

	return nil
}