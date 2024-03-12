#ifndef SkGoWStream_h
#define SkGoWStream_h

#include "include/core/SkStream.h"

class SkGoWStream : public SkWStream {
public:
    SkGoWStream(uintptr_t gws) : fBytesWritten(0), fGoWStream(gws) {}

    bool write(const void* buffer, size_t size) override;
    size_t bytesWritten() const override { return fBytesWritten; };

private:
    typedef SkWStream INHERITED;
    uintptr_t fGoWStream;
    size_t fBytesWritten;
};

#endif
