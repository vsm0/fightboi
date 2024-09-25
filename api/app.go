package api

import (
	"github.com/vsm0/fightboi/gfx"
	"github.com/vsm0/fightboi/lua"
)

type App struct {
	lua.Runtime
	gfx.Canvas
}

func New(w, h float64) (*App, error) {
	a := &App{
		Runtime: *lua.New(),
	}

	canvas, err := gfx.NewCanvas(w, h)
	a.Canvas = *canvas
	
	return a, err
}
