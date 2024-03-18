package pdf

type Layer struct {
	container

	isMain bool
}

type Layers struct {
	element

	items []*Layer
}

func (ls *Layers) children() []drawable {
	return asDrawable(ls.items)
}

func (ls *Layers) messure(available size) sizePlan {
	for _, l := range ls.items {
		if l.isMain {
			return l.messure(available)
		}
	}

	panic("No main layer defined")
}

func (ls *Layers) draw(available size) {
	for _, l := range ls.items {
		m := l.messure(available)
		if m.pType == wrap {
			continue
		}
		l.draw(available)
	}
}

func (ls *Layers) Layer(main bool) *Container {
	c := &container{}

	l := Layer{
		isMain: main,
	}
	l.child = c
	ls.items = append(ls.items, &l)

	return c
}
