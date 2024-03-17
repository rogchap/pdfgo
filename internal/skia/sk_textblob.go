package skia

//#include "sk_textblob.h"
import "C"
import "unsafe"

type TextBlob struct {
	handle *C.sk_textblob_t
}

func NewTextBlob(text string) *TextBlob {
	fm := NewSystemFontMgr()
	tf := fm.MatchFamily("")

	font := NewFont()
	font.SetTypeface(tf)
	count := font.GlyphCount(text)

	builder := NewTextBlobBuilder()
	runbuf := builder.AllocPosRun(font, count)
	font.Glyphs(text, count, runbuf)
	font.GlyphPositions(count, runbuf)

	return builder.Make()
}

type TextBlobBuilder struct {
	handle *C.sk_textblob_builder_t
}

func NewTextBlobBuilder() *TextBlobBuilder {
	return &TextBlobBuilder{
		handle: C.sk_textblob_builder_new(),
	}
}

func (b *TextBlobBuilder) AllocPosRun(font *Font, count int) *TextBlobBuilderRunBuffer {
	rb := &C.sk_textblob_builder_runbuffer_t{}
	C.sk_textblob_builder_alloc_run_pos(b.handle, font.handle, C.int(count), nil, rb)
	return rb
}

func (b *TextBlobBuilder) Make() *TextBlob {
	return &TextBlob{
		handle: C.sk_textblob_builder_make(b.handle),
	}
}

func (buf *TextBlobBuilderRunBuffer) Glyphs(count int) []uint16 {
	return unsafe.Slice((*uint16)(buf.glyphs), count)
}

func (buf *TextBlobBuilderRunBuffer) Pos(count int) []Point {
	return unsafe.Slice((*Point)(buf.pos), count)
}
