package pdf

type fixedContainer struct {
	container

	width, height float32
}

func (f *fixedContainer) messure(available size) sizePlan {
	width, height := f.width, f.height
	if width > available.width {
		width = available.width
	}
	if height > available.height {
		height = available.height
	}

	m := f.container.messure(size{width, height})

	if m.pType == wrap {
		return sizePlan{pType: wrap}
	}

	s := size{width, height}

	if m.pType == partial {
		return sizePlan{pType: partial, size: s}
	}

	return sizePlan{size: s}
}

func (f *fixedContainer) draw(available size) {
	width, height := f.width, f.height
	if width > available.width {
		width = available.width
	}
	if height > available.height {
		height = available.height
	}

	// TODO: Canvas transate offset if RTL
	f.container.draw(size{width, height})
}
