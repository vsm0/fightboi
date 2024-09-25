package lua

import (
	"os"

	"github.com/arnodel/golua/lib"
	"github.com/arnodel/golua/lib/packagelib"
	rt "github.com/arnodel/golua/runtime"
)

type Runtime struct {
	rt.Runtime
}

type ApiFunc struct {
	Fun func(t *rt.Thread, c *rt.GoCont) (rt.Cont, error)
	Name string
	Argc int
	Variadic bool
}

func New() *Runtime {
	return &Runtime{
		Runtime: *rt.New(os.Stdout),
	}
}

func (r *Runtime) Load(loaders ...packagelib.Loader) {
	lib.LoadLibs(&r.Runtime, loaders...)
}

func (r *Runtime) Register(globals ...ApiFunc) {
	for _, g := range globals {
		r.SetEnvGoFunc(r.GlobalEnv(), g.Name, g.Fun, g.Argc, g.Variadic)
	}
}

func (r *Runtime) Run(code []byte, debug string) error {
	chunk, err := r.CompileAndLoadLuaChunk(debug, code, rt.TableValue(r.GlobalEnv()))
	if err != nil {
		return err
	}

	_, err = rt.Call1(r.MainThread(), rt.FunctionValue(chunk))

	return err
}
