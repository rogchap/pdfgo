package pdf

import (
	"rogchap.com/skia"
)

type Image struct {
	element

	raw *skia.Image
}

func (i *Image) messure(available size) sizePlan {
	w, h := i.raw.Width(), i.raw.Height()

	return sizePlan{
		size: size{
			width:  float32(w),
			height: float32(h),
		},
	}
}

func (i *Image) draw(available size) {
	if i.raw == nil {
		return
	}

	i.skdoc.canvas.DrawImage(i.raw, 0, 0)
}

func newImageFromFile(path string) *Image {
	raw := skia.NewImageFromFile(path)
	if raw == nil {
		return nil
	}
	return &Image{
		raw: raw,
	}
}
