package tec

import "github.com/reusee/dms"

func init() {
	dms.AddCastType((*func(int, Changeset))(nil), func(fn interface{}, args []interface{}) {
		fn.(func(int, Changeset))(args[0].(int), args[1].(Changeset))
	})
}
