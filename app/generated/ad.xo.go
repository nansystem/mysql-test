// Package generated contains the types for schema 'devdb'.
package generated

import (
	"errors"
)

// Ad represents a row from 'devdb.ad'.
type Ad struct {
	ID          int    `json:"id"`          // id
	Title       string `json:"title"`       // title
	Description string `json:"description"` // description
	Type        int8   `json:"type"`        // type
	StartAt     int    `json:"start_at"`    // start_at
	EndAt       int    `json:"end_at"`      // end_at
	UpdatedAt   int    `json:"updated_at"`  // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Ad exists in the database.
func (a *Ad) Exists() bool {
	return a._exists
}

// Deleted provides information if the Ad has been deleted from the database.
func (a *Ad) Deleted() bool {
	return a._deleted
}

// Insert inserts the Ad to the database.
func (a *Ad) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if a._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO devdb.ad (` +
		`title, description, type, start_at, end_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, a.Title, a.Description, a.Type, a.StartAt, a.EndAt, a.UpdatedAt)
	res, err := db.Exec(sqlstr, a.Title, a.Description, a.Type, a.StartAt, a.EndAt, a.UpdatedAt)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	a.ID = int(id)
	a._exists = true

	return nil
}

// Update updates the Ad in the database.
func (a *Ad) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !a._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if a._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE devdb.ad SET ` +
		`title = ?, description = ?, type = ?, start_at = ?, end_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, a.Title, a.Description, a.Type, a.StartAt, a.EndAt, a.UpdatedAt, a.ID)
	_, err = db.Exec(sqlstr, a.Title, a.Description, a.Type, a.StartAt, a.EndAt, a.UpdatedAt, a.ID)
	return err
}

// Save saves the Ad to the database.
func (a *Ad) Save(db XODB) error {
	if a.Exists() {
		return a.Update(db)
	}

	return a.Insert(db)
}

// Delete deletes the Ad from the database.
func (a *Ad) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !a._exists {
		return nil
	}

	// if deleted, bail
	if a._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM devdb.ad WHERE id = ?`

	// run query
	XOLog(sqlstr, a.ID)
	_, err = db.Exec(sqlstr, a.ID)
	if err != nil {
		return err
	}

	// set deleted
	a._deleted = true

	return nil
}

// AdByID retrieves a row from 'devdb.ad' as a Ad.
//
// Generated from index 'ad_id_pkey'.
func AdByID(db XODB, id int) (*Ad, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, title, description, type, start_at, end_at, updated_at ` +
		`FROM devdb.ad ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	a := Ad{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&a.ID, &a.Title, &a.Description, &a.Type, &a.StartAt, &a.EndAt, &a.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

// AdsByTypeEndAt retrieves a row from 'devdb.ad' as a Ad.
//
// Generated from index 'idx_ad_type_end_at'.
func AdsByTypeEndAt(db XODB, typ int8, endAt int) ([]*Ad, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, title, description, type, start_at, end_at, updated_at ` +
		`FROM devdb.ad ` +
		`WHERE type = ? AND end_at = ?`

	// run query
	XOLog(sqlstr, typ, endAt)
	q, err := db.Query(sqlstr, typ, endAt)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Ad{}
	for q.Next() {
		a := Ad{
			_exists: true,
		}

		// scan
		err = q.Scan(&a.ID, &a.Title, &a.Description, &a.Type, &a.StartAt, &a.EndAt, &a.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &a)
	}

	return res, nil
}

// AdsByTypeEndAtID retrieves a row from 'devdb.ad' as a Ad.
//
// Generated from index 'idx_ad_type_end_at_id'.
func AdsByTypeEndAtID(db XODB, typ int8, endAt int, id int) ([]*Ad, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, title, description, type, start_at, end_at, updated_at ` +
		`FROM devdb.ad ` +
		`WHERE type = ? AND end_at = ? AND id = ?`

	// run query
	XOLog(sqlstr, typ, endAt, id)
	q, err := db.Query(sqlstr, typ, endAt, id)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Ad{}
	for q.Next() {
		a := Ad{
			_exists: true,
		}

		// scan
		err = q.Scan(&a.ID, &a.Title, &a.Description, &a.Type, &a.StartAt, &a.EndAt, &a.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &a)
	}

	return res, nil
}

// AdsByIDTypeEndAt retrieves a row from 'devdb.ad' as a Ad.
//
// Generated from index 'idx_id_ad_type_end_at'.
func AdsByIDTypeEndAt(db XODB, id int, typ int8, endAt int) ([]*Ad, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, title, description, type, start_at, end_at, updated_at ` +
		`FROM devdb.ad ` +
		`WHERE id = ? AND type = ? AND end_at = ?`

	// run query
	XOLog(sqlstr, id, typ, endAt)
	q, err := db.Query(sqlstr, id, typ, endAt)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Ad{}
	for q.Next() {
		a := Ad{
			_exists: true,
		}

		// scan
		err = q.Scan(&a.ID, &a.Title, &a.Description, &a.Type, &a.StartAt, &a.EndAt, &a.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &a)
	}

	return res, nil
}
