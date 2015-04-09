package tec

import (
	"testing"

	"github.com/reusee/dms"
)

func TestChangeset(t *testing.T) {
	c := New()
	c.Load(mod{func(loader dms.Loader) {
		var modChangeset *ModChangeset
		loader.Require("mod changeset", &modChangeset)

		var apply func(Changeset)
		loader.Require("apply changeset", &apply)
		apply(nil)
		if len(modChangeset.Changesets) != 1 {
			t.Fatal("changesets len")
		}
		if modChangeset.csIndex != 0 {
			t.Fatal("csIndex")
		}

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
		if len(modChangeset.Changesets) != 2 {
			t.Fatal("changesets len")
		}
		if modChangeset.csIndex != 1 {
			t.Fatal("csIndex")
		}

		apply(nil)
		if len(modChangeset.Changesets) != 3 {
			t.Fatal("changesets len")
		}
		if modChangeset.csIndex != 2 {
			t.Fatal("csIndex")
		}

		var revert func()
		loader.Require("revert changeset", &revert)
		revert()
		if modChangeset.csIndex != 1 {
			t.Fatal("csIndex")
		}
		revert()
		if modChangeset.csIndex != 0 {
			t.Fatal("csIndex")
		}
		revert()
		if modChangeset.csIndex != -1 {
			t.Fatal("csIndex")
		}
		revert()
		if modChangeset.csIndex != -1 {
			t.Fatal("csIndex")
		}
	}})
}
