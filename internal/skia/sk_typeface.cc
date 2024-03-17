#include "include/core/SkFontMgr.h"
#include "include/core/SkFontStyle.h"
#include "include/core/SkTypeface.h"

#include "include/ports/SkFontMgr_mac_ct.h"

#include "sk_typeface.h"

#include "sk_types_priv.h"

// font manager
sk_fontmgr_t* sk_fontmgr_create_mac_default(void) {
    return ToFontMgr(SkFontMgr_New_CoreText(NULL).release());
}

sk_typeface_t* sk_fontmgr_match_family_style(sk_fontmgr_t* fontmgr, const char* familyName, sk_fontstyle_t* style) {
    return ToTypeface(AsFontMgr(fontmgr)->matchFamilyStyle(familyName, *AsFontStyle(style)).release());
}

sk_typeface_t* sk_fontmgr_create_from_file(sk_fontmgr_t* fontmgr, const char* path, int index) {
    return ToTypeface(AsFontMgr(fontmgr)->makeFromFile(path, index).release());
}

// font style
sk_fontstyle_t* sk_fontstyle_new(int weight, int width, sk_font_style_slant_t slant) {
    return ToFontStyle(new SkFontStyle(weight, width,(SkFontStyle::Slant)slant));
}

