package skia

//#include "sk_stream.h"
import "C"

import (
	"io"
	"syscall"
	"unsafe"
)

type goWStream struct {
	handle *C.sk_gowstream_t
	writer io.Writer
}

func newGoWStream(w io.Writer) *goWStream {
	ws := &goWStream{
		writer: w,
	}
	ptr := uintptr(unsafe.Pointer(ws))
	ws.handle = C.sk_gowstream_t_new(C.uintptr_t(ptr))

	return ws
}

//export goWStreamWrite
func goWStreamWrite(gws uintptr, buffer unsafe.Pointer, size C.int) {
	ws := (*goWStream)(unsafe.Pointer(gws))
	// TODO return bytes written and if error return false back to C
	ws.writer.Write(C.GoBytes(buffer, size))
}

type FileStream struct {
	handle *C.sk_wstream_filestream_t
}

func NewFileStream(path string) *FileStream {
	ptr, _ := syscall.BytePtrFromString(path)
	return &FileStream{
		handle: C.sk_filewstream_new((*C.char)(unsafe.Pointer(ptr))),
	}
}
