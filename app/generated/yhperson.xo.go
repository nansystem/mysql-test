// Package generated contains the types for schema 'devdb'.
package generated

import (
	"errors"
)

// YhPerson represents a row from 'devdb.yh_people'.
type YhPerson struct {
	PrefID uint   `json:"pref_id"` // pref_id
	Name   string `json:"name"`    // name
	Age    uint   `json:"age"`     // age

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the YhPerson exists in the database.
func (yp *YhPerson) Exists() bool {
	return yp._exists
}

// Deleted provides information if the YhPerson has been deleted from the database.
func (yp *YhPerson) Deleted() bool {
	return yp._deleted
}

// Insert inserts the YhPerson to the database.
func (yp *YhPerson) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if yp._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO devdb.yh_people (` +
		`pref_id, name, age` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, yp.PrefID, yp.Name, yp.Age)
	_, err = db.Exec(sqlstr, yp.PrefID, yp.Name, yp.Age)
	if err != nil {
		return err
	}

	// set existence
	yp._exists = true

	return nil
}

// Update updates the YhPerson in the database.
func (yp *YhPerson) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !yp._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if yp._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query with composite primary key
	const sqlstr = `UPDATE devdb.yh_people SET ` +
		`age = ?` +
		` WHERE pref_id = ? AND name = ?`

	// run query
	XOLog(sqlstr, yp.Age, yp.PrefID, yp.Name)
	_, err = db.Exec(sqlstr, yp.Age, yp.PrefID, yp.Name)
	return err
}

// Save saves the YhPerson to the database.
func (yp *YhPerson) Save(db XODB) error {
	if yp.Exists() {
		return yp.Update(db)
	}

	return yp.Insert(db)
}

// Delete deletes the YhPerson from the database.
func (yp *YhPerson) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !yp._exists {
		return nil
	}

	// if deleted, bail
	if yp._deleted {
		return nil
	}

	// sql query with composite primary key
	const sqlstr = `DELETE FROM devdb.yh_people WHERE pref_id = ? AND name = ?`

	// run query
	XOLog(sqlstr, yp.PrefID, yp.Name)
	_, err = db.Exec(sqlstr, yp.PrefID, yp.Name)
	if err != nil {
		return err
	}

	// set deleted
	yp._deleted = true

	return nil
}

// YhPersonByName retrieves a row from 'devdb.yh_people' as a YhPerson.
//
// Generated from index 'yh_people_name_pkey'.
func YhPersonByName(db XODB, name string) (*YhPerson, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`pref_id, name, age ` +
		`FROM devdb.yh_people ` +
		`WHERE name = ?`

	// run query
	XOLog(sqlstr, name)
	yp := YhPerson{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, name).Scan(&yp.PrefID, &yp.Name, &yp.Age)
	if err != nil {
		return nil, err
	}

	return &yp, nil
}
