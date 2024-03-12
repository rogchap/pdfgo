#include "include/core/SkCanvas.h"

#include "sk_canvas.h"

#include "sk_types_priv.h"

void sk_canvas_draw_rect(sk_canvas_t* ccanvas, const sk_rect_t* crect, const sk_paint_t* cpaint) {
    AsCanvas(ccanvas)->drawRect(*AsRect(crect), *AsPaint(cpaint));
}

