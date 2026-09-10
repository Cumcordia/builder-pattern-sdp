package main

type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{builder: b}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) buildGetRequest() (Result, error) {
	return d.builder.
		setCurl().setURL().getFinalResult()
}

func (d *Director) buildPostRequest() (Result, error) {
	return d.builder.setCurl().setURL().setData().setFlags().getFinalResult()
}