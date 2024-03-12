package pdf

type Page struct {
	s *size

	// background element
	// foreground element

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
	c.child = p.content

	// c.layers(func(ls *layers) {
	// 	ls.layer(false).child = p.background
	// 	ls.layer(true).child = p.content // TODO change to header/content/footer
	// 	ls.layer(false).child = p.foreground
	// })
}

func (p *Page) Content() *Container {
	c := &container{}
	p.content = c
	return c
}
