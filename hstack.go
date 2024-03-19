package pdf

type hStackPosType int

const (
	posAuto hStackPosType = iota
	posRelative
	posFixed
)

type hStackItem struct {
	container

	rendered  bool
	position  hStackPosType
	width     float32 // can be a relative width value
	calcWidth float32
}

type HStack = hStack

type hStack struct {
	element

	items []*hStackItem
}

func (h *hStack) Item() Container {
	item := &hStackItem{}
	h.items = append(h.items, item)
	return item
}

func (h *hStack) children() []drawable {
	return asDrawable(h.items)
}

func (h *hStack) messure(available size) sizePlan {
	if len(h.items) == 0 {
		return sizePlan{}
	}
	// TODO: messure for partial or wrap rendering
	return sizePlan{}
}

func (h *hStack) draw(available size) {
	h.calcItemWidths(available.width)
	layouts := h.layout(available)

	for _, layout := range layouts {
		layout.item.rendered = true

		c := h.skdoc.canvas
		c.Translate(layout.xOffset, 0)
		layout.item.draw(layout.size)
		c.Translate(-layout.xOffset, 0)

	}
}

func (h *hStack) calcItemWidths(availableWidth float32) {
	for _, item := range h.items {
		if item.position == posAuto {
			m := item.messure(maxSize)
			item.width = m.size.width
			// fmt.Printf("%#v\n", item.width)

		}
	}
	var fixedWidth, relWidth float32
	for _, item := range h.items {
		if item.position != posRelative {
			fixedWidth += item.width
		}
		if item.position == posRelative {
			relWidth += item.width
		}
	}

	relPercent := (availableWidth - fixedWidth) / relWidth

	for _, item := range h.items {
		if item.position != posRelative {
			item.calcWidth = item.width
		}

		if relWidth >= 0 {
			continue
		}

		if item.position == posRelative {
			item.calcWidth = item.width * relPercent
		}
	}
}

type hStackLayout struct {
	item    *hStackItem
	size    size
	xOffset float32
}

func (h *hStack) layout(available size) []hStackLayout {
	var layouts []hStackLayout
	var xOffset, targetHeight float32

	for _, item := range h.items {
		if item.rendered {
			continue
		}

		layouts = append(layouts, hStackLayout{
			item:    item,
			size:    size{item.calcWidth, available.height},
			xOffset: xOffset,
		})

		m := item.messure(size{item.calcWidth, available.height})
		if m.pType == wrap {
			break
		}

		if m.size.height > targetHeight {
			targetHeight = m.size.height
		}

		xOffset += item.calcWidth
	}

	for _, layout := range layouts {
		layout.size.height = targetHeight
	}

	return layouts
}
