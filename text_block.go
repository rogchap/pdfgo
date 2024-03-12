package pdf

import "rogchap.com/pdf/internal/skia"

type TextSpan struct {
	text string
}

func (ts *TextSpan) draw(skdoc *skiaDoc) {
	// TODO draw text for real

	tb := skia.NewTextBlob(ts.text)
	p := skia.NewPaint(0xFF, 0, 0, 0)
	skdoc.canvas.DrawText(tb, 100, 100, p)
}

type TextBlock struct {
	element

	items []*TextSpan
}

func (tb *TextBlock) Span(text string) {
	tb.items = append(tb.items, &TextSpan{
		text: text,
	})
}

func (tb *TextBlock) messure(available size) sizePlan {
	// TODO: messure
	return sizePlan{}
}

func (tb *TextBlock) draw(sp sizePlan) {
	for _, span := range tb.items {
		span.draw(tb.skdoc)
	}
}
