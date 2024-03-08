package skia

import (
	"fmt"
	"testing"
)

func TestAPI(t *testing.T) {
	p := NewPixmap()
	p.Reset()
	fmt.Printf("%#v\n", p)
}
