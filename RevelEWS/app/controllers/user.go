package controllers

import (
	"encoding/json"

	"github.com/nlack/ews-qr-app/RevelEWS/app/models"

	"github.com/revel/revel"
)

type User struct {
	GorpController
}


func (c User) Participant() revel.Result {
	return c.Render()
}

func (c User) Instructor() revel.Result {
	return c.Render()
}

func (c User) parseUser() (models.User, error) {
	user := models.User{}
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	return user, err
}

func (c User) Add() revel.Result {
	if user, err := c.parseUser(); err != nil {
		return c.RenderText("Unable to parse the User from JSON.")
	} else {
		// Validate the model
		user.Validate(c.Validation)
		if c.Validation.HasErrors() {
			// Do something better here!
			return c.RenderText("You have error in your User.")
		} else {
			if err := c.Txn.Insert(&user); err != nil {
				return c.RenderText(
					"Error inserting record into database!")
			} else {
				return c.RenderJSON(user)
			}
		}
	}
}

func (c User) Get(id int64) revel.Result {
	course := new(models.User)
	err := c.Txn.SelectOne(course,
		`SELECT * FROM User WHERE id = ?`, id)
	if err != nil {
		return c.RenderText("Error.  Item probably doesn't exist.")
	}
	return c.RenderJSON(course)
}

func (c User) List() revel.Result {
	lastId := parseIntOrDefault(c.Params.Get("lid"), -1)
	limit := parseUintOrDefault(c.Params.Get("limit"), uint64(25))
	users, err := c.Txn.Select(models.User{},
		`SELECT * FROM User WHERE Id > ? LIMIT ?`, lastId, limit)
	if err != nil {
		return c.RenderText(
			"Error trying to get records from DB.")
	}
	return c.RenderJSON(users)
}

func (c User) Update(id int64) revel.Result {
	user, err := c.parseUser()
	if err != nil {
		return c.RenderText("Unable to parse the User from JSON.")
	}
	// Ensure the Id is set.
	user.Id = id
	success, err := c.Txn.Update(&user)
	if err != nil || success == 0 {
		return c.RenderText("Unable to update User.")
	}
	return c.RenderText("Updated %v", id)
}

func (c User) Delete(id int64) revel.Result {
	success, err := c.Txn.Delete(&models.User{Id: id})
	if err != nil || success == 0 {
		return c.RenderText("Failed to remove User")
	}
	return c.RenderText("Deleted %v", id)
}
