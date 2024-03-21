package pdf

type Container interface {
	Element(el drawable)
	Background(color string) Container
	Padding(left, top, right, bottom float32) Container
	Fixed(width, height float32) Container

	Layers(cb func(layers *Layers))
	VStack(cb func(stack *VStack))
	HStack(cb func(stack *HStack))
	Text(s string) *TextSpan
	TextBlock(cb func(text *TextBlock))
	PageBreak()
}

type container struct {
	element

	child drawable
}

func (c *container) children() []drawable {
	return []drawable{c.child}
}

func (c *container) messure(available size) sizePlan {
	if c.child == nil {
		return sizePlan{}
	}
	return c.child.messure(available)
}

func (c *container) draw(available size) {
	if c.child == nil {
		return
	}
	c.child.draw(available)
}

// func (c *container) layers(cb func(ls *layers)) {
// 	ls := &layers{}
// 	cb(ls)
// 	// c.child = ls
// }

func (c *container) Element(el drawable) {
	c.child = el
}

func (c *container) Background(color string) Container {
	b := &background{
		color: color,
	}
	c.child = b

	return b
}

func (c *container) Padding(left, top, right, bottom float32) Container {
	if left == 0 && top == 0 && right == 0 && bottom == 0 {
		// no need to add a padding container
		return c
	}

	p := &padding{
		left:   left,
		top:    top,
		right:  right,
		bottom: bottom,
	}
	c.child = p

	return p
}

func (c *container) Fixed(width, height float32) Container {
	f := &fixedContainer{
		width:  width,
		height: height,
	}
	c.child = f

	return f
}

func (c *container) Layers(cb func(layers *Layers)) {
	ls := &Layers{}
	cb(ls)
	c.child = ls
}

func (c *container) VStack(cb func(stack *VStack)) {
	s := &vStack{}
	cb(s)
	c.child = s
}

func (c *container) HStack(cb func(stack *HStack)) {
	s := &hStack{}
	cb(s)
	c.child = s
}

func (c *container) Text(s string) *TextSpan {
	tb := &TextBlock{}
	span := tb.Span(s)
	c.child = tb
	return span
}

func (c *container) TextBlock(cb func(text *TextBlock)) {
	tb := &TextBlock{}
	cb(tb)
	c.child = tb
}

func (c *container) PageBreak() {
	c.child = &pageBreak{}
}
