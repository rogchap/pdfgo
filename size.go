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

var maxSize = size{10_000, 10_000} // is this big enough? don't want it to be too big

var pageSizeA4 = size{2384.2, 3370.8}
