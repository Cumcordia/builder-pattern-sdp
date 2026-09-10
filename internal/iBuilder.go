package main

type IBuilder interface {
	setCurl() IBuilder
	setURL() IBuilder
	setData() IBuilder
	setFlags() IBuilder
	getFinalURL() (URL, error)
}

func getBuilder(builderType string) IBuilder {
	if builderType == "GET" {
		return newGetBuilder()
	}

	if builderType == "POST" {
		return newPostBuilder()
	}

	return nil
}
