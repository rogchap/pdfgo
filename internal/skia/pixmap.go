package skia

//#include "sk_pixmap.h"
import "C"

type Pixmap struct {
	handle *C.sk_pixmap_t
}

func NewPixmap() *Pixmap {
	return &Pixmap{
		handle: C.sk_pixmap_new(),
	}
}

func (p *Pixmap) Reset() {
	C.sk_pixmap_reset(p.handle)
}
