package gfxlib

import (
	"github.com/vsm0/fightboi/api"

	"github.com/arnodel/golua/lib/packagelib"
	rt "github.com/arnodel/golua/runtime"
)

func Loader(a *api.App) packagelib.Loader {
	l := packagelib.Loader{}
	l.Name = "gfx"

	l.Load = func(r *rt.Runtime) (rt.Value, func()) {
		pkg := rt.NewTable()
		pkgVal := rt.TableValue(pkg)
		r.SetEnvGoFunc(pkg, "pix", a.Pix, 3, false)
		r.SetEnvGoFunc(pkg, "rect", a.Rect, 5, false)

		return pkgVal, nil
	}

	return l
}
