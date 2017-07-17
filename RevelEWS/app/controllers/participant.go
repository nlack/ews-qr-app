package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/gamegos/jsend"
	"github.com/nlack/ews-qr-app/RevelEWS/app/models"
	"github.com/revel/revel"
)

type ParticipantController struct {
	GorpController
}

func (c ParticipantController) parseParticipant() (models.Participant, error) {
	participant := models.Participant{}
	err := json.NewDecoder(c.Request.Body).Decode(&participant)
	return participant, err
}

func (c ParticipantController) Add() revel.Result {
	if participant, err := c.parseParticipant(); err != nil {
		return c.RenderText("Unable to parse the Participant from JSON.")
	} else {
		// Validate the model
		participant.Validate(c.Validation)
		if c.Validation.HasErrors() {
			// Do something better here!
			return c.RenderText("You have error in your Participant.")
		} else {
			if err := c.Txn.Insert(&participant); err != nil {
				return c.RenderText(
					"Error inserting record into database!")
			} else {
				return c.RenderJSON(participant)
			}
		}
	}
}

func (p ParticipantController) Login( /*username string	parameter nur fuer formulare*/ ) revel.Result {

	//fmt.Println(username)
	//fmt.Println("params" + p.Params.Get("username"))

	participant := new(models.Participant)
	if err := p.Params.BindJSON(participant); err != nil { //err
		fmt.Println(err.Error())
	}
	err := p.Txn.SelectOne(participant,
		`SELECT * FROM Participant WHERE username = ? and password = ?`, participant.Username, participant.Password)
	if err != nil {
		jsend.Wrap(p.Response.Out).Status(404).Message(err.Error()).Send()
		return nil
	}
	type AccessKey struct {
		AccessKey string
	}
	accesskey := AccessKey{participant.AccessKey}
	jsend.Wrap(p.Response.Out).Data(accesskey).Send()
	return nil //p.RenderJSON(targetParticipant)
}

func (p ParticipantController) Show() revel.Result {
	participant := new(models.Participant)
	if err := p.Params.BindJSON(participant); err != nil { //err
		fmt.Println(err.Error())
	}
	err := p.Txn.SelectOne(participant,
		`SELECT * FROM Participant WHERE id = ? AND accesskey = ?`, p.Params.Get("id"), participant.AccessKey)
	if err != nil {
		jsend.Wrap(p.Response.Out).Status(404).Message(err.Error()).Send()
		return nil
	}
	type Participant struct {
		Firstname string
		Lastname  string
		QRHash    string
	}
	participant2 := Participant{participant.Firstname, participant.Lastname, participant.QRHash}
	jsend.Wrap(p.Response.Out).Data(participant2).Send()
	return nil
}

/*
REQUEST

{
  "name": "TestUser1",
  "password": "test12345#?!\"_-/"
}
RESPONSE

{
  "code": "OK|ERROR",
  "message": "blah blah blah" // bei Error leer.,
  "accesskey": "2039m4c8094875043mxxncowtn" // bei Error nicht vorhanden.
}

*/
