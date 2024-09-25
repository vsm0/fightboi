package api

import (
	"errors"
	"image/color"

	rt "github.com/arnodel/golua/runtime"
)

func (a *App) Rect(r *rt.Thread, c *rt.GoCont) (rt.Cont, error) {
	var x, y, w, h, col int64

	err := c.CheckNArgs(5)
	if err == nil {
		x, err = c.IntArg(0)
	}
	if err == nil {
		y, err = c.IntArg(1)
	}
	if err == nil {
		w, err = c.IntArg(2)
	}
	if err == nil {
		h, err = c.IntArg(3)
	}
	if err == nil {
		col, err = c.IntArg(4)
	}

	if col < 0 || col >= 16 {
		return nil, errors.New("invalid color index")
	}

	for yy := y; yy < y + h; yy++ {
		for xx := x; xx < x + w; xx++ {
			a.Set(int(xx), int(yy), color.RGBA{0xff, 0xff, 0xff, 0xff})
		}
	}

	return c.PushingNext(nil), nil
}
