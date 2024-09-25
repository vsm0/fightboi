package api

import (
	"github.com/vsm0/fightboi/gfx"
	"github.com/vsm0/fightboi/lua"
)

type Api struct {
	lua.Runtime
	gfx.Canvas
}
