#include "sk_types_priv.h"
#include "sk_stream.h"

sk_gowstream_t* sk_gowstream_t_new(uintptr_t gws) {
    return ToGoWStream(new SkGoWStream(gws));
}

sk_wstream_filestream_t* sk_filewstream_new(const char* path) {
    return ToFileWStream(new SkFILEWStream(path));
}

