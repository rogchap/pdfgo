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

	return sizePlan{size: size{width, height}}
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
