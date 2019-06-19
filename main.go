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

//setSetupChannel set setup channel to use for server status
func setSetupChannel(status chan string) chan string {
	gStatus = status
	return gStatus
}

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
	log.Println("Listening on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}



