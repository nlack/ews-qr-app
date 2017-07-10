package controllers

import "github.com/revel/revel"

type Participant struct {
	App
}

func (c Participant) Login() revel.Result {
	return c.Render()
}

func (c Participant) LoginUser(username, password string, remember bool) revel.Result {
	return nil
}
