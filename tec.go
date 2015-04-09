package tec

import (
	"sync"
	"time"

	"github.com/reusee/dms"
)

func init() {
	dms.MaxResolveTime = time.Second * 1
}

type Tec struct {
	*dms.Sys
}

var mods = []dms.Mod{}

func addMod(mod dms.Mod) {
	mods = append(mods, mod)
}

func New() *Tec {
	sys := dms.New()
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
