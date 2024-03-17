package skia

//#include "sk_types.h"
import "C"

import "unsafe"

type Rect struct {
	Left   float32
	Top    float32
	Right  float32
	Bottom float32
}

func NewRect(x, y, w, h float32) Rect {
	return Rect{
		Left:   x,
		Top:    y,
		Right:  w,
		Bottom: h,
	}
}

func (r *Rect) cptr() *C.sk_rect_t {
	if r == nil {
		return nil
	}

	return (*C.sk_rect_t)(unsafe.Pointer(r))
}

type Point = C.sk_point_t

func NewPoint(x, y float32) Point {
	return C.sk_point_t{
		x: C.float(x),
		y: C.float(y),
	}
}

func (p Point) X() float32 {
	return float32(p.x)
}

func (p Point) Y() float32 {
	return float32(p.y)
}

type TextBlobBuilderRunBuffer = C.sk_textblob_builder_runbuffer_t

type FontStyleSlant = C.sk_font_style_slant_t

const (
	FontStyleSlantUright FontStyleSlant = iota
	FontStyleSlantItalic
	FontStyleSlantOblique
)
