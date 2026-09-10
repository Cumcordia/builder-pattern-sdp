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
	d.builder.setCurl()
	d.builder.setURL()
	return d.builder.getFinalURL()
}

func (d *Director) buildPostURL() (URL, error) {
	d.builder.setCurl()
	d.builder.setFlags()
	d.builder.setData()
	d.builder.setURL()
	return d.builder.getFinalURL()
}
