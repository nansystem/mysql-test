// Package generated contains the types for schema 'devdb'.
package generated

import (
	"errors"
	"time"
)

// PushInfo represents a row from 'devdb.push_info'.
type PushInfo struct {
	PushID     uint      `json:"push_id"`     // push_id
	ModDate    time.Time `json:"mod_date"`    // mod_date
	DeletedFlg int8      `json:"deleted_flg"` // deleted_flg

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the PushInfo exists in the database.
func (pi *PushInfo) Exists() bool {
	return pi._exists
}

// Deleted provides information if the PushInfo has been deleted from the database.
func (pi *PushInfo) Deleted() bool {
	return pi._deleted
}

// Insert inserts the PushInfo to the database.
func (pi *PushInfo) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if pi._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO devdb.push_info (` +
		`push_id, mod_date, deleted_flg` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, pi.PushID, pi.ModDate, pi.DeletedFlg)
	_, err = db.Exec(sqlstr, pi.PushID, pi.ModDate, pi.DeletedFlg)
	if err != nil {
		return err
	}

	// set existence
	pi._exists = true

	return nil
}

// Update updates the PushInfo in the database.
func (pi *PushInfo) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !pi._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if pi._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE devdb.push_info SET ` +
		`mod_date = ?, deleted_flg = ?` +
		` WHERE push_id = ?`

	// run query
	XOLog(sqlstr, pi.ModDate, pi.DeletedFlg, pi.PushID)
	_, err = db.Exec(sqlstr, pi.ModDate, pi.DeletedFlg, pi.PushID)
	return err
}

// Save saves the PushInfo to the database.
func (pi *PushInfo) Save(db XODB) error {
	if pi.Exists() {
		return pi.Update(db)
	}

	return pi.Insert(db)
}

// Delete deletes the PushInfo from the database.
func (pi *PushInfo) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !pi._exists {
		return nil
	}

	// if deleted, bail
	if pi._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM devdb.push_info WHERE push_id = ?`

	// run query
	XOLog(sqlstr, pi.PushID)
	_, err = db.Exec(sqlstr, pi.PushID)
	if err != nil {
		return err
	}

	// set deleted
	pi._deleted = true

	return nil
}

// PushInfoByPushID retrieves a row from 'devdb.push_info' as a PushInfo.
//
// Generated from index 'push_info_push_id_pkey'.
func PushInfoByPushID(db XODB, pushID uint) (*PushInfo, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`push_id, mod_date, deleted_flg ` +
		`FROM devdb.push_info ` +
		`WHERE push_id = ?`

	// run query
	XOLog(sqlstr, pushID)
	pi := PushInfo{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, pushID).Scan(&pi.PushID, &pi.ModDate, &pi.DeletedFlg)
	if err != nil {
		return nil, err
	}

	return &pi, nil
}
