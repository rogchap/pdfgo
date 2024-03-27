# PDF Go

PDF Go is a powerful, open source, PDF generation library for the Go programming language.

## Feature

* Powerful layout engine
    * Page layers
    * horizontal and vertical stacks
    * Page headers and footers
    * Automatically wrap content over multiple pages
* Comprehensive text rendering
    * Text styles (italic, bold, color, line height etc)
    * Fonts manager including system fonts (yes, color emoji support 😍)
    * Font family substitutions (fallback)
    * Font subsetting (reduce PDF file size)
    * Font shaping (support Arabic fonts)
    * Right-to-left (RTL) content (Note: WIP)
* ...

## Usage

This project is a WIP; no release has been made, and is pre-v0.1

## Example
```go
type myDoc struct {}

func (d *myDoc) Build(b *pdf.Doc) {
	d.Page(func(page *pdf.Page) {
        page.Content().
            Text("PDF Go").FontSize(64)
    }
}

func main() {
	f, _ := os.Create("output.pdf")
	defer f.Close()
	pdf.Generate(f, &myDoc{})
}
```

## Acknowledgements 

This library is has been inspired by the great work done by [Marcin Ziąbek](https://github.com/MarcinZiabek) and the [QuestPDF](https://www.questpdf.com/) community.
