package backend

import (
	"github.com/martini-contrib/render"
)

func (c Contrller) Default(r render.Render) {

	opt := render.HTMLOptions{
		Layout: c.PathOptions.Layout,
	}
	r.HTML(200, c.ViewPath+"default/index", nil, opt)
}