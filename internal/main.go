package main

import (
	"fmt"
	"log"
)

func main() {
	getBuilderVar := getBuilder("GET")
	postBuilder := getBuilder("POST")

	director := newDirector(getBuilderVar)
	getResult, err := director.buildGetRequest()
	if err != nil {
		log.Fatal(err)
	}
	if curl, ok := getResult.(CurlCommand); ok {
		fmt.Println(curl.Command)
	}

	director.setBuilder(postBuilder)
	postResult, err := director.buildPostRequest()
	if err != nil {
		log.Fatal(err)
	}
	if req, ok := postResult.(HTTPRequestObject); ok {
		fmt.Printf("%s %s\nHeaders: %v\nBody: %s\n", req.Method, req.URL, req.Headers, req.Body)
	}
}