package pdf

import (
	"rogchap.com/skia"
)

type drawable interface {
	children() []drawable
	messure(available size) sizePlan
	draw(available size)

	setSkDoc(skdoc *skiaDoc)
	setPageCtx(pCtx *pageContext)
}

type resetable interface {
	reset()
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
	pCtx  *pageContext
}

func (e *element) children() []drawable {
	return nil
}

func (e *element) setSkDoc(skdoc *skiaDoc) {
	e.skdoc = skdoc
}

func (e *element) setPageCtx(pCtx *pageContext) {
	e.pCtx = pCtx
}

func (e *element) canvas() *skia.Canvas {
	return e.skdoc.canvas
}
