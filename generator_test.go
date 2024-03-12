package pdf_test

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"rogchap.com/pdf"
)

type testDoc struct{}

func (d *testDoc) Build(c *pdf.DocContainer) {
	c.Page(func(page *pdf.Page) {
		page.Size(400, 600)
		page.Content().
			Background("green").
			Text(func(text *pdf.TextBlock) {
				text.Span("Hello World")
				text.Span("Something else")
			})
	})
}

func TestGenerator(t *testing.T) {
	var buf bytes.Buffer
	doc := testDoc{}

	pdf.Generate(&buf, &doc)
	fmt.Printf("%s\n", buf.Bytes())
	os.WriteFile("output.pdf", buf.Bytes(), 0o666)
}
