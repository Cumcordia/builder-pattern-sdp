package main

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

func (g *PostBuilder) getFinalURL() URL {
	return URL{
		curl:  g.curl,
		url:   g.url,
		data:  g.data,
		flags: g.flags,
	}
}
