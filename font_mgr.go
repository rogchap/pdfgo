package pdf

import (
	"path/filepath"

	"github.com/go-text/typesetting/font"
	"github.com/go-text/typesetting/fontscan"
	"github.com/go-text/typesetting/opentype/api/metadata"
	"rogchap.com/skia"
)

type nilLogger struct{}

func (nilLogger) Printf(format string, args ...any) {}

type fontMgr struct {
	skFontMgr *skia.FontMgr
	hbFontMgr *fontscan.FontMap

	skFamMap map[string]string
}

func (f *fontMgr) skTypeface(face font.Face) *skia.Typeface {
	fam, aspect := f.hbFontMgr.FontMetadata(face.Font)

	fStyle := skFontStyle(aspect)

	var tf *skia.Typeface
	if skFam, ok := f.skFamMap[fam]; ok {
		tf = f.skFontMgr.MatchFamilyStyle(skFam, fStyle)
	}

	loc := f.hbFontMgr.FontLocation(face.Font)
	if tf == nil {
		// some skia fonts are not stored in the cache with the same normalized string as hb
		// We are assuming that the location is a filename that matches that of the family name
		_, fname := filepath.Split(loc.File)
		tf = f.skFontMgr.MatchFamilyStyle(fname[:len(fname)-len(filepath.Ext(fname))], fStyle)
	}

	if tf == nil {
		// fallback to default typeface
		// this assumes the platform always returns a typeface for default
		tf = f.skFontMgr.MatchFamilyStyle("", fStyle)
	}

	return tf
}

func skFontStyle(aspect metadata.Aspect) *skia.FontStyle {
	weight := skia.FontStyleWeight(aspect.Weight)

	var width skia.FontStyleWidth
	switch aspect.Stretch {
	case metadata.StretchUltraCondensed:
		width = skia.FontStyleWidthUltraCondensed
	case metadata.StretchExtraCondensed:
		width = skia.FontStyleWidthExtraCondensed
	case metadata.StretchCondensed:
		width = skia.FontStyleWidthCondensed
	case metadata.StretchSemiCondensed:
		width = skia.FontStyleWidthSemiCondensed
	case metadata.StretchNormal:
		width = skia.FontStyleWidthNormal
	case metadata.StretchSemiExpanded:
		width = skia.FontStyleWidthSemiExpanded
	case metadata.StretchExpanded:
		width = skia.FontStyleWidthExpanded
	case metadata.StretchExtraExpanded:
		width = skia.FontStyleWidthExpanded
	case metadata.StretchUltraExpanded:
		width = skia.FontStyleWidthUltraExpanded
	}

	slant := skia.FontStyleSlantUright
	if aspect.Style == metadata.StyleItalic {
		slant = skia.FontStyleSlantItalic
	}

	fStyle := skia.NewFontStyle(weight, width, slant)
	return fStyle
}

func initFontMgr() *fontMgr {
	var fm fontMgr

	fm.skFontMgr = skia.NewSystemFontMgr()
	fm.hbFontMgr = fontscan.NewFontMap(nilLogger{}) // TODO: Could font logging be helpful to the caller?
	fm.hbFontMgr.UseSystemFonts("")                 // TODO: We need to be able to set the cache directory for a caller

	// setup a map between the hb font name and the sk font name
	skFontFams := fm.skFontMgr.Families()
	fm.skFamMap = make(map[string]string, len(skFontFams))
	for _, fam := range skFontFams {
		fm.skFamMap[metadata.NormalizeFamily(fam)] = fam
	}

	return &fm
}

// TODO: due to settings of the font managers (like the cache) this
// initialisation should be lazy and use sync.Once
var defaultFontMgr = initFontMgr()
