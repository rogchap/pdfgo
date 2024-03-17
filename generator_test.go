package pdf_test

import (
	"bytes"
	"os"
	"testing"

	"rogchap.com/pdf"
)

type testDoc struct{}

func (d *testDoc) Build(c *pdf.DocContainer) {
	c.Page(func(page *pdf.Page) {
		page.Size(400, 600)
		page.Content().
			// Background("green").
			Text(func(text *pdf.TextBlock) {
				text.Span("Welcome to the bustling city streets, where the rhythm of life never seems to slow down. 🏙️ From the early morning rush to the late-night revelry, there's always something happening around every corner. The aroma of freshly brewed coffee wafts through the air as people hurry to catch their morning commute 🚶‍♂️🚶‍♀️, while street performers entertain crowds with their mesmerizing tunes. 世界您好 As the sun sets, the cityscape transforms into a kaleidoscope of neon lights, painting the sky with vibrant hues. 🌆 Amidst the hustle and bustle, friendships are forged, dreams are chased, and love finds its way into the most unexpected places. So, take a stroll down these lively streets and let the city's energy sweep you off your feet. 🌟💫").
					FontSize(24)
			})
	})
}

func TestGenerator(t *testing.T) {
	var buf bytes.Buffer
	doc := testDoc{}

	pdf.Generate(&buf, &doc)
	// fmt.Printf("%s\n", buf.Bytes())
	os.WriteFile("output.pdf", buf.Bytes(), 0o666)
}
