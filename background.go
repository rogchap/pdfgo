package pdf

import (
	"image/color"

	"rogchap.com/skia"
)

type background struct {
	container

	color string
}

func (b *background) messure(available size) sizePlan {
	if b.child != nil {
		return b.child.messure(available)
	}
	return sizePlan{}
}

func (b *background) draw(available size) {
	c := parseHexColor(b.color)

	p := skia.NewPaint(c)
	r := skia.NewRect(0, 0, available.width, available.height)
	b.skdoc.canvas.DrawRect(&r, p)

	if b.child != nil {
		b.child.draw(available)
	}
}

func parseHexColor(s string) color.RGBA {
	white := color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}

	if s == "" {
		return white
	}

	var c color.RGBA

	if s[0] == '#' {
		s = s[1:]
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		return 0
	}

	switch len(s) {
	case 8:
		c.A = hexToByte(s[0])<<4 + hexToByte(s[1])
		c.R = hexToByte(s[2])<<4 + hexToByte(s[3])
		c.G = hexToByte(s[4])<<4 + hexToByte(s[5])
		c.B = hexToByte(s[6])<<4 + hexToByte(s[7])
	case 6:
		c.A = 0xFF
		c.R = hexToByte(s[0])<<4 + hexToByte(s[1])
		c.G = hexToByte(s[2])<<4 + hexToByte(s[3])
		c.B = hexToByte(s[4])<<4 + hexToByte(s[5])
	case 3:
		c.A = 0xFF
		c.R = hexToByte(s[0]) * 17
		c.G = hexToByte(s[1]) * 17
		c.B = hexToByte(s[2]) * 17
	default:
		c = white
	}
	return c
}
