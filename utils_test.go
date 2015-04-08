package tec

import "github.com/reusee/dms"

type mod struct {
	load func(dms.Loader)
}

func (m mod) Load(loader dms.Loader) {
	m.load(loader)
}
