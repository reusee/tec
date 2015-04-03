package tec

type Changeset struct {
	Id  int
	Ops []Op
}

type Op struct {
	Type OpType
	Pos  int
	Text []byte
}

type OpType int

const (
	OpInsert OpType = iota
	OpDelete
)

type State interface {
	Apply(Changeset)
	Revert(Changeset)
}
