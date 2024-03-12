package pdf

import (
	"fmt"

	"rogchap.com/pdf/internal/skia"
)

type background struct {
	container

	color string // TODO ARGB
}

func (b *background) messure(available size) sizePlan {
	if b.child != nil {
		return b.child.messure(available)
	}
	return sizePlan{}
}

func (b *background) draw(available sizePlan) {
	fmt.Printf("%#v\n", b.skdoc.canvas)

	p := skia.NewPaint(0xFF, 0xDE, 0x22, 0xCD)
	r := skia.NewRect(0, 0, available.size.width, available.size.height)
	b.skdoc.canvas.DrawRect(&r, p)
}
