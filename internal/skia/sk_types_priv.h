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

DEF_STRUCT_MAP(SkRect, sk_rect_t, Rect)

#endif
