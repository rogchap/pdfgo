#include "include/core/SkPaint.h"

#include "sk_paint.h"

#include "sk_types_priv.h"

sk_paint_t* sk_paint_new(void) {
    return ToPaint(new SkPaint());
}

void sk_paint_set_color(sk_paint_t* cpaint, sk_color_t c) {
    AsPaint(cpaint)->setColor(c);
}
