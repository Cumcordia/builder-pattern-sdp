package main

import "errors"

type PostBuilder struct {
	curl  string
	url   string
	data  string
	flags string
}

func newPostBuilder() *PostBuilder {
	return &PostBuilder{}
}

func (g *PostBuilder) setCurl() IBuilder{
	g.curl = "curl"
	return g
}

func (g *PostBuilder) setURL() IBuilder {
	g.url = "https://example.com"
	return g
}

func (g *PostBuilder) setFlags() IBuilder{
	g.flags = "-X POST -d"
	return g
}

func (g *PostBuilder) setData() IBuilder{
	g.data = "data"
	return g
}

func (g *PostBuilder) getFinalURL() (URL, error) {

	if g.curl == "" || g.url == "" || g.flags == "" || g.data == "" {
		return URL{}, errors.New("URL is null")
	}
	return URL{
		g.curl,
		g.flags,
		g.data,
		g.url,
	}, nil
}
