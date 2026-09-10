package main

type IBuilder interface {
	setCurl()
	setURL()
	setData()
	setFlags()
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
