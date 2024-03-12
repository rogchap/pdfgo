package pdf

type layer struct {
	container

	isMain bool
}

type layers struct {
	children []*layer
}

func (ls *layers) messure(available size) sizePlan {
	for _, l := range ls.children {
		if l.isMain {
			return l.messure(available)
		}
	}

	panic("No main layer defined")
}

func (ls *layers) draw(available sizePlan) {
	for _, l := range ls.children {
		m := l.messure(available.size)
		if m.pType == wrap {
			continue
		}
		l.draw(available)
	}
}

func (ls *layers) layer(main bool) *container {
	c := &container{}

	l := layer{
		isMain: main,
	}
	l.child = c

	return c
}
