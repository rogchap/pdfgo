package skia

//#include "sk_canvas.h"
import "C"

type Canvas struct {
	handle *C.sk_canvas_t
}

func (c *Canvas) DrawRect(rect *Rect, paint *Paint) {
	C.sk_canvas_draw_rect(c.handle, rect.cptr(), paint.handle)
}

func (c *Canvas) DrawText(text *TextBlob, x, y float32, paint *Paint) {
	C.sk_canvas_draw_text_blob(c.handle, text.handle, C.float(x), C.float(y), paint.handle)
}
