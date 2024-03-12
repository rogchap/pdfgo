package pdf

type pageContext struct {
	currentPage int
}

func (p *pageContext) Inc() {
	p.currentPage += 1
}

func (p *pageContext) Dec() {
	p.currentPage -= 1
}
