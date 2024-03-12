#ifndef sk_types_DEFINED
#define sk_types_DEFINED

#include <stdint.h>

#ifdef __cplusplus
    #define CPP_BEGIN_GUARD     extern "C" {
    #define CPP_END_GUARD       }
#else
    #define CPP_BEGIN_GUARD
    #define CPP_END_GUARD
#endif

CPP_BEGIN_GUARD

typedef uint32_t sk_color_t;

typedef struct sk_pixmap_t sk_pixmap_t;
typedef struct sk_wstream_filestream_t sk_wstream_filestream_t;
typedef struct sk_gowstream_t sk_gowstream_t;
typedef struct sk_document_t sk_document_t;
typedef struct sk_wstream_t sk_wstream_t;
typedef struct sk_canvas_t sk_canvas_t;
typedef struct sk_paint_t sk_paint_t;

typedef struct {
    float   left;
    float   top;
    float   right;
    float   bottom;
} sk_rect_t;


CPP_END_GUARD

#endif
