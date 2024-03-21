package pdf

import (
	"image/color"

	"rogchap.com/skia"
)

type TextSpan struct {
	text  string
	style TextStyle

	glyphs []*shapedGlyph
}

func (ts *TextSpan) FontSize(s float32) {
	ts.style.FontSize = s
}

type messureResult struct {
	width      float32
	height     float32
	lineHeight float32
	ascent     float32
	descent    float32
	startIdx   int
	endIdx     int
	nextIdx    int
	totalIdx   int
}

func (ts *TextSpan) draw(skcanvas *skia.Canvas, startIdx, endIdx int) {
	glyphsToDraw := ts.glyphs[startIdx : endIdx+1]
	if len(glyphsToDraw) == 0 {
		return
	}

	runs := make(map[*textRun][]*shapedGlyph)

	for _, glyph := range glyphsToDraw {
		runs[glyph.parentRun] = append(runs[glyph.parentRun], glyph)
	}

	builder := skia.NewTextBlobBuilder()

	for run, glyphs := range runs {
		skfont := skia.NewFont()
		skfont.SetTypeface(defaultFontMgr.skTypeface(run.face))
		skfont.SetSize(run.size)

		count := len(glyphs)
		runbuf := builder.AllocPosRun(skfont, count)
		tbg := runbuf.Glyphs(count)
		tbp := runbuf.Pos(count)

		for idx, g := range glyphs {
			tbg[idx] = g.glyphID
			tbp[idx] = g.position
		}
	}

	// TODO paint should be from TextStyle
	p := skia.NewPaint(color.RGBA{0, 0, 0, 0xff})

	xOffset := glyphsToDraw[0].position.X()
	skcanvas.DrawText(builder.Make(), -xOffset, 0, p)
}

func (ts *TextSpan) messure(startIdx int, availableWidth float32) *messureResult {
	ts.glyphs = nil

	runs := shape(ts.text, ts.style)

	for _, run := range runs {
		ts.glyphs = append(ts.glyphs, run.glyphs...)
	}

	if len(ts.glyphs) == 0 {
		return nil
	}

	endIdx := maxEndIdx(ts.glyphs, startIdx, availableWidth)

	if endIdx < startIdx {
		return nil
	}

	nextIdx := endIdx

	if endIdx != len(ts.glyphs)-1 {
		// need to wrap at last space
		// space glyphID = 0x3
		var spaceGlyph uint16 = 0x3
		if ts.glyphs[endIdx].glyphID != spaceGlyph && ts.glyphs[endIdx+1].glyphID == spaceGlyph {
			nextIdx = endIdx + 2
		} else {
			// find the last space
			lastSpaceIdx := endIdx
			for lastSpaceIdx >= startIdx {
				if ts.glyphs[lastSpaceIdx].glyphID == spaceGlyph {
					break
				}
				lastSpaceIdx--
			}

			if lastSpaceIdx > 1 && lastSpaceIdx >= startIdx {
				endIdx = lastSpaceIdx - 1
				nextIdx = lastSpaceIdx + 1
			} else {
				nextIdx = endIdx + 1
			}
		}
	}

	start := ts.glyphs[startIdx]
	end := ts.glyphs[endIdx]

	width := end.position.X() - start.position.X() + end.width

	var height float32
	var ascent, descent float32
	for _, glyph := range ts.glyphs[startIdx:endIdx] {
		if glyph.parentRun.lineThickness > height {
			height = glyph.parentRun.lineThickness
		}
		if glyph.parentRun.ascent > ascent {
			ascent = glyph.parentRun.ascent
		}
		if glyph.parentRun.descent < descent {
			descent = glyph.parentRun.descent
		}
	}

	lineHeight := ts.style.LineHeight
	if lineHeight < 1 {
		lineHeight = 1
	}

	return &messureResult{
		width:      width,
		height:     height,
		lineHeight: lineHeight,
		ascent:     ascent,
		descent:    descent,
		startIdx:   startIdx,
		endIdx:     endIdx,
		nextIdx:    nextIdx,
		totalIdx:   len(ts.glyphs) - 1,
	}
}

func maxEndIdx(glyphs []*shapedGlyph, startIdx int, availableWidth float32) int {
	maxWidth := availableWidth + glyphs[startIdx].position.X()

	idx := startIdx
	for idx < len(glyphs) {
		g := glyphs[idx]
		if g.position.X()+g.width > maxWidth {
			break
		}
		idx++
	}
	return idx - 1
}

type TextBlock struct {
	element

	items          []*TextSpan
	renderQueue    []*TextSpan
	currentItemIdx int
}

func (tb *TextBlock) Span(text string) *TextSpan {
	ts := &TextSpan{
		text: text,
	}
	ts.style.LineHeight = 1
	tb.items = append(tb.items, ts)
	return ts
}

