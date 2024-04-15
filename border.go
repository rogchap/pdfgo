package pdf

import "rogchap.com/skia"

type BorderStyle struct {
	Color string

	Left   float32
	Top    float32
	Right  float32
	Bottom float32

	RadiusTopLeft     float32
	RadiusTopRight    float32
	RadiusBottomRight float32
	RadiusBottomLeft  float32
}

type border struct {
	container

	style *BorderStyle
}

func defaultBorderStyle() *BorderStyle {
	return &BorderStyle{
		Color:  "#000",
		Left:   1,
		Top:    1,
		Right:  1,
		Bottom: 1,
	}
}

func (b *border) draw(available size) {
	b.container.draw(available)

	c := b.canvas()
	w, h := available.width, available.height

	style := b.style
	if style == nil {
		style = defaultBorderStyle()
	}

	paint := skia.NewPaint(parseHexColor(style.Color))
	paint.SetStyle(skia.PaintStyleStroke)
	paint.SetStrokeCap(skia.StrokeCapSquare)

	if style.Top > 0 {
		paint.SetStrokeWidth(style.Top)
		path := skia.NewPath()
		path.MoveTo(0+style.RadiusTopLeft, 0)
		path.LineTo(w-style.RadiusTopRight, 0)
		if style.RadiusTopRight > 0 {
			path.ArcTo(skia.NewRect(w-style.RadiusTopRight*2, 0, w, style.RadiusTopRight*2), -90, 90, false)
		}
		c.DrawPath(path, paint)
	}

	if style.Right > 0 {
		paint.SetStrokeWidth(style.Right)
		path := skia.NewPath()
		path.MoveTo(w, style.RadiusTopRight)
		path.LineTo(w, h-style.RadiusBottomRight)
		if style.RadiusBottomRight > 0 {
			path.ArcTo(skia.NewRect(w-style.RadiusBottomRight*2, h-style.RadiusBottomRight*2, w, h), 0, 90, false)
		}
		c.DrawPath(path, paint)
	}

	if style.Bottom > 0 {
		paint.SetStrokeWidth(style.Bottom)
		path := skia.NewPath()
		path.MoveTo(w-style.RadiusBottomRight, h)
		path.LineTo(style.RadiusBottomLeft, h)
		if style.RadiusBottomLeft > 0 {
			path.ArcTo(skia.NewRect(0, h-style.RadiusBottomLeft*2, style.RadiusBottomLeft*2, h), 90, 90, false)
		}
		c.DrawPath(path, paint)
	}

	if style.Left > 0 {
		paint.SetStrokeWidth(style.Left)
		path := skia.NewPath()
		path.MoveTo(0, h-style.RadiusBottomLeft)
		path.LineTo(0, style.RadiusTopLeft)
		if style.RadiusTopLeft > 0 {
			path.ArcTo(skia.NewRect(0, 0, style.RadiusTopLeft*2, style.RadiusTopLeft*2), 180, 90, false)
		}
		c.DrawPath(path, paint)
	}
}
