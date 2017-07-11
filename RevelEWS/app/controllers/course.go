package controllers

import (
	"encoding/json"

	"github.com/nlack/ews-qr-app/RevelEWS/app/models"

	"github.com/revel/revel"
)

type Course struct {
	GorpController
}

func (c Course) parseCourse() (models.Course, error) {
	biditem := models.Course{}
	err := json.NewDecoder(c.Request.Body).Decode(&biditem)
	return biditem, err
}

func (c Course) Add() revel.Result {
	if biditem, err := c.parseCourse(); err != nil {
		return c.RenderText("Unable to parse the Course from JSON.")
	} else {
		// Validate the model
		biditem.Validate(c.Validation)
		if c.Validation.HasErrors() {
			// Do something better here!
			return c.RenderText("You have error in your Course.")
		} else {
			if err := c.Txn.Insert(&biditem); err != nil {
				return c.RenderText(
					"Error inserting record into database!")
			} else {
				return c.RenderJSON(biditem)
			}
		}
	}
}

func (c Course) Get(id int64) revel.Result {
	course := new(models.Course)
	err := c.Txn.SelectOne(course,
		`SELECT * FROM Course WHERE id = ?`, id)
	if err != nil {
		return c.RenderText("Error.  Item probably doesn't exist.")
	}
	return c.RenderJSON(course)
}

func (c Course) List() revel.Result {
	lastId := parseIntOrDefault(c.Params.Get("lid"), -1)
	limit := parseUintOrDefault(c.Params.Get("limit"), uint64(25))
	biditems, err := c.Txn.Select(models.Course{},
		`SELECT * FROM Course WHERE Id > ? LIMIT ?`, lastId, limit)
	if err != nil {
		return c.RenderText(
			"Error trying to get records from DB.")
	}
	return c.RenderJSON(biditems)
}

func (c Course) Update(id int64) revel.Result {
	biditem, err := c.parseCourse()
	if err != nil {
		return c.RenderText("Unable to parse the Course from JSON.")
	}
	// Ensure the Id is set.
	biditem.Id = id
	success, err := c.Txn.Update(&biditem)
	if err != nil || success == 0 {
		return c.RenderText("Unable to update bid item.")
	}
	return c.RenderText("Updated %v", id)
}

func (c Course) Delete(id int64) revel.Result {
	success, err := c.Txn.Delete(&models.Course{Id: id})
	if err != nil || success == 0 {
		return c.RenderText("Failed to remove Course")
	}
	return c.RenderText("Deleted %v", id)
}
