package pdf

type Document interface {
	Build(c *DocContainer)
}

type DocContainer struct {
	pages []*Page
}

func (c *DocContainer) Page(cb func(page *Page)) {
	page := &Page{}
	cb(page)
	c.pages = append(c.pages, page)
}

func (c *DocContainer) build() *container {
	if len(c.pages) == 0 {
		return nil
	}

	cont := &container{}
	if len(c.pages) == 1 {
		c.pages[0].build(cont)
		return cont
	}

	cont.VStack(func(stack *VStack) {
		for idx, page := range c.pages {
			if idx != 0 {
				stack.Item().PageBreak()
			}
			pageCont := &container{}
			page.build(pageCont)
			stack.Item().Element(pageCont)
		}
	})
	return cont
}
