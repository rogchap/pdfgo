package skia

import (
	"bytes"
	"fmt"
	"testing"
)

func TestAPI(t *testing.T) {
	p := NewPixmap()
	p.Reset()
	fmt.Printf("%#v\n", p)

	var buf bytes.Buffer
	doc := NewDocument(&buf)
	c := doc.BeginPage(400, 600, &Rect{0, 0, 300, 400})
	_ = c
	doc.Close()
	fmt.Printf("%s\n", buf.Bytes())
}
