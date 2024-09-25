package lua

import (
	"os"

	"github.com/arnodel/golua/lib"
	"github.com/arnodel/golua/lib/coroutine"
	"github.com/arnodel/golua/lib/mathlib"
	"github.com/arnodel/golua/lib/packagelib"
	"github.com/arnodel/golua/lib/stringlib"
	"github.com/arnodel/golua/lib/tablelib"
	"github.com/arnodel/golua/lib/utf8lib"
	rt "github.com/arnodel/golua/runtime"
)

type Runtime struct {
	Lua *rt.Runtime
}

func New() *Runtime {
	r := &Runtime{
		Lua: rt.New(os.Stdout),
	}

	r.Load(
		packagelib.LibLoader,
		coroutine.LibLoader,
		mathlib.LibLoader,
		stringlib.LibLoader,
		tablelib.LibLoader,
		utf8lib.LibLoader,
	)

	return r
}

func (r *Runtime) Load(loaders ...packagelib.Loader) {
	lib.LoadLibs(r.Lua, loaders...)
}

func (r *Runtime) Run(code []byte, debug string) error {
	chunk, err := r.Lua.CompileAndLoadLuaChunk(debug, code, rt.TableValue(r.Lua.GlobalEnv()))
	if err != nil {
		return err
	}

	_, err = rt.Call1(r.Lua.MainThread(), rt.FunctionValue(chunk))

	return err
}
