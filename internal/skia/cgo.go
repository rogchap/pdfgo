package skia

// #cgo LDFLAGS: -lskia
// #cgo CXXFLAGS: -I${SRCDIR}/../../external/skia -std=c++17
// #cgo darwin LDFLAGS: -framework CoreFoundation -framework CoreGraphics -framework CoreText -framework CoreServices
// #cgo darwin,arm64 LDFLAGS: -L${SRCDIR}/../../external/skia/out/mac-apple
import "C"
