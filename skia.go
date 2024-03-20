package pdf

import (
	"io"

	"rogchap.com/skia"
)

type skiaDoc struct {
	doc    *skia.Document
	canvas *skia.Canvas

	pageSize size
}

func newSkiaDoc(w io.Writer) *skiaDoc {
	return &skiaDoc{
		doc: skia.NewDocument(w),
	}
}

func (sk *skiaDoc) beginPage(sp sizePlan) {
	sk.pageSize = sp.size
	sk.canvas = sk.doc.BeginPage(sp.size.width, sp.size.height, nil)
}

func (sk *skiaDoc) endPage() {
	sk.doc.EndPage()
	// TODO: Dispose of skia.Canvas
	sk.canvas = nil
}

func (sk *skiaDoc) endDoc() {
	sk.doc.Close()
	sk.doc.Dispose()
	sk.doc = nil
}
