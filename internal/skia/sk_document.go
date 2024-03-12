package skia

//#include "sk_document.h"
import "C"

import (
	"io"
	"unsafe"
)

type Document struct {
	handle *C.sk_document_t

	gws *goWStream
}

func NewDocument(w io.Writer) *Document {
	gws := newGoWStream(w)
	return &Document{
		handle: C.sk_document_create_pdf_from_stream((*C.sk_wstream_t)(unsafe.Pointer(gws.handle))),
		gws:    gws,
	}
}

func (d *Document) BeginPage(width, height float32, content *Rect) *Canvas {
	if width == 0 || height == 0 {
		panic("skia: can't begin a page with zero width or height")
	}

	return &Canvas{
		handle: C.sk_document_begin_page(d.handle, C.float(width), C.float(height), content.cptr()),
	}
}

func (d *Document) EndPage() {
	C.sk_document_end_page(d.handle)
}

func (d *Document) Close() {
	C.sk_document_close(d.handle)
}

func (d *Document) Dispose() {
	C.sk_document_unref(d.handle)
	d.handle = nil
	d.gws = nil
	d = nil
}
