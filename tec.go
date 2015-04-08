package tec

import (
	"sync"

	"github.com/reusee/dms"
)

type Tec struct {
	*dms.Sys
}

func New() *Tec {
	sys := dms.New()
	mods := []dms.Mod{
		new(ModChangeset),
	}
	wg := new(sync.WaitGroup)
	wg.Add(len(mods))
	for _, mod := range mods {
		mod := mod
		go func() {
			sys.Load(mod)
			wg.Done()
		}()
	}
	wg.Wait()
	return &Tec{
		Sys: sys,
	}
}

func (t *Tec) Close() {
	t.Sys.Close()
}
