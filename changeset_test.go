package tec

import (
	"testing"

	"github.com/reusee/dms"
)

func TestChangeset(t *testing.T) {
	c := New()
	c.Load(mod{func(loader dms.Loader) {
		var apply func(Changeset)
		loader.Require("apply changeset", &apply)
		apply(nil)

		var applyCast *dms.Cast
		loader.Require("changeset apply cast", &applyCast)
		called := false
		applyCast.Add(func(i int, c Changeset) {
			called = true
		})
		apply(nil)
		if !called {
			t.Fail()
		}

		apply(nil)

		var revert func()
		loader.Require("revert changeset", &revert)
		revert()
		revert()
		revert()
		revert()
	}})
}
