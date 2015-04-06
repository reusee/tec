package tec

import "sync"

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

type State interface {
	Apply(int, Changeset)
	Revert(int, Changeset)
}

type Tec struct {
	Changesets []Changeset
	csIndex    int
	States     [][][]State
	jobs       chan func()
	jobWg      *sync.WaitGroup
	closed     chan struct{}
}

func New(states [][][]State) *Tec {
	// start job runners
	max := 0
	for _, sss := range states {
		for _, ss := range sss {
			if l := len(ss); l > max {
				max = l
			}
		}
	}
	jobs := make(chan func())
	closed := make(chan struct{})
	wg := new(sync.WaitGroup)
	for i := 0; i < max; i++ {
		go func() {
			for {
				select {
				case <-closed:
					return
				case job := <-jobs:
					job()
					wg.Done()
				}
			}
		}()
	}

	return &Tec{
		csIndex: -1,
		States:  states,
		jobs:    jobs,
		jobWg:   wg,
		closed:  closed,
	}
}

func (t *Tec) Close() {
	close(t.closed)
}

func (t *Tec) Apply(c Changeset) {
	t.Changesets = t.Changesets[:t.csIndex+1]
	t.Changesets = append(t.Changesets, c)
	t.csIndex = len(t.Changesets) - 1
	for _, sss := range t.States {
		t.jobWg.Add(len(sss))
		for _, ss := range sss {
			ss := ss
			t.jobs <- func() {
				for _, s := range ss {
					s.Apply(t.csIndex, c)
				}
			}
		}
		t.jobWg.Wait()
	}
}

func (t *Tec) Revert() {
	if t.csIndex < 0 {
		return
	}
	c := t.Changesets[t.csIndex]
	for _, sss := range t.States {
		t.jobWg.Add(len(sss))
		for _, ss := range sss {
			ss := ss
			t.jobs <- func() {
				for _, s := range ss {
					s.Revert(t.csIndex, c)
				}
			}
		}
		t.jobWg.Wait()
	}
	t.Changesets = t.Changesets[:len(t.Changesets)-1]
	t.csIndex--
}
