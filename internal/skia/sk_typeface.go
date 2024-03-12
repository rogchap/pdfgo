package skia

//#include "sk_typeface.h"
import "C"

import (
	"syscall"
	"unsafe"
)

type Typeface struct {
	handle *C.sk_typeface_t
}

// font manager

type FontMgr struct {
	handle *C.sk_fontmgr_t
}

// TODO: Move function to arch spacific file
func NewSystemFontMgr() *FontMgr {
	return &FontMgr{
		handle: C.sk_fontmgr_create_mac_default(),
	}
}

func (f *FontMgr) MatchFamily(familyName string) *Typeface {
	var sPtr *C.char = nil
	if familyName != "" {
		ptr, _ := syscall.BytePtrFromString(familyName)
		sPtr = (*C.char)(unsafe.Pointer(ptr))
	}
	return &Typeface{
		handle: C.sk_fontmgr_match_family_style(f.handle, sPtr, FontStyleNormal.handle),
	}
}

// font style

type FontStyle struct {
	handle *C.sk_fontstyle_t
}

func NewFontStyle(weight, width int, slant FontStyleSlant) *FontStyle {
	return &FontStyle{
		handle: C.sk_fontstyle_new(C.int(weight), C.int(width), slant),
	}
}

var FontStyleNormal = NewFontStyle(400, 5, FontStyleSlantUright)
