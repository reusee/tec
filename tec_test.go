package tec

import "testing"

func TestNew(t *testing.T) {
	c := New()
	defer c.Close()
}
