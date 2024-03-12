#include "include/core/SkTextBlob.h"

#include "sk_textblob.h"

#include "sk_types_priv.h"

sk_textblob_builder_t* sk_textblob_builder_new(void) {
    return ToTextBlobBuilder(new SkTextBlobBuilder());
}

sk_textblob_t* sk_textblob_builder_make(sk_textblob_builder_t* builder) {
    return ToTextBlob(AsTextBlobBuilder(builder)->make().release());
}

void sk_textblob_builder_alloc_run_pos(sk_textblob_builder_t* builder, const sk_font_t* font, int count, const sk_rect_t* bounds, sk_textblob_builder_runbuffer_t* runbuffer) {
    *runbuffer = ToTextBlobBuilderRunBuffer(AsTextBlobBuilder(builder)->allocRunPos(AsFont(*font), count, AsRect(bounds)));
}
