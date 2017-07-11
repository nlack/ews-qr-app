package controllers

import (
	"github.com/nlack/ews-qr-app/RevelEWS/app/routes"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Impressum() revel.Result {
	return c.Render()
}

func (c App) Login(username, password string, instructor bool) revel.Result {

	//if instructor redirect to course list
	//if no instructor redirect to qr code
	if instructor {
		return c.Redirect(routes.User.Instructor())
	} else {
		return c.Redirect(routes.User.Participant())
	}

}
