package main

type GetBuilder struct {
	curl  string
	url   string
	data  string
	flags string
}

func newGetBuilder() *GetBuilder {
	return &GetBuilder{}
}

func (g *GetBuilder) setCurl() {
	g.curl = "curl"
}

func (g *GetBuilder) setURL() {
	g.url = "https://example.com"
}

func (g *GetBuilder) setFlags() {}

func (g *GetBuilder) setData() {}

func (g *GetBuilder) getFinalURL() URL {
	return URL{
		curl: g.curl,
		url:  g.url,
	}
}
