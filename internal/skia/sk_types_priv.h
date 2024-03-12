#ifndef sk_types_priv_DEFINED
#define sk_types_priv_DEFINED

#include "sk_types.h"

#include "include/core/SkTypes.h"

#include "SkGoWStream.h"

// Define a mapping between a C++ type and the C type.
//
// Usual Values:
//  - C++  |  SkType   |  SkSomeType
//  - C    |  sk_type  |  sk_some_type_t
//  - Map  |  Name     |  ToSomeType / AsSomeType
//
#define DEF_MAP_DECL(SkType, sk_type, Name, Declaration, Ns)        \
    Declaration;                                                    \
    static inline const Ns::SkType& As##Name(const sk_type& t) {    \
        return reinterpret_cast<const Ns::SkType&>(t);              \
    }                                                               \
    static inline const Ns::SkType* As##Name(const sk_type* t) {    \
        return reinterpret_cast<const Ns::SkType*>(t);              \
    }                                                               \
    static inline Ns::SkType& As##Name(sk_type& t) {                \
        return reinterpret_cast<Ns::SkType&>(t);                    \
    }                                                               \
    static inline Ns::SkType* As##Name(sk_type* t) {                \
        return reinterpret_cast<Ns::SkType*>(t);                    \
    }                                                               \
    static inline const sk_type& To##Name(const Ns::SkType& t) {    \
        return reinterpret_cast<const sk_type&>(t);                 \
    }                                                               \
    static inline const sk_type* To##Name(const Ns::SkType* t) {    \
        return reinterpret_cast<const sk_type*>(t);                 \
    }                                                               \
    static inline sk_type& To##Name(Ns::SkType& t) {                \
        return reinterpret_cast<sk_type&>(t);                       \
    }                                                               \
    static inline sk_type* To##Name(Ns::SkType* t) {                \
        return reinterpret_cast<sk_type*>(t);                       \
    }

#define DEF_CLASS_MAP(SkType, sk_type, Name)                   \
    DEF_MAP_DECL(SkType, sk_type, Name, class SkType, )

#define DEF_STRUCT_MAP(SkType, sk_type, Name)                  \
    DEF_MAP_DECL(SkType, sk_type, Name, struct SkType, )

DEF_CLASS_MAP(SkPixmap, sk_pixmap_t, Pixmap)
DEF_CLASS_MAP(SkGoWStream, sk_gowstream_t, GoWStream)
DEF_CLASS_MAP(SkFILEWStream, sk_wstream_filestream_t, FileWStream)
DEF_CLASS_MAP(SkDocument, sk_document_t, Document)
DEF_CLASS_MAP(SkWStream, sk_wstream_t, WStream)
DEF_CLASS_MAP(SkCanvas, sk_canvas_t, Canvas)
DEF_CLASS_MAP(SkPaint, sk_paint_t, Paint)
DEF_CLASS_MAP(SkTextBlob, sk_textblob_t, TextBlob)
DEF_CLASS_MAP(SkTextBlobBuilder, sk_textblob_builder_t, TextBlobBuilder)
DEF_CLASS_MAP(SkFont, sk_font_t, Font)
DEF_CLASS_MAP(SkTypeface, sk_typeface_t, Typeface)
DEF_CLASS_MAP(SkFontMgr, sk_fontmgr_t, FontMgr)
DEF_CLASS_MAP(SkFontStyle, sk_fontstyle_t, FontStyle)

DEF_STRUCT_MAP(SkRect, sk_rect_t, Rect)
DEF_STRUCT_MAP(SkPoint, sk_point_t, Point)

#include "include/core/SkTextBlob.h"
static inline SkTextBlobBuilder::RunBuffer AsTextBlobBuilderRunBuffer(const sk_textblob_builder_runbuffer_t* runbuffer) {
    return {
        (SkGlyphID*)runbuffer->glyphs,
        (SkScalar*)runbuffer->pos,
        (char*)runbuffer->utf8text,
        (uint32_t*)runbuffer->clusters,
    };
}

static inline sk_textblob_builder_runbuffer_t ToTextBlobBuilderRunBuffer(const SkTextBlobBuilder::RunBuffer runbuffer) {
    return {
        runbuffer.glyphs,
        runbuffer.pos,
        runbuffer.utf8text,
        runbuffer.clusters,
    };
}

#endif
