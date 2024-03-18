package pdf

type drawable interface {
	children() []drawable
	messure(available size) sizePlan
	draw(available size)

	setSkDoc(skdoc *skiaDoc)
}

func asDrawable[T drawable](s []T) []drawable {
	rtn := make([]drawable, len(s))
	for i := range s {
		rtn[i] = s[i]
	}
	return rtn
}

type element struct {
	skdoc *skiaDoc
}

func (e *element) children() []drawable {
	return nil
}

func (e *element) setSkDoc(skdoc *skiaDoc) {
	e.skdoc = skdoc
}
