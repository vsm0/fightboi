package main

import (
	"github.com/vsm0/fightboi/api"
	"github.com/vsm0/fightboi/gfx"
	"github.com/vsm0/fightboi/lua"

	"image/color"
	"time"

	px "github.com/gopxl/pixel/v2"
	gl "github.com/gopxl/pixel/v2/backends/opengl"
)

func main() {
	gl.Run(run)
}

func run() {
	cfg := gl.WindowConfig{
		Title: "FightBoi Game System",
		Bounds: px.R(0, 0, 480, 360),
		Resizable: true,
		VSync: true,
	}
	win, err := gl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	defer win.Destroy()

	canvas, err := gfx.NewCanvas(128, 128)
	if err != nil {
		panic(err)
	}

	a := &api.Api{
		Runtime: *lua.New(),
		Canvas: *canvas,
	}

	a.Register(a.Pix())
	a.Run([]byte("pix(50, 50, 1)"), "test")

	tps := time.Duration(time.Second/60)
	timer := time.NewTicker(tps)

	for range timer.C {
		if win.Closed() {
			break
		}
		win.Update()
		win.Clear(color.RGBA{0x1f, 0x1f, 0x1f, 0x1f})
		
		canvas.Draw(win)
	}
}
