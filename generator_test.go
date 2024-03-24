package pdf_test

import (
	"os"
	"testing"

	"rogchap.com/pdf"
)

type testDoc struct{}

func (d *testDoc) Build(c *pdf.DocContainer) {
	c.Page(func(page *pdf.Page) {
		// page.PageSize(pdf.PageSizeA4Landscape)
		page.MarginV(200)
		page.MarginH(75)
		// page.Color("#f00")

		page.Content().
			// Background("#020").
			VStack(func(stack *pdf.VStack) {
				stack.Space(50)

				stack.Item().ImageFile("testdata/park1.jpg")
				stack.Item().Text(textCopy).FontSize(24)
				stack.Item().Text(textCopy).FontSize(34)
				stack.Item().Text(textCopy).FontSize(44)
				stack.Item().Text(textCopy).FontSize(54)
				stack.Item().Text(textCopy).FontSize(64)
				stack.Item().Text(textCopy).FontSize(64)
				stack.Item().Text(textCopy).FontSize(64)
				stack.Item().PageBreak()
				stack.Item().Text(textCopy).FontSize(64)
				stack.Item().Text(textCopy).FontSize(54)
				stack.Item().Text(textCopy).FontSize(54)
				stack.Item().ImageFile("testdata/park2.jpg")
				stack.Item().Text(textCopy).FontSize(54)
				stack.Item().Text(textCopy).FontSize(54)
				stack.Item().Text(textCopy).FontSize(54)
			})
	})
	c.Page(func(page *pdf.Page) {
		page.PageSize(pdf.PageSizeA4Landscape)
		page.Content().HStack(func(hstack *pdf.HStack) {
			hstack.Item().Text(textCopy)
			hstack.Item().Text(textCopy)
		})
	})
}

func TestGenerator(t *testing.T) {
	f, _ := os.Create("output.pdf")
	defer f.Close()
	pdf.Generate(f, &testDoc{})
}

var textCopy = "Welcome to the bustling city streets, where the rhythm of life never seems to slow down. 🏙️ From the early morning rush to the late-night revelry, there's always something happening around every corner. The aroma of freshly brewed coffee wafts through the air as people hurry to catch their morning commute 🚶‍♂️🚶‍♀️, while street performers entertain crowds with their mesmerizing tunes. 世界您好 As the sun sets, the cityscape transforms into a kaleidoscope of neon lights, painting the sky with vibrant hues. 🌆 Amidst the hustle and bustle, friendships are forged, dreams are chased, and love finds its way into the most unexpected places. So, take a stroll down these lively streets and let the city's energy sweep you off your feet. 🌟💫"
