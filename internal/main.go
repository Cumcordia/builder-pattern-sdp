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

	fmt.Print(getUrl.curl + " ")
	fmt.Println(getUrl.url)

	director.setBuilder(postBuilder)
	postUrl, err := director.buildPostURL()

	if err != nil{
		log.Fatal(err)
	}

	fmt.Print(postUrl.curl  + " ")
	fmt.Print(postUrl.flags  + " ")
	fmt.Print(postUrl.data  + " ")
	fmt.Println(postUrl.url)
}
