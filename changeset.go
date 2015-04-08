package tec

import "github.com/reusee/dms"

type ModChangeset struct {
	Changesets []Changeset
	csIndex    int
	applyCast  *dms.Cast
	revertCast *dms.Cast
}

type Changeset []Change

type Change struct {
	Op   Op
	Pos  int
	Text []byte
}

type Op int

const (
	OpInsert Op = iota
	OpDelete
)

func (m *ModChangeset) Load(loader dms.Loader) {
	m.applyCast = dms.NewCast((*func(int, Changeset))(nil))
	loader.Provide("changeset apply cast", m.applyCast)
	m.revertCast = dms.NewCast((*func(int, Changeset))(nil))
	loader.Provide("changeset revert cast", m.revertCast)

	loader.Provide("apply changeset", m.Apply)
	loader.Provide("revert changeset", m.Revert)
}

func (m *ModChangeset) Apply(c Changeset) {
	if m.csIndex > 0 {
		m.Changesets = m.Changesets[:m.csIndex+1] // clear reverted changesets
	}
	m.Changesets = append(m.Changesets, c)
	m.csIndex = len(m.Changesets) - 1
	m.applyCast.Pcall(m.csIndex, c)
}

func (m *ModChangeset) Revert() {
	if m.csIndex < 0 {
		return
	}
	c := m.Changesets[m.csIndex]
	m.revertCast.Pcall(m.csIndex, c)
	m.csIndex--
}
