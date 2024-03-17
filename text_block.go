package pdf

import (
	"fmt"

	"rogchap.com/pdf/internal/skia"
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
	fmt.Printf("startIdx: %v, endIdx: %v\n", startIdx, endIdx)

	glyphsToDraw := ts.glyphs[startIdx : endIdx+1]
	if len(glyphsToDraw) == 0 {
		return
	}

	fmt.Printf("text: %v all: %v todraw: %v\n", len([]rune(ts.text)), len(ts.glyphs), len(glyphsToDraw))

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

		fmt.Println("run count:", count)

		runbuf := builder.AllocPosRun(skfont, count)
		tbg := runbuf.Glyphs(count)
		tbp := runbuf.Pos(count)

		for idx, g := range glyphs {
			tbg[idx] = g.glyphID
			tbp[idx] = g.position
		}
	}

	// TODO paint should be from TextStyle
	p := skia.NewPaint(0xFF, 0, 0, 0)

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

	// TODO: is this the line bounds gap?
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

func (tb *TextBlock) messure(available size) sizePlan {
	if len(tb.items) == 0 {
		return sizePlan{}
	}

	return sizePlan{}
	// panic("")
}

func (tb *TextBlock) draw(sp sizePlan) {
	lines := tb.splitIntoLines(sp.size.width, sp.size.height)

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

	var queue []*TextSpan
	for _, item := range tb.items {
		queue = append(queue, item)
	}
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

	// var currentHeight float32

	for len(queue) > 0 {
		line := nextLine()

		if len(line) == 0 {
			return lines
		}

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
		}

		lines = append(lines, textLine{
			elements: line,

			textHeight: textHeight,
			lineHeight: lineHeight,

			ascent:  ascent,
			descent: descent,
		})

	}

	return lines
}
