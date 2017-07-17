package controllers

import (
	"fmt"

	"github.com/gamegos/jsend"
	"github.com/nlack/ews-qr-app/RevelEWS/app/models"
	"github.com/revel/revel"
)

type InstructorController struct {
	GorpController
}

func (p InstructorController) Login( /*username string	parameter nur fuer formulare*/ ) revel.Result {

	//fmt.Println(username)
	//fmt.Println("params" + p.Params.Get("username"))

	instructor := new(models.Instructor)
	if err := p.Params.BindJSON(instructor); err != nil { //err
		fmt.Println(err.Error())
	}
	err := p.Txn.SelectOne(instructor,
		`SELECT * FROM Instructor WHERE username = ? and password = ?`, instructor.Username, instructor.Password)
	if err != nil {
		jsend.Wrap(p.Response.Out).Status(404).Message(err.Error()).Send()
		return nil
	}
	type AccessKey struct {
		AccessKey string
	}
	accesskey := AccessKey{instructor.AccessKey}
	jsend.Wrap(p.Response.Out).Data(accesskey).Send()
	return nil //p.RenderJSON(targetInstructor)
}

func (p InstructorController) Show() revel.Result {
	return nil
}
