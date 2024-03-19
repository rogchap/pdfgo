package pdf

type padding struct {
	container

	left   float32
	top    float32
	right  float32
	bottom float32
}

func (p *padding) messure(available size) sizePlan {
	if p.child == nil {
		return sizePlan{}
	}

	s := size{
		width:  available.width - p.left - p.right,
		height: available.height - p.top - p.bottom,
	}

	m := p.container.messure(s)

	return sizePlan{
		pType: m.pType,
		size: size{
			width:  m.size.width + p.left + p.right,
			height: m.size.height + p.top + p.bottom,
		},
	}
}

func (p *padding) draw(available size) {
	if p.child == nil {
		return
	}

	s := size{
		width:  available.width - p.left - p.right,
		height: available.height - p.top - p.bottom,
	}

	c := p.skdoc.canvas
	c.Translate(p.left, p.top)
	p.child.draw(s)
	c.Translate(-p.left, -p.top)
}
