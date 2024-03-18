package pdf

type Page struct {
	s *size

	background drawable
	foreground drawable

	// header  element
	content drawable
	// footer  element
}

func (p *Page) Size(w, h float32) {
	p.s = &size{width: w, height: h}
}

func (p *Page) PageSize(s size) {
	p.s = &s
}

func (p *Page) build(c *container) {
	c.Layers(func(layers *Layers) {
		layers.Layer(false).Element(p.background)
		layers.Layer(true).Element(p.content) // TODO change to header/content/footer
		layers.Layer(false).Element(p.foreground)
	})
}

func (p *Page) Content() *Container {
	c := &container{}
	p.content = c
	return c
}