func (tb *TextBlock) reset() {
	tb.currentItemIdx = 0
	tb.renderQueue = nil

	for _, item := range tb.items {
		tb.renderQueue = append(tb.renderQueue, item)
	}
}

func (tb *TextBlock) messure(available size) sizePlan {
	if len(tb.renderQueue) == 0 {
		return sizePlan{}
	}

	// TODO: We should cache mesurements as we do this twice, trice ... maybe more; once for messure and again for draw
	// In a lot of cases the values in both will be the same (unless we have to go over multiple pages
	lines := tb.splitIntoLines(available.width, available.height)

	if len(lines) == 0 {
		return sizePlan{
			pType: wrap,
		}
	}

	var width, height float32
	for _, line := range lines {
		if line.width > width {
			width = line.width
		}
		height += line.lineHeight
	}

	if width > available.width || height > available.height {
		return sizePlan{pType: wrap}
	}

	var willRenderCount int
	for _, line := range lines {
		for _, el := range line.elements {
			if el.messurement.endIdx == el.messurement.totalIdx {
				willRenderCount++
			}
		}
	}

	if willRenderCount != len(tb.renderQueue) {
		return sizePlan{
			pType: partial,
			size: size{
				width:  width,
				height: height,
			},
		}
	}

	return sizePlan{
		size: size{
			width:  width,
			height: height,
		},
	}
}

func (tb *TextBlock) draw(available size) {
	lines := tb.splitIntoLines(available.width, available.height)

	if len(lines) == 0 {
		return
	}

	var topOffset float32
	for _, line := range lines {
		var leftOffset float32 = 0 // left aligned for now

		for _, item := range line.elements {

			tb.skdoc.canvas.Translate(leftOffset, topOffset+line.ascent)
			item.item.draw(tb.skdoc.canvas, item.messurement.startIdx, item.messurement.endIdx)
			tb.skdoc.canvas.Translate(-leftOffset, -(topOffset + line.ascent))

			leftOffset += item.messurement.width
		}

		topOffset += line.lineHeight
	}

	for _, line := range lines {
		for _, el := range line.elements {
			if el.messurement.endIdx == el.messurement.totalIdx {
				tb.renderQueue = tb.renderQueue[1:]
			}
		}
	}

	lastLine := lines[len(lines)-1]
	lastEl := lastLine.elements[len(lastLine.elements)-1]
	tb.currentItemIdx = lastEl.messurement.nextIdx

	if lastEl.messurement.endIdx == lastEl.messurement.totalIdx {
		tb.currentItemIdx = 0
	}

	if len(tb.renderQueue) == 0 {
		tb.reset()
	}
}

type lineElement struct {
	item        *TextSpan
	messurement *messureResult
}

type textLine struct {
	elements []lineElement

	textHeight float32
	lineHeight float32
	ascent     float32
	descent    float32
	width      float32
}

func (tb *TextBlock) splitIntoLines(availableWidth, availableHeight float32) []textLine {
	var lines []textLine

	queue := make([]*TextSpan, len(tb.renderQueue))
	copy(queue, tb.renderQueue)

	currentItemIdx := tb.currentItemIdx

	nextLine := func() []lineElement {
		var currentWidth float32
		var lineElements []lineElement

		for {
			if len(queue) == 0 {
				break
			}

			currentSpan := queue[0]

			if len(lineElements) != 0 {
				currentSpan.text = " " + currentSpan.text
			}

			result := currentSpan.messure(currentItemIdx, availableWidth-currentWidth)

			lineElements = append(lineElements, lineElement{
				item:        currentSpan,
				messurement: result,
			})
			currentWidth += result.width
			currentItemIdx = result.nextIdx

			if result.endIdx != result.nextIdx {
				break
			}

			currentItemIdx = 0
			queue = queue[1:]
		}
		return lineElements
	}

	var currentHeight float32

	for len(queue) > 0 {
		line := nextLine()

		if len(line) == 0 {
			return lines
		}

		var width float32
		var textHeight, lineHeight float32
		var ascent, descent float32
		for _, item := range line {
			if item.messurement.height > textHeight {
				textHeight = item.messurement.height
			}
			if item.messurement.lineHeight*item.messurement.height > lineHeight {
				lineHeight = item.messurement.lineHeight * item.messurement.height
			}
			if item.messurement.ascent > ascent {
				ascent = item.messurement.ascent
			}
			if item.messurement.descent < descent {
				descent = item.messurement.descent
			}
			width += item.messurement.width
		}

		if currentHeight+lineHeight > availableHeight {
			break
		}

		currentHeight += lineHeight

		lines = append(lines, textLine{
			elements: line,

			width:      width,
			textHeight: textHeight,
			lineHeight: lineHeight,

			ascent:  ascent,
			descent: descent,
		})

	}

	return lines
}
