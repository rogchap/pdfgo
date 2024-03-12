#include "sk_types.h"

CPP_BEGIN_GUARD

sk_font_t* sk_font_new(void);

void sk_font_set_typeface(sk_font_t* font, sk_typeface_t* value);
void sk_font_set_size(sk_font_t* font, float value);
void sk_font_set_scale_x(sk_font_t* font, float value);

int sk_font_text_to_glyphs(const sk_font_t* font, const void* text, size_t byteLength, sk_text_encoding_t encoding, uint16_t glyphs[], int maxGlyphCount);
void sk_font_get_pos(const sk_font_t* font, const uint16_t glyphs[], int count, sk_point_t pos[], sk_point_t* origin);


CPP_END_GUARD


