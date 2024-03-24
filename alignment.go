package pdf

type alignType int

const (
	alignLeft alignType = iota
	alignCenter
	alignRight
)

type alignment struct {
	container

	align alignType
}

func (a *alignment) draw(available size) {
	if a.child == nil {
		return
	}

	m := a.messure(available)
	if m.pType == wrap {
		return
	}

	diff := available.width - m.size.width
	var xOffset float32
	switch a.align {
	case alignCenter:
		xOffset = diff * 0.5
	case alignRight:
		xOffset = diff
	default:
		xOffset = 0
	}

	c := a.skdoc.canvas
	c.Translate(xOffset, 0)
	a.container.draw(m.size)
	c.Translate(-xOffset, 0)
}
