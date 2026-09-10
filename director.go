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

func (d *Director) buildURL() URL {
	d.builder.setCurl()
	d.builder.setData()
	d.builder.setFlags()
	d.builder.setURL()
	return d.builder.getFinalURL()
}
