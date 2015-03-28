package frontend

import (
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"github.com/starboychina/martini-mvc/src/config"
	. "github.com/starboychina/martini-mvc/src/helpers/utilities"
)

type Contrller struct {
	PathOptions
}

func (c Contrller) autoRedirect(session sessions.Session, r render.Render) {
	uri := session.Get(config.SessionRedirect)
	if nil == uri {
		r.Redirect("/", 301)
	} else {
		session.Delete(config.SessionRedirect)
		r.Redirect(uri.(string), 301)
	}
}
