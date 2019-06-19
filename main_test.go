package main

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

// Hook in test suite
func TestIntegration(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))

}

type IntegrationTestSuite struct {
	suite.Suite
}

func (test *IntegrationTestSuite) SetupSuite() {
	//The rpc client will be used across
	//all test calls within the subroutine
	//so we disable the client from closing
	//in our tests

	//Run server
	c := make(chan string)
	setSetupChannel(c)

	go main()

	//Wait for server to start
	output := <-c
	log.Println(output)
	if output != "done" {
		panic("not started")
	}

	//Add delay before sending first request
	time.Sleep(10 * time.Millisecond)
}
