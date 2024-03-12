
#include "include/core/SkPixmap.h"

#include "sk_types_priv.h"
#include "sk_pixmap.h"
#include "SkGoWStream.h"

sk_pixmap_t* sk_pixmap_new(void) {
    return ToPixmap(new SkPixmap());
}

void sk_pixmap_reset(sk_pixmap_t* cpixmap) {
    AsPixmap(cpixmap)->reset();
}


