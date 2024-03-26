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

	space float32
	items []*hStackItem
}

func (h *hStack) Space(s float32) {
	h.space = s
}

func (h *hStack) RelativeItem(size float32) Container {
	if size < 1 {
		size = 1
	}
	item := &hStackItem{
		position: posRelative,
		width:    size,
	}
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

	h.calcItemWidths(available.width)
	layouts := h.layout(available)

	lastLayout := layouts[len(layouts)-1]
	width := lastLayout.xOffset + lastLayout.size.width
	var height float32
	for _, layout := range layouts {

		if layout.ptype == wrap {
			return sizePlan{pType: wrap}
		}

		if layout.size.height > height {
			height = layout.size.height
		}
	}

	if width > available.width || height > available.height {
		return sizePlan{pType: wrap}
	}

	s := size{width, height}
	for _, layout := range layouts {
		if !layout.item.rendered && layout.ptype == partial {
			return sizePlan{pType: partial, size: s}
		}
	}

	return sizePlan{size: s}
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
	spacing := float32(len(h.items)-1) * h.space
	relPercent := (availableWidth - fixedWidth - spacing) / relWidth

	for _, item := range h.items {
		if item.position != posRelative {
			item.calcWidth = item.width
		}

		if relWidth <= 0 {
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
	ptype   sizePlanType
}

func (h *hStack) layout(available size) []hStackLayout {
	var layouts []hStackLayout
	var xOffset, targetHeight float32

	for _, item := range h.items {
		if item.rendered {
			continue
		}

		m := item.messure(size{item.calcWidth, available.height})
		layouts = append(layouts, hStackLayout{
			item:    item,
			size:    size{item.calcWidth, available.height},
			xOffset: xOffset,
			ptype:   m.pType,
		})

		if m.pType == wrap {
			break
		}

		if m.size.height > targetHeight {
			targetHeight = m.size.height
		}

		xOffset += item.calcWidth + h.space
	}

	for idx := range layouts {
		layouts[idx].size.height = targetHeight
	}

	return layouts
}
