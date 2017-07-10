package controllers

import "github.com/revel/revel"

type Instructor struct {
	App
}

func (c Instructor) Login() revel.Result {
	return c.Render()
}

func (c Instructor) UserLogin(username, password string, remember bool) revel.Result {
	return nil
}
