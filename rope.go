package tec

import "github.com/reusee/dms"

func init() {
	addMod(new(ModRope))
}

type ModRope struct {
}

func (m *ModRope) Load(loader dms.Loader) {
	loader.Provide("mod rope", m)

	var applyCast, revertCast *dms.Cast
	loader.Require("changeset apply cast", &applyCast)
	loader.Require("changeset revert cast", &revertCast)
	applyCast.Add(m.OnApply)
	revertCast.Add(m.OnRevert)
}

func (m *ModRope) OnApply(i int, cs Changeset) {
	//TODO
}

func (m *ModRope) OnRevert(i int, cs Changeset) {
	//TODO
}
