package skia

//#include "sk_font.h"
import "C"

import (
	"syscall"
	"unicode/utf8"
	"unsafe"
)

type Font struct {
	handle *C.sk_font_t
}

func NewFont() *Font {
	return &Font{
		handle: C.sk_font_new(),
	}
}

func (f *Font) SetTypeface(t *Typeface) {
	C.sk_font_set_typeface(f.handle, t.handle)
}

func (f *Font) SetSize(s float32) {
	C.sk_font_set_size(f.handle, C.float(s))
}

func (f *Font) SetScale(s float32) {
	C.sk_font_set_scale_x(f.handle, C.float(s))
}

func (f *Font) GlyphCount(text string) int {
	ptr, _ := syscall.BytePtrFromString(text)
	l := C.ulong(utf8.RuneCountInString(text))

	count := C.sk_font_text_to_glyphs(f.handle, unsafe.Pointer(ptr), l, C.UTF8_SK_TEXT_ENCODING, nil, 0)
	return int(count)
}

func (f *Font) Glyphs(text string, glyphCount int, buf *TextBlobBuilderRunBuffer) {
	ptr, _ := syscall.BytePtrFromString(text)
	l := C.ulong(utf8.RuneCountInString(text))

	C.sk_font_text_to_glyphs(f.handle, unsafe.Pointer(ptr), l, C.UTF8_SK_TEXT_ENCODING, (*C.ushort)(buf.glyphs), C.int(glyphCount))
}

func (f *Font) GlyphPositions(glyphCount int, buf *TextBlobBuilderRunBuffer) {
	// TODO: pass in origin
	C.sk_font_get_pos(f.handle, (*C.ushort)(buf.glyphs), C.int(glyphCount), (*C.sk_point_t)(buf.pos), &C.sk_point_t{})
}
