package pdf

import (
	"math"
)

type Cell struct {
	container

	row      int
	col      int
	rendered bool
}

type Table struct {
	element

	cols  []*column
	cells []*Cell

	rowCells [][]*Cell
}

type column struct {
	relWidth   float32
	constWidth float32

	calcWidth float32
}

func (t *Table) RelCol(size float32) {
	if size < 1 {
		size = 1
	}
	col := column{
		relWidth: size,
	}
	t.cols = append(t.cols, &col)
}

func (t *Table) ConstCol(width float32) {
	col := column{
		constWidth: width,
	}
	t.cols = append(t.cols, &col)
}

func (t *Table) Cell() *Cell {
	c := &Cell{}
	t.cells = append(t.cells, c)

	return c
}

func (t *Table) organizeCells() {
	currentRow := 1
	currentCol := 1

	colCount := len(t.cols)

	for idx, cell := range t.cells {
		if cell.row != 0 && cell.col != 0 {
			currentRow, currentCol = cell.row, cell.col
			continue
		}

		// TODO: deal with col span and row span

		cell.row = currentRow
		cell.col = currentCol

		currentCol++

		if colCount <= 1 {
			currentRow++
			continue
		}

		if (idx+1)%colCount == 0 {
			currentRow++
			currentCol = 1
		}
	}
}

func (t *Table) reset() {
	for _, cell := range t.cells {
		cell.rendered = false
	}
}

func (t *Table) children() []drawable {
	return asDrawable(t.cells)
}

func (t *Table) messure(available size) sizePlan {
	if len(t.cells) == 0 {
		return fullRenderZero()
	}

	// TODO add cache
	t.calcColWidths(available.width)
	layouts := t.layout(available)

	var width, height float32
	for _, col := range t.cols {
		width += col.calcWidth
	}
	for _, layout := range layouts {
		h := layout.yOffset + layout.size.height
		if h > height {
			height = h
		}
	}

	// TODO: deal with partial page rendering
	return fullRender(size{width, height})
}

func (t *Table) draw(available size) {
	if len(t.cells) == 0 {
		return
	}

	t.calcColWidths(available.width)
	layouts := t.layout(available)

	for _, layout := range layouts {
		if layout.pType == full {
			layout.cell.rendered = true
		}
		if layout.pType == wrap {
			continue
		}

		canvas := t.canvas()
		canvas.Translate(layout.xOffset, layout.yOffset)
		layout.cell.draw(layout.size)
		canvas.Translate(-layout.xOffset, -layout.yOffset)
	}
}

func (t *Table) calcColWidths(availableWidth float32) {
	var cWidth, rWidth float32
	for _, col := range t.cols {
		cWidth += col.constWidth
		rWidth += col.relWidth
	}

	relPercent := (availableWidth - cWidth) / rWidth

	for _, col := range t.cols {
		col.calcWidth = col.constWidth + col.relWidth*relPercent
	}
}

type tableLayout struct {
	pType   sizePlanType
	cell    *Cell
	size    size
	xOffset float32
	yOffset float32
}

func (t *Table) layout(available size) []tableLayout {
	var layouts []tableLayout

	xOffsets := make([]float32, len(t.cols)+1)
	xOffsets[0] = 0

	yOffsets := make(map[int]float64)

	for i := 1; i <= len(t.cols); i++ {
		xOffsets[i] = t.cols[i-1].calcWidth + xOffsets[i-1]
	}

	currentRow := 1

	for _, cell := range t.cells {
		if cell.row > currentRow {
			yOffsets[currentRow] = math.Max(yOffsets[currentRow], yOffsets[currentRow-1])

			if yOffsets[currentRow-1] > float64(available.height) {
				break
			}

			for row := cell.row; row <= cell.row-currentRow; row++ {
				yOffsets[row] = math.Max(yOffsets[row-1], yOffsets[row])
			}
			currentRow = cell.row
		}

		yOffset := yOffsets[cell.row-1]

		width := xOffsets[cell.col] - xOffsets[cell.col-1]
		height := available.height - float32(yOffset)

		m := cell.messure(size{width, height})

		yOffsets[cell.row] = math.Max(yOffsets[cell.row], yOffset+float64(m.size.height))

		layouts = append(layouts, tableLayout{
			pType:   m.pType,
			cell:    cell,
			size:    size{width, m.size.height},
			xOffset: xOffsets[cell.col-1],
			yOffset: float32(yOffset),
		})
	}

	return layouts
}
