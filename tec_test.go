package tec

import (
	"testing"

	"github.com/reusee/dms"
)

func TestNew(t *testing.T) {
	c := New()
	defer c.Close()
	loaded := false
	c.Load(mod{func(loader dms.Loader) {
		loaded = true
	}})
	if !loaded {
		t.Fatal("load")
	}
}
