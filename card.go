package pdf

type card struct {
	element

	header  drawable
	content drawable
	footer  drawable
}

func (c *card) children() []drawable {
	return []drawable{
		c.header,
		c.content,
		c.footer,
	}
}

func (c *card) messure(available size) sizePlan {
	layouts := c.layout(available)

	var width, height float32
	var isPartial bool
	for _, layout := range layouts {
		if layout.sPlan.size.width > width {
			width = layout.sPlan.size.width
		}
		if layout.sPlan.size.height > height {
			height = layout.sPlan.size.height
		}

		if layout.sPlan.pType == partial {
			isPartial = true
		}
	}

	if width > available.width || height > available.height {
		return sizePlan{pType: wrap}
	}

	pType := full
	if isPartial {
		pType = partial
	}

	return sizePlan{pType: pType, size: size{width, height}}
}

func (c *card) draw(available size) {
	layouts := c.layout(available)

	var width float32
	for _, layout := range layouts {
		if layout.sPlan.size.width > width {
			width = layout.sPlan.size.width
		}
	}

	for _, layout := range layouts {
		if layout.item != nil {
			s := size{width, layout.sPlan.size.height}
			c.skdoc.canvas.Translate(0, layout.yOffset)
			layout.item.draw(s)
			c.skdoc.canvas.Translate(0, -layout.yOffset)
		}
	}
}

type cardLayout struct {
	item    drawable
	sPlan   sizePlan
	yOffset float32
}

func (c *card) layout(available size) []cardLayout {
	var layouts []cardLayout

	var headerM sizePlan
	if c.header != nil {
		headerM = c.header.messure(available)
	}

	var footerM sizePlan
	if c.footer != nil {
		footerM = c.footer.messure(available)
	}

	var contentM sizePlan
	if c.content != nil {
		contentM = c.content.messure(size{available.width, available.height - headerM.size.height - footerM.size.height})
	}

	layouts = append(layouts, cardLayout{
		item:    c.header,
		sPlan:   headerM,
		yOffset: 0,
	})

	layouts = append(layouts, cardLayout{
		item:    c.content,
		sPlan:   contentM,
		yOffset: headerM.size.height,
	})

	layouts = append(layouts, cardLayout{
		item:    c.footer,
		sPlan:   footerM,
		yOffset: headerM.size.height + contentM.size.height,
	})

	return layouts
}
