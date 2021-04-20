// Package generated contains the types for schema 'devdb'.
package generated

// Code generated by xo. DO NOT EDIT.

import (
	// "database/sql"
	// "database/sql/driver"
	// "encoding/csv"
	"errors"
	// "fmt"
	// "regexp"
	// "strings"
	// "time"
)

// Cv represents a row from 'devdb.cv'.
type Cv struct {
	ID        int  `json:"id"`         // id
	AdID      int  `json:"ad_id"`      // ad_id
	UserID    int  `json:"user_id"`    // user_id
	Status    int8 `json:"status"`     // status
	CreatedAt int  `json:"created_at"` // created_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Cv exists in the database.
func (c *Cv) Exists() bool {
	return c._exists
}

// Deleted provides information if the Cv has been deleted from the database.
func (c *Cv) Deleted() bool {
	return c._deleted
}

// Insert inserts the Cv to the database.
func (c *Cv) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO devdb.cv (` +
		`id, ad_id, user_id, status, created_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, c.ID, c.AdID, c.UserID, c.Status, c.CreatedAt)
	_, err = db.Exec(sqlstr, c.ID, c.AdID, c.UserID, c.Status, c.CreatedAt)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Update updates the Cv in the database.
func (c *Cv) Update(db XODB) error {
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
	const sqlstr = `UPDATE devdb.cv SET ` +
		`ad_id = ?, user_id = ?, status = ?, created_at = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, c.AdID, c.UserID, c.Status, c.CreatedAt, c.ID)
	_, err = db.Exec(sqlstr, c.AdID, c.UserID, c.Status, c.CreatedAt, c.ID)
	return err
}

// Save saves the Cv to the database.
func (c *Cv) Save(db XODB) error {
	if c.Exists() {
		return c.Update(db)
	}

	return c.Insert(db)
}

// Delete deletes the Cv from the database.
func (c *Cv) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM devdb.cv WHERE id = ?`

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

// CvByID retrieves a row from 'devdb.cv' as a Cv.
//
// Generated from index 'cv_id_pkey'.
func CvByID(db XODB, id int) (*Cv, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, ad_id, user_id, status, created_at ` +
		`FROM devdb.cv ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	c := Cv{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&c.ID, &c.AdID, &c.UserID, &c.Status, &c.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &c, nil
}