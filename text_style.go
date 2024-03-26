package pdf

import "github.com/go-text/typesetting/di"

type TextDirection = di.Direction

type FontWeight int

const (
	FontWeightThin       FontWeight = 100
	FontWeightExtraLight            = 200
	FontWeightLight                 = 300
	FontWeightNormal                = 400
	FontWeightMedium                = 500
	FontWeightSemiBold              = 600
	FontWeightBold                  = 700
	FontWeightExtraBold             = 800
	FontWeightBlack                 = 900
	FontWeightExtraBlack            = 1000
)

type TextStyle struct {
	FontSize   float32
	LineHeight float32
	Color      string
	FontFamily []string
	FontWeight FontWeight
	Italic     bool

	Direction TextDirection
}
