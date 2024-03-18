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
	// TODO: messure for partial or wrap rendering
	return sizePlan{}
}

func (v *vStack) draw(available size) {
	layouts := v.layout(available)

	for _, layout := range layouts {
		layout.item.rendered = true

		c := v.skdoc.canvas
		c.Translate(0, layout.yOffset)
		layout.item.draw(size{available.width, layout.size.height})
		c.Translate(0, -layout.yOffset)
	}
}

func (v *vStack) Item() *Container {
	item := &vStackItem{}
	v.items = append(v.items, item)
	return &item.container
}

type vStackLayout struct {
	item    *vStackItem
	size    size
	yOffset float32
}

func (v *vStack) layout(available size) []vStackLayout {
	var layouts []vStackLayout
	var topOffset, targetWidth float32

	for _, item := range v.items {
		if item.rendered {
			continue
		}

		availableHeight := available.height - topOffset
		if availableHeight <= 0 {
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
