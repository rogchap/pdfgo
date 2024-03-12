package skia

//#include "sk_paint.h"
import "C"

type Paint struct {
	handle *C.sk_paint_t
}

func NewPaint(a, r, g, b uint8) *Paint {
	p := C.sk_paint_new()
	C.sk_paint_set_color(p, colorSetARGB(a, r, g, b))
	return &Paint{p}
}

func colorSetARGB(a, r, g, b uint8) C.sk_color_t {
	return C.sk_color_t((uint32(a) << 24) | (uint32(r) << 16) | (uint32(g) << 8) | uint32(b))
}
