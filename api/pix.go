package api

import (
	"errors"
	"image/color"
	
	rt "github.com/arnodel/golua/runtime"
)

func (a *App) Pix(t *rt.Thread, c *rt.GoCont) (rt.Cont, error) {
	var x, y, col int64

	err := c.CheckNArgs(3)
	if err == nil {
		x, err = c.IntArg(0)
	}
	if err == nil {
		y, err = c.IntArg(1)
	}
	if err == nil {
		col, err = c.IntArg(2)
	}
	if err != nil {
		return nil, err
	}

	if col < 0 || col >= 16 {
		return nil, errors.New("invalid color index")
	}

	// for now, just set white
	a.Set(int(x), int(y), color.RGBA{0xff, 0xff, 0xff, 0xff})

	return c.PushingNext(nil), nil
}
