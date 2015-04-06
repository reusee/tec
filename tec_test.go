package tec

import (
	"sync/atomic"
	"testing"
)

func TestNew(t *testing.T) {
	c := New(nil)
	defer c.Close()
}

type fooState struct {
	i *int64
	n int64
}

func (f fooState) Apply(n int, c Changeset) {
	atomic.AddInt64(f.i, f.n)
}
func (f fooState) Revert(n int, c Changeset) {
	atomic.AddInt64(f.i, -f.n)
}

func TestApplyAndRevert(t *testing.T) {
	var n int64
	c := New([][][]State{
		[][]State{
			[]State{fooState{&n, 1}},
			[]State{fooState{&n, 2}},
		},
		[][]State{
			[]State{fooState{&n, 4}, fooState{&n, 8}},
		},
	})
	defer c.Close()
	c.Apply(nil)
	if n != 15 {
		t.Fail()
	}
	c.Revert()
	if n != 0 {
		t.Fail()
	}
	c.Revert()
	if n != 0 {
		t.Fail()
	}
	c.Apply(nil)
	if n != 15 {
		t.Fail()
	}
	c.Revert()
	if n != 0 {
		t.Fail()
	}
	c.Revert()
	if n != 0 {
		t.Fail()
	}
}
