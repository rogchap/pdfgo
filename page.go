package pdf

type Page struct {
	s *size

	marginLeft   float32
	marginTop    float32
	marginRight  float32
	marginBottom float32

	background drawable
	foreground drawable

	// header  element
	content drawable
	// footer  element
}

func (p *Page) Size(w, h float32) {
	p.s = &size{width: w, height: h}
}

func (p *Page) PageSize(s PageSize) {
	p.s = &size{s.width, s.height}
}

func (p *Page) Margin(left, top, right, bottom float32) {
	p.marginLeft = left
	p.marginRight = right
	p.marginTop = top
	p.marginBottom = bottom
}

func (p *Page) MarginVH(v float32) {
	p.marginLeft = v
	p.marginRight = v
	p.marginTop = v
	p.marginBottom = v
}

func (p *Page) MarginV(v float32) {
	p.marginTop = v
	p.marginBottom = v
}

func (p *Page) MarginH(v float32) {
	p.marginLeft = v
	p.marginRight = v
}

func (p *Page) build(c *container) {
	c.Layers(func(layers *Layers) {
		layers.Layer(false).Element(p.background)

		defaultSize := asSize(PageSizeA4)
		if p.s == nil {
			p.s = &defaultSize
		}

		layers.Layer(true).
			Fixed(p.s.width, p.s.height).
			Padding(p.marginLeft, p.marginTop, p.marginRight, p.marginBottom).
			Element(p.content) // TODO change to header/content/footer

		layers.Layer(false).Element(p.foreground)
	})
}

func (p *Page) Content() Container {
	c := &container{}
	p.content = c
	return c
}
