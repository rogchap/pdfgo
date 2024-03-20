package pdf

import (
	"github.com/go-text/typesetting/font"
	"github.com/go-text/typesetting/shaping"
	"golang.org/x/image/math/fixed"
	"rogchap.com/skia"
)

type textRun struct {
	glyphs        []*shapedGlyph
	face          font.Face
	size          float32
	ascent        float32
	descent       float32
	lineThickness float32
}

type shapedGlyph struct {
	glyphID   uint16
	position  skia.Point
	width     float32
	parentRun *textRun
}

// TODO: Do we need to consider concurrency?
var shaper = shaping.HarfbuzzShaper{}

func shape(text string, style TextStyle) []*textRun {
	var fontSize float32 = 16 // default
	if style.FontSize != 0 {
		fontSize = style.FontSize
	}

	runes := []rune(text)

	in := shaping.Input{
		Text:      runes,
		RunStart:  0,
		RunEnd:    len(runes),
		Direction: style.Direction,
		Size:      float32ToFixed266(fontSize),
	}

	ins := shaping.SplitByFace(in, defaultFontMgr.hbFontMgr)

	var x, y float32
	var runs []*textRun

	for _, in := range ins {
		out := shaper.Shape(in)

		run := textRun{
			face:          out.Face,
			size:          fixed266ToFloat32(out.Size),
			ascent:        fixed266ToFloat32(out.LineBounds.Ascent),
			descent:       fixed266ToFloat32(out.LineBounds.Descent),
			lineThickness: fixed266ToFloat32(out.LineBounds.LineThickness()),
		}

		for _, g := range out.Glyphs {

			run.glyphs = append(run.glyphs, &shapedGlyph{
				glyphID: uint16(g.GlyphID),
				position: skia.NewPoint(
					x+fixed266ToFloat32(g.XOffset),
					y-fixed266ToFloat32(g.YOffset),
				),
				width:     fixed266ToFloat32(g.XAdvance),
				parentRun: &run,
			})

			x += fixed266ToFloat32(g.XAdvance)
			y += fixed266ToFloat32(g.YAdvance)
		}

		runs = append(runs, &run)
	}

	return runs
}

func fixed266ToFloat32(i fixed.Int26_6) float32 {
	return float32(float64(i) / (1 << 6))
}

func float32ToFixed266(f float32) fixed.Int26_6 {
	return fixed.Int26_6(float64(f) * (1 << 6))
}
