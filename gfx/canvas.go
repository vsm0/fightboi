package gfx

import (
	"errors"
	"image/color"
	"math"

	px "github.com/gopxl/pixel/v2"
	gl "github.com/gopxl/pixel/v2/backends/opengl"
)

type Canvas struct {
	Data []color.RGBA
	px.Sprite
}

func NewCanvas(w, h float64) (*Canvas, error) {
	c := &Canvas{}

	if w < 1 || h < 1 {
		return c, errors.New("invalid width/height")
	}

	r := px.R(0, 0, w, h)
	pd := px.MakePictureData(r)
	c.Data = pd.Pix

	for i := range c.Data {
		c.Data[i] = color.RGBA{A: 0xff}
	}

	c.Sprite = *px.NewSprite(pd, r)

	return c, nil
}

func (c *Canvas) Set(x, y int, col color.RGBA) error {
	w := int(c.Frame().W())
	h := int(c.Frame().H())
	if x < 0 || y < 0 || x >= w || y >= h {
		return errors.New("invalid x/y")
	}

	c.Data[x + y * w] = col
	return nil
}

func (c *Canvas) Draw(win *gl.Window) {
	ww := win.Bounds().W()
	wh := win.Bounds().H()
	w := c.Frame().W()
	h := c.Frame().H()
	scale := math.Min(ww / w, wh / h)
	pos := win.Bounds().Center()
	mtx := px.IM.Moved(pos).ScaledXY(pos, px.V(scale, scale))

	c.Sprite.Draw(win, mtx)
}
