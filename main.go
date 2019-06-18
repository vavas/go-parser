package main

import (

	"log"
	"net/http"

	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json"

)

var (
	gStatus chan string
)

//getSetupChannel get setup channel for server status
func getSetupChannel() chan string {
	return gStatus
}

func main() {

	log.Println("Loading services...")

	// rpc server for project endpoint
	s := rpc.NewServer()
	s.RegisterService(new(UrlService), "url")
	s.RegisterCodec(json.NewCodec(), "application/json")

	http.Handle("/rpc", s)

	ch := getSetupChannel()
	if ch != nil {
		ch <- "done"
	}

	log.Println("Starting listener...")
	log.Println("Listening on port 80...")
	log.Fatal(http.ListenAndServe(":80", nil))
}



