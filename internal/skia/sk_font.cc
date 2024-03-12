#include "include/core/SkFont.h"

#include "sk_font.h"

#include "sk_types_priv.h"

sk_font_t* sk_font_new(void) {
    return ToFont(new SkFont());
}

void sk_font_set_typeface(sk_font_t* font, sk_typeface_t* value) {
    AsFont(font)->setTypeface(sk_ref_sp(AsTypeface(value)));
}

void sk_font_set_size(sk_font_t* font, float value) {
    AsFont(font)->setSize(value);
}

void sk_font_set_scale_x(sk_font_t* font, float value) {
    AsFont(font)->setScaleX(value);
}

int sk_font_text_to_glyphs(const sk_font_t* font, const void* text, size_t byteLength, sk_text_encoding_t encoding, uint16_t glyphs[], int maxGlyphCount) {
    return AsFont(font)->textToGlyphs(text, byteLength, (SkTextEncoding)encoding, glyphs, maxGlyphCount);
}

void sk_font_get_pos(const sk_font_t* font, const uint16_t glyphs[], int count, sk_point_t pos[], sk_point_t* origin) {
    AsFont(font)->getPos(glyphs, count, AsPoint(pos), *AsPoint(origin));
}
