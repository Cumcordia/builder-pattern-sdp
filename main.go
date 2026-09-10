package main

import "fmt"

func main() {
	getBuilderV := getBuilder("GET")
	postBuilder := getBuilder("POST")

	director := newDirector(getBuilderV)
	getUrl := director.builder.getFinalURL()

	fmt.Println(getUrl.curl)
	fmt.Println(getUrl.url)

	director.setBuilder(postBuilder)
	postUrl := director.builder.getFinalURL()

	fmt.Println(postUrl.curl)
	fmt.Println(postUrl.flags)
	fmt.Println(postUrl.data)
	fmt.Println(postUrl.url)
}
