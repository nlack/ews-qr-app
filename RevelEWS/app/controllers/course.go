package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/gamegos/jsend"
	"github.com/nlack/ews-qr-app/RevelEWS/app/models"
	"github.com/revel/revel"
)

type CourseController struct {
	GorpController
}

func (c CourseController) parseCourse() (models.Course, error) {
	course := models.Course{}
	err := json.NewDecoder(c.Request.Body).Decode(&course)
	return course, err
}

func (c CourseController) Add() revel.Result {

	course := new(models.Course)
	if err := c.Params.BindJSON(course); err != nil { //err
		fmt.Println(err.Error())
	}
	tempInstructor := new(models.Instructor)
	err := c.Txn.SelectOne(tempInstructor, `SELECT * FROM Instructor WHERE accesskey = ?`, course.InstructorKey)
	if err != nil {
		jsend.Wrap(c.Response.Out).Status(404).Message("Instr key: " + err.Error()).Send()
		return nil
	}
	course.Validate(c.Validation)
	if c.Validation.HasErrors() {
		// Do something better here!
		return c.RenderText("You have error in your Course.")
	}
	if err := c.Txn.Insert(course); err != nil {
		return c.RenderText(
			err.Error())
	} else {
		jsend.Wrap(c.Response.Out).Status(200).Send()
		return nil
	}
	// Validate the model
	/*course.Validate(c.Validation)
	if c.Validation.HasErrors() {
		// Do something better here!
		return c.RenderText("You have error in your Course.")
	} else {
		if err := c.Txn.Insert(&course); err != nil {
			return c.RenderText(
				"Error inserting record into database!")
		} else {
			return c.RenderJSON(course)
		}
	}*/
}

func (c CourseController) Get(id int64) revel.Result {
	course := new(models.Course)
	err := c.Txn.SelectOne(course,
		`SELECT * FROM Course WHERE id = ?`, id)
	if err != nil {
		return c.RenderText("Error.  Item probably doesn't exist.")
	}
	return c.RenderJSON(course)
}

func (c CourseController) List() revel.Result {
	lastId := parseIntOrDefault(c.Params.Get("lid"), -1)
	limit := parseUintOrDefault(c.Params.Get("limit"), uint64(25))
	courses, err := c.Txn.Select(models.Course{},
		`SELECT * FROM Course WHERE Id > ? LIMIT ?`, lastId, limit)
	if err != nil {
		return c.RenderText(
			"Error trying to get records from DB.")
	}
	return c.RenderJSON(courses)
}

func (c CourseController) Update(id int64) revel.Result {
	course, err := c.parseCourse()
	if err != nil {
		return c.RenderText("Unable to parse the Course from JSON.")
	}
	// Ensure the Id is set.
	course.Id = id
	success, err := c.Txn.Update(&course)
	if err != nil || success == 0 {
		return c.RenderText("Unable to update bid item.")
	}
	return c.RenderText("Updated %v", id)
}

func (c CourseController) Delete(id int64) revel.Result {
	success, err := c.Txn.Delete(&models.Course{Id: id})
	if err != nil || success == 0 {
		return c.RenderText("Failed to remove Course")
	}
	return c.RenderText("Deleted %v", id)
}

func (c CourseController) AddParticipant() revel.Result {
	return nil
}
