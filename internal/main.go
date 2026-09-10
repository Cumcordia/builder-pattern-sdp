package main

import "fmt"

func main() {
	getBuilderVar := getBuilder("GET")
	postBuilder := getBuilder("POST")

	director := newDirector(getBuilderVar)
	getUrl := director.buildGetURL()

	fmt.Print(getUrl.curl + " ")
	fmt.Println(getUrl.url)

	director.setBuilder(postBuilder)
	postUrl := director.buildPostURL()

	fmt.Print(postUrl.curl  + " ")
	fmt.Print(postUrl.flags  + " ")
	fmt.Print(postUrl.data  + " ")
	fmt.Println(postUrl.url)
}
