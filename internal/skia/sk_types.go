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
