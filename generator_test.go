package pdf_test

import (
	"bytes"
	"os"
	"testing"
	"time"

	"rogchap.com/pdf"

	"golang.org/x/exp/trace"
)

type testDoc struct{}

func (d *testDoc) Build(doc *pdf.Doc) {
	doc.Author = "rogchap"
	doc.Title = "Kitchen Sink"

	doc.Creation = time.Now()

	doc.Page(func(page *pdf.Page) {
		// page.PageSize(pdf.PageSizeA4Landscape)
		page.MarginVH(75)
		// page.Color("#f00")

		page.Header().Padding(0, 0, 0, 50).AlignCenter().Text("Header").Color("#f00")
		page.Footer().
			Padding(0, 50, 0, 0).
			AlignRight().
			TextBlock(func(text *pdf.TextBlock) {
				text.Span("Page ")
				text.CurrentPage()
			})

		page.Content().
			// Background("#020").
			VStack(func(stack *pdf.VStack) {
				stack.Space(50)

				stack.Item().Text(textCopy).FontSize(24)
				stack.Item().Text(textCopy).FontSize(34).Italic(true)
				stack.Item().Text(textCopy).FontSize(44)

				stack.Item().HStack(func(hstack *pdf.HStack) {
					hstack.RelativeItem(1).ImageFile("testdata/park1.jpg")
					hstack.RelativeItem(1).AlignCenter().ImageFile("testdata/park2.jpg")
					hstack.RelativeItem(1).AlignRight().ImageFile("testdata/park3.jpg")
				})

				stack.Item().Text(textCopy).FontSize(54).LineHeight(1.5)
				stack.Item().Text(textCopy).FontSize(64)
				stack.Item().Text(textCopy).FontSize(64).FontWeight(pdf.FontWeightLight)
				stack.Item().Text(textCopy).FontSize(64)
				stack.Item().PageBreak()
				stack.Item().Text(textCopy).FontSize(54)
				stack.Item().Text(textCopy).FontSize(54).Bold()
				stack.Item().Text(textCopy).FontSize(54)
				stack.Item().AlignCenter().ImageFile("testdata/park4.jpg")
				stack.Item().Text(textCopy).FontSize(54)
				stack.Item().Text(textCopy).FontSize(54).FontFamily("Times New Roman")
				stack.Item().Text(textCopy).FontSize(54)
				stack.Item().AlignCenter().Text("THE END").FontWeight(pdf.FontWeightBlack).FontSize(74).Color("#df00fa")
			})
	})
	doc.Page(func(page *pdf.Page) {
		page.PageSize(pdf.PageSizeA4Landscape)
		page.MarginVH(75)

		page.Content().HStack(func(hstack *pdf.HStack) {
			hstack.Space(50)

			hstack.RelativeItem(1).Text(textCopy).FontSize(36)
			hstack.RelativeItem(1).Text(textCopy).FontSize(62)
		})
	})
}

func TestGenerator(t *testing.T) {
	fr := trace.NewFlightRecorder()
	fr.Start()

	f, _ := os.Create("output.pdf")
	defer f.Close()
	pdf.Generate(f, &testDoc{})

	var b bytes.Buffer
	fr.WriteTo(&b)
	os.WriteFile("trace.out", b.Bytes(), 0o755)
}

var textCopy = "Welcome to the bustling city streets, where the rhythm of life never seems to slow down. 🏙️From the early morning rush to the late-night revelry, there's always something happening around every corner. The aroma of freshly brewed coffee wafts through the air as people hurry to catch their morning commute 🚶‍♂️🚶‍♀️, while street performers entertain crowds with their mesmerizing tunes. 世欢迎来到丛林. As the sun sets, the cityscape transforms into a kaleidoscope of neon lights, painting the sky with vibrant hues. 🌆Amidst the hustle and bustle, friendships are forged, dreams are chased, and love finds its way into the most unexpected places. So, take a stroll down these lively streets and let the city's energy sweep you off your feet. 🌟💫"
