package main

import "errors"

type GetBuilder struct {
	curl  string
	url   string
	data  string
	flags string
}

func newGetBuilder() *GetBuilder {
	return &GetBuilder{}
}

func (g *GetBuilder) setCurl() IBuilder {
	g.curl = Ccurl
	return  g
}

func (g *GetBuilder) setURL() IBuilder {
	g.url = Curl
	return g
}

func (g *GetBuilder) setFlags() IBuilder {
	return g
}

func (g *GetBuilder) setData() IBuilder {
	return g
}

func (g *GetBuilder) getFinalURL() (URL, error) {
	if g.curl == "" || g.url == "" {
		return URL{}, errors.New("URL is null")
	}

	return URL{
		curl: g.curl,
		url:  g.url,
	}, nil
}
