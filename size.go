package pdf

type size struct {
	width, height float32
}

type sizePlanType uint8

const (
	full sizePlanType = iota
	partial
	wrap
)

type sizePlan struct {
	size  size
	pType sizePlanType
}

var maxSize = size{100_000, 100_000}

type PageSize size

var (
	PageSizeA4          = PageSize{2384.2, 3370.8}
	PageSizeA4Landscape = flip(PageSizeA4)
)

func flip(s PageSize) PageSize {
	return PageSize{
		width:  s.height,
		height: s.width,
	}
}

func asSize(s PageSize) size {
	return size{
		width:  s.width,
		height: s.height,
	}
}
