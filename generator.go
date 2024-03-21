package pdf

import (
	"io"
)

func Generate(w io.Writer, doc Document) {
	skdoc := newSkiaDoc(w)

	container := &DocContainer{}
	doc.Build(container)

	content := container.build()

	pctx := &pageContext{}
	_ = pctx

	walkChildren(content, func(el drawable) {
		if el == nil {
			return
		}

		el.setSkDoc(skdoc)
		// TODO: set page context

		if el, ok := el.(resetable); ok {
			el.reset()
		}
	})

	for {
		pctx.Inc()
		sp := content.messure(maxSize)

		skdoc.beginPage(sp)
		content.draw(sp.size)
		skdoc.endPage()

		if sp.pType == full {
			break
		}
	}

	skdoc.endDoc()
}

func walkChildren(el drawable, cb func(el drawable)) {
	if el == nil {
		return
	}

	for _, c := range el.children() {
		walkChildren(c, cb)
	}

	cb(el)
}
