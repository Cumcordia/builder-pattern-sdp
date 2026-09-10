package main

type Result interface {
	isResult()
}

type CurlCommand struct {
	Command string
}

func (CurlCommand) isResult() {}

type HTTPRequestObject struct{
	Method string
	URL string
	Headers string
	Body string
}

func (HTTPRequestObject) isResult(){}