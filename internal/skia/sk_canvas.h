#include "sk_types.h"

CPP_BEGIN_GUARD

void sk_canvas_draw_rect(sk_canvas_t* ccanvas, const sk_rect_t* crect, const sk_paint_t* cpaint);
void sk_canvas_draw_text_blob (sk_canvas_t* ccanvas, sk_textblob_t* text, float x, float y, const sk_paint_t* cpaint);

void sk_canvas_translate(sk_canvas_t* ccanvas, float dx, float dy);

CPP_END_GUARD
