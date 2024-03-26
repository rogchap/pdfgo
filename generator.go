package pdf

import (
	"io"

	"rogchap.com/skia"
)

func Generate(w io.Writer, builder DocBuilder) {
	doc := &Doc{
		Producer: "rogchap.com/skia",
	}
	builder.Build(doc)
	content := doc.build()

	metadata := skia.PDFMetadata{
		Title:           doc.Title,
		Author:          doc.Author,
		Subject:         doc.Subject,
		Keywords:        doc.Keywords,
		Creator:         doc.Creator,
		Producer:        doc.Producer,
		Creation:        doc.Creation,
		Modified:        doc.Modified,
		PDFA:            doc.PDFA,
		RasterDPI:       doc.RasterDPI,
		EncodingQuality: doc.EncodingQuality,
	}

	skdoc := newSkiaDoc(w, metadata)

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
