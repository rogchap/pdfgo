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

	container := &container{}
	if len(c.pages) == 1 {
		c.pages[0].build(container)
		return container
	}

	// TODO: Render column with page breaks
	panic("multiple pages not supported yet")
}
