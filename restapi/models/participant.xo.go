// Package models contains the types for schema 'testtt'.
package models

// GENERATED BY XO. DO NOT EDIT.

import "errors"

// Participant represents a row from 'testtt.participant'.
type Participant struct {
	ID        int    `json:"id"`                       // id
	Name      string `json:"name"`                     // name
	Password  string `json:"password"`                 // password
	Firstname string `json:"firstname"`                // firstname
	Lastname  string `json:"lastname"`                 // lastname
	Qrhash    string `json:"qrhash" validate:"len=25"` // qrhash

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Participant exists in the database.
func (p *Participant) Exists() bool {
	return p._exists
}

// Deleted provides information if the Participant has been deleted from the database.
func (p *Participant) Deleted() bool {
	return p._deleted
}

// Insert inserts the Participant to the database.
func (p *Participant) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if p._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO testtt.participant (` +
		`name, password, firstname, lastname, qrhash` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, p.Name, p.Password, p.Firstname, p.Lastname, p.Qrhash)
	res, err := db.Exec(sqlstr, p.Name, p.Password, p.Firstname, p.Lastname, p.Qrhash)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	p.ID = int(id)
	p._exists = true

	return nil
}

// Update updates the Participant in the database.
func (p *Participant) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !p._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if p._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE testtt.participant SET ` +
		`name = ?, password = ?, firstname = ?, lastname = ?, qrhash = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, p.Name, p.Password, p.Firstname, p.Lastname, p.Qrhash, p.ID)
	_, err = db.Exec(sqlstr, p.Name, p.Password, p.Firstname, p.Lastname, p.Qrhash, p.ID)
	return err
}

// Save saves the Participant to the database.
func (p *Participant) Save(db XODB) error {
	if p.Exists() {
		return p.Update(db)
	}

	return p.Insert(db)
}

// Delete deletes the Participant from the database.
func (p *Participant) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !p._exists {
		return nil
	}

	// if deleted, bail
	if p._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM testtt.participant WHERE id = ?`

	// run query
	XOLog(sqlstr, p.ID)
	_, err = db.Exec(sqlstr, p.ID)
	if err != nil {
		return err
	}

	// set deleted
	p._deleted = true

	return nil
}

// ParticipantByName retrieves a row from 'testtt.participant' as a Participant.
//
// Generated from index 'name'.
func ParticipantByName(db XODB, name string) (*Participant, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, password, firstname, lastname, qrhash ` +
		`FROM testtt.participant ` +
		`WHERE name = ?`

	// run query
	XOLog(sqlstr, name)
	p := Participant{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, name).Scan(&p.ID, &p.Name, &p.Password, &p.Firstname, &p.Lastname, &p.Qrhash)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func ParticipantByNameAndPW(db XODB, name string, password string) (*Participant, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, password, firstname, lastname, qrhash ` +
		`FROM testtt.participant ` +
		`WHERE name = ? AND password = ?`

	// run query
	XOLog(sqlstr, name, password)
	p := Participant{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, name, password).Scan(&p.ID, &p.Name, &p.Password, &p.Firstname, &p.Lastname, &p.Qrhash)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

// ParticipantByID retrieves a row from 'testtt.participant' as a Participant.
//
// Generated from index 'participant_id_pkey'.
func ParticipantByID(db XODB, id int) (*Participant, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, password, firstname, lastname, qrhash ` +
		`FROM testtt.participant ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	p := Participant{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&p.ID, &p.Name, &p.Password, &p.Firstname, &p.Lastname, &p.Qrhash)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
