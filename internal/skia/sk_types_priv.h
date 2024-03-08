#include "sk_types.h"

#include "include/core/SkTypes.h"

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

DEF_CLASS_MAP(SkPixmap, sk_pixmap_t, Pixmap)
