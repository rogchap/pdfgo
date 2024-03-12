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
	})

	for {
		pctx.Inc()
		sp := content.messure(maxSize)
		// TODO: check if plan is wrap and then panic

		// TODO: Temp fix until we do proper page messuring
		sp.size = pageSizeA4

		skdoc.beginPage(sp)
		content.draw(sp)
		skdoc.endPage()

		if sp.pType == full {
			break
		}
	}

	// TODO: messure Space Plan to deal with wrapping content

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
