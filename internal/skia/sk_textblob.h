#include "sk_types.h"

CPP_BEGIN_GUARD

sk_textblob_builder_t* sk_textblob_builder_new(void);
sk_textblob_t* sk_textblob_builder_make(sk_textblob_builder_t* builder);
void sk_textblob_builder_alloc_run_pos(sk_textblob_builder_t* builder, const sk_font_t* font, int count, const sk_rect_t* bounds, sk_textblob_builder_runbuffer_t* runbuffer);

CPP_END_GUARD


