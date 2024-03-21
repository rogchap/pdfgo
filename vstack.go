package pdf

type vStackItem struct {
	container

	rendered bool
}

type VStack = vStack

type vStack struct {
	element

	items []*vStackItem
}

func (v *vStack) children() []drawable {
	return asDrawable(v.items)
}

func (v *vStack) messure(available size) sizePlan {
	if len(v.items) == 0 {
		return sizePlan{}
	}
	layouts := v.layout(available)

	if len(layouts) == 0 {
		return sizePlan{pType: wrap}
	}

	var width, height float32
	var willRenderCount int
	for _, layout := range layouts {
		if layout.size.width > width {
			width = layout.size.width
		}
		if layout.size.height > height {
			height = layout.size.height
		}

		if width > available.width || height > available.height {
			return sizePlan{pType: wrap}
		}

		if layout.pType == full {
			willRenderCount++
		}
	}

	var renderedCount int
	for _, item := range v.items {
		if item.rendered {
			renderedCount++
		}
	}

	s := size{width, height}

	if willRenderCount+renderedCount == len(v.items) {
		return sizePlan{
			size: s,
		}
	}

	return sizePlan{
		pType: partial,
		size:  s,
	}
}

func (v *vStack) draw(available size) {
	if len(v.items) == 0 {
		return
	}

	layouts := v.layout(available)

	for _, layout := range layouts {
		if layout.pType == full {
			layout.item.rendered = true
		}

		c := v.skdoc.canvas
		c.Translate(0, layout.yOffset)
		layout.item.draw(size{available.width, layout.size.height})
		c.Translate(0, -layout.yOffset)
	}
}

func (v *vStack) Item() Container {
	item := &vStackItem{}
	v.items = append(v.items, item)
	return item
}

type vStackLayout struct {
	item    *vStackItem
	size    size
	yOffset float32
	pType   sizePlanType
}

func (v *vStack) layout(available size) []vStackLayout {
	var layouts []vStackLayout
	var topOffset, targetWidth float32

	for _, item := range v.items {
		if item.rendered {
			continue
		}

		availableHeight := available.height - topOffset
		if availableHeight < 0 {
			break
		}

		m := item.messure(size{available.width, availableHeight})
		if m.pType == wrap {
			break
		}

		layouts = append(layouts, vStackLayout{
			item:    item,
			size:    m.size,
			yOffset: topOffset,
			pType:   m.pType,
		})

		if m.size.width > targetWidth {
			targetWidth = m.size.width
		}

		if m.pType == partial {
			break
		}

		topOffset += m.size.height
	}

	for _, l := range layouts {
		l.size.width = targetWidth
	}

	return layouts
}
