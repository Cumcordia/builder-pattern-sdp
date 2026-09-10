package main

import (
	"fmt"
	"log"
)

func main() {
	getBuilderVar := getBuilder("GET")
	postBuilder := getBuilder("POST")

	director := newDirector(getBuilderVar)
	getUrl, err := director.buildGetURL()

	fmt.Println(getUrl.curl, getUrl.url)

	director.setBuilder(postBuilder)
	postUrl, err := director.buildPostURL()

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(postUrl.curl, postUrl.flags, postUrl.data, postUrl.url)
}
