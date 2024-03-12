#include <stdint.h>
#include "sk_types.h"

CPP_BEGIN_GUARD

sk_gowstream_t* sk_gowstream_t_new(uintptr_t gws);
sk_wstream_filestream_t* sk_filewstream_new(const char* path);

CPP_END_GUARD

