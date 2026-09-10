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

func (g *PostBuilder) setCurl() {
	g.curl = "curl"
}

func (g *PostBuilder) setURL() {
	g.url = "https://example.com"
}

func (g *PostBuilder) setFlags() {
	g.flags = "-X POST -d"
}

func (g *PostBuilder) setData() {
	g.data = "data"
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
