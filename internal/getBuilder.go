package main

import "errors"

type GetBuilder struct {
	curl string
	url  string
}

func newGetBuilder() *GetBuilder {
	return &GetBuilder{}
}

func (g *GetBuilder) setCurl() IBuilder {
	g.curl = defaultCurl
	return g
}

func (g *GetBuilder) setURL() IBuilder {
	g.url = defaultURL
	return g
}

// GET не использует data/flags — методы нужны только для соответствия интерфейсу.
func (g *GetBuilder) setData() IBuilder {
	return g
}

func (g *GetBuilder) setFlags() IBuilder {
	return g
}

func (g *GetBuilder) getFinalResult() (Result, error) {
	if g.curl == "" || g.url == "" {
		return nil, errors.New("GetBuilder: curl or url is not set")
	}
	return CurlCommand{
		Command: g.curl + " " + g.url,
	}, nil
}