# PDF Go

## Feature

* Powerful page layouts
* Fonts manager
* Font family substitutions
* ...

## Usage

This project is a WIP; no release has been made, and is pre-v0.1

## Example
```go
type myDoc struct {}

func (d *myDoc) Build(c *pdf.DocContainer) {
	c.Page(func(page *pdf.Page) {
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
