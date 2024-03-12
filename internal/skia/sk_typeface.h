#include "sk_types.h"

CPP_BEGIN_GUARD

// font manager
sk_fontmgr_t* sk_fontmgr_create_mac_default(void);
sk_typeface_t* sk_fontmgr_match_family_style(sk_fontmgr_t* fontmgr, const char* familyName, sk_fontstyle_t* style);

// font style
sk_fontstyle_t* sk_fontstyle_new(int weight, int width, sk_font_style_slant_t slant);

CPP_END_GUARD



