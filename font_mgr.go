package pdf

import (
	"path/filepath"

	"github.com/go-text/typesetting/font"
	"github.com/go-text/typesetting/fontscan"
	"rogchap.com/skia"
)

type nilLogger struct{}

func (nilLogger) Printf(format string, args ...any) {}

type fontMgr struct {
	skFontMgr *skia.FontMgr
	hbFontMgr *fontscan.FontMap

	typefaceCache map[string]*skia.Typeface
}

func (f *fontMgr) skTypeface(face font.Face) *skia.Typeface {
	fam, _ := f.hbFontMgr.FontMetadata(face.Font)
	tf, inCache := f.typefaceCache[fam]

	if tf == nil {
		tf = f.skFontMgr.MatchFamily(fam)
	}

	loc := f.hbFontMgr.FontLocation(face.Font)
	if tf == nil {
		// some skia fonts are not stored in the cache with the same normalized string as hb
		// We are assuming that the location is a filename that matches that of the family name
		_, fname := filepath.Split(loc.File)
		tf = f.skFontMgr.MatchFamily(fname[:len(fname)-len(filepath.Ext(fname))])
	}

	if tf == nil {
		tf = f.skFontMgr.CreateFromFile(loc.File, int(loc.Index))
	}

	if tf == nil {
		// fallback to default typeface
		// this assumes the platform always returns a typeface for default
		tf = f.skFontMgr.MatchFamily("")
	}

	if !inCache {
		f.typefaceCache[fam] = tf
	}

	return tf
}

func initFontMgr() *fontMgr {
	var fm fontMgr

	fm.skFontMgr = skia.NewSystemFontMgr()
	fm.hbFontMgr = fontscan.NewFontMap(nilLogger{}) // TODO: Could font logging be helpful to the caller?
	fm.hbFontMgr.UseSystemFonts("")                 // TODO: We need to be able to set the cache directory for a caller

	fm.typefaceCache = make(map[string]*skia.Typeface)

	return &fm
}

// TODO: due to settings of the font managers (like the cache) this
// initialisation should be lazy and use sync.Once
var defaultFontMgr = initFontMgr()
