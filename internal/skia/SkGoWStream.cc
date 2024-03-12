#include "SkGoWStream.h"
#include "_cgo_export.h"

bool SkGoWStream::write(const void* buffer, size_t size) {
    // TODO: fBytesWritten should come from the return of the Go function
    fBytesWritten += size;

    goWStreamWrite(fGoWStream, (void*)buffer, size);

    // TODO: return false if the Go code errors
    return true;
}
