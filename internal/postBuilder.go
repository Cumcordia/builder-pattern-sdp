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

func (p *PostBuilder) setCurl() IBuilder {
	p.curl = defaultCurl
	return p
}

func (p *PostBuilder) setURL() IBuilder {
	p.url = defaultURL
	return p
}

func (p *PostBuilder) setData() IBuilder {
	p.data = defaultData
	return p
}

func (p *PostBuilder) setFlags() IBuilder {
	p.flags = defaultFlags
	return p
}

func (p *PostBuilder) getFinalResult() (Result, error) {
	if p.url == "" {
		return nil, errors.New("PostBuilder: url is not set")
	}
	if p.data == "" {
		return nil, errors.New("PostBuilder: data is required for POST")
	}
	return HTTPRequestObject{
		Method: "POST",
		URL:    p.url,
		Headers: 
			"Content-Type application/x-www-form-urlencoded",
		Body: p.data,
	}, nil
}