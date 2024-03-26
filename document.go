package pdf

import "time"

type DocBuilder interface {
	Build(doc *Doc)
}

type Doc struct {
	pages []*Page

	// The document's title.
	Title string
	// The name of the person who created the document.
	Author string
	// The subject of the document.
	Subject string
	// Keywords associated with the document.
	Keywords string
	// If the document was converted to PDF from another format, the name of the conforming product that created the
	// original document from which it was converted.
	Creator string
	// The product that is converting this document to PDF.
	Producer string
	// The date and time the document was created.
	Creation time.Time
	// The date and time the document was most recently modified.
	Modified time.Time

	// If true, include XMP metadata, a document UUID, and sRGB output intent information.
	// This adds length to the document and makes it non-reproducable, but are necessary features for PDF/A-2b conformance
	PDFA bool
	// The DPI (pixels-per-inch) at which features without native PDF support will be rasterized (e.g.
	// draw image with perspective, draw text with perspective, ...) A larger DPI would create a PDF that reflects the
	// original intent with better fidelity, but it can make for larger PDF files too, which would use more memory while
	// rendering, and it would be slower to be processed or sent online or to printer.
	RasterDPI float32
	// Encoding quality controls the trade-off between size and quality.
	// By default this is set to 101 percent, which corresponds to lossless encoding. If this value is set to a
	// value <= 100, and the image is opaque, it will be encoded (using JPEG) with that quality setting.
	EncodingQuality int
}

func (d *Doc) Page(cb func(page *Page)) {
	page := &Page{}
	cb(page)
	d.pages = append(d.pages, page)
}

func (c *Doc) build() drawable {
	if len(c.pages) == 0 {
		return nil
	}

	cont := &container{}
	if len(c.pages) == 1 {
		c.pages[0].build(cont)
		return cont
	}

	cont.VStack(func(stack *VStack) {
		for idx, page := range c.pages {
			if idx != 0 {
				stack.Item().PageBreak()
			}
			pageCont := &container{}
			page.build(pageCont)
			stack.Item().Element(pageCont)
		}
	})
	return cont
}
