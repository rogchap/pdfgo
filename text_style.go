package pdf

import "github.com/go-text/typesetting/di"

type TextDirection = di.Direction

type TextStyle struct {
	FontSize   float32
	LineHeight float32
	Direction  TextDirection
}
