package main

type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) buildGetURL() (URL, error) {
	return d.builder.setCurl().setURL().getFinalURL()
}

func (d *Director) buildPostURL() (URL, error) {
	return d.builder.setCurl().setFlags().setData().setURL().getFinalURL()
}
