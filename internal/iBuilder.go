package main

type IBuilder interface {
	setCurl() IBuilder
	setURL() IBuilder
	setData() IBuilder
	setFlags() IBuilder
	getFinalResult() (Result, error)
}

const (
	defaultCurl  = "curl"
	defaultURL   = "https://example.com"
	defaultFlags = "-X POST -d"
	defaultData  = "data"
)

func getBuilder(builderType string) IBuilder {
	if builderType == "GET" {
		return newGetBuilder()
	}
	if builderType == "POST" {
		return newPostBuilder()
	}
	return nil
}