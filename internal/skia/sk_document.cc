#include "include/core/SkDocument.h"
#include "include/docs/SkPDFDocument.h"
// #include "include/core/SkRect.h"

#include "sk_document.h"

#include "sk_types_priv.h"

void sk_document_unref(sk_document_t* document) {
    SkSafeUnref(AsDocument(document));
}

sk_document_t* sk_document_create_pdf_from_stream(sk_wstream_t* stream) {
    return ToDocument(SkPDF::MakeDocument(AsWStream(stream)).release());
}

sk_canvas_t* sk_document_begin_page(sk_document_t* document, float width, float height, const sk_rect_t* content) {
    return ToCanvas(AsDocument(document)->beginPage(width, height, AsRect(content)));
}

void sk_document_end_page(sk_document_t* document) {
    AsDocument(document)->endPage();
}

void sk_document_close(sk_document_t* document) {
    AsDocument(document)->close();
}

