package pdf

type drawable interface {
	children() []drawable
	messure(available size) sizePlan
	draw(available sizePlan)

	setSkDoc(skdoc *skiaDoc)
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
