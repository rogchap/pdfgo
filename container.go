package pdf

import "fmt"

type Container = container

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

func (c *container) draw(sp sizePlan) {
	fmt.Printf("%#v\n", c.child)
	if c.child == nil {
		return
	}
	c.child.draw(sp)
}

// func (c *container) layers(cb func(ls *layers)) {
// 	ls := &layers{}
// 	cb(ls)
// 	// c.child = ls
// }

func (c *container) Background(color string) *Container {
	b := &background{
		color: color,
	}
	c.child = b

	return &b.container
}

func (c *container) Text(cb func(text *TextBlock)) {
	tb := &TextBlock{}
	cb(tb)
	c.child = tb
}
