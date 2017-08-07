// Package models contains the types for schema 'testtt'.
package models

// GENERATED BY XO. DO NOT EDIT.

import (
	"errors"
	"time"
)

// Course represents a row from 'testtt.course'.
type Course struct {
	ID           int       `json:"id"`            // id
	Name         string    `json:"name"`          // name
	Date         time.Time `json:"date"`          // date
	Participants string    `json:"participants"`  // participants
	InstructorID int       `json:"instructor_id"` // instructor_id

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Course exists in the database.
func (c *Course) Exists() bool {
	return c._exists
}

// Deleted provides information if the Course has been deleted from the database.
func (c *Course) Deleted() bool {
	return c._deleted
}

// Insert inserts the Course to the database.
func (c *Course) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO testtt.course (` +
		`name, date, participants, instructor_id` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, c.Name, c.Date, c.Participants, c.InstructorID)
	res, err := db.Exec(sqlstr, c.Name, c.Date, c.Participants, c.InstructorID)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	c.ID = int(id)
	c._exists = true

	return nil
}

// Update updates the Course in the database.
func (c *Course) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !c._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if c._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE testtt.course SET ` +
		`name = ?, date = ?, participants = ?, instructor_id = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, c.Name, c.Date, c.Participants, c.InstructorID, c.ID)
	_, err = db.Exec(sqlstr, c.Name, c.Date, c.Participants, c.InstructorID, c.ID)
	return err
}

// Save saves the Course to the database.
func (c *Course) Save(db XODB) error {
	if c.Exists() {
		return c.Update(db)
	}

	return c.Insert(db)
}

// Delete deletes the Course from the database.
func (c *Course) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !c._exists {
		return nil
	}

	// if deleted, bail
	if c._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM testtt.course WHERE id = ?`

	// run query
	XOLog(sqlstr, c.ID)
	_, err = db.Exec(sqlstr, c.ID)
	if err != nil {
		return err
	}

	// set deleted
	c._deleted = true

	return nil
}

// Instructor returns the Instructor associated with the Course's InstructorID (instructor_id).
//
// Generated from foreign key 'course_ibfk_1'.
func (c *Course) Instructor(db XODB) (*Instructor, error) {
	return InstructorByID(db, c.InstructorID)
}

// CourseByID retrieves a row from 'testtt.course' as a Course.
//
// Generated from index 'course_id_pkey'.
func CourseByID(db XODB, id int) (*Course, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, date, participants, instructor_id ` +
		`FROM testtt.course ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	c := Course{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&c.ID, &c.Name, &c.Date, &c.Participants, &c.InstructorID)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

// CoursesByInstructorID retrieves a row from 'testtt.course' as a Course.
//
// Generated from index 'instructor_id'.
func CoursesByInstructorID(db XODB, instructorID int) ([]*Course, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, date, participants, instructor_id ` +
		`FROM testtt.course ` +
		`WHERE instructor_id = ?`

	// run query
	XOLog(sqlstr, instructorID)
	q, err := db.Query(sqlstr, instructorID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Course{}
	for q.Next() {
		c := Course{
			_exists: true,
		}

		// scan
		err = q.Scan(&c.ID, &c.Name, &c.Date, &c.Participants, &c.InstructorID)
		if err != nil {
			return nil, err
		}

		res = append(res, &c)
	}

	return res, nil
}
