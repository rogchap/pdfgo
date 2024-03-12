#include "sk_types.h"

CPP_BEGIN_GUARD

void sk_document_unref(sk_document_t* document);

sk_document_t* sk_document_create_pdf_from_stream(sk_wstream_t* stream);

sk_canvas_t* sk_document_begin_page(sk_document_t* document, float width, float height, const sk_rect_t* content);
void sk_document_end_page(sk_document_t* document);
void sk_document_close(sk_document_t* document);

CPP_END_GUARD

