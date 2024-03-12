package pdf

import (
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

func (b *background) draw(sp sizePlan) {
	// TODO: Change to real paint
	p := skia.NewPaint(0xFF, 0xF3, 0xE2, 0xD3)
	r := skia.NewRect(0, 0, sp.size.width, sp.size.height)
	b.skdoc.canvas.DrawRect(&r, p)

	if b.child != nil {
		b.child.draw(sp)
	}
}
