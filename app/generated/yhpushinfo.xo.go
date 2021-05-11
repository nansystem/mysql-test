// Package generated contains the types for schema 'devdb'.
package generated

import (
	"errors"
	"time"
)

// YhPushInfo represents a row from 'devdb.yh_push_info'.
type YhPushInfo struct {
	PushID     uint      `json:"push_id"`     // push_id
	ModDate    time.Time `json:"mod_date"`    // mod_date
	DeletedFlg int8      `json:"deleted_flg"` // deleted_flg

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the YhPushInfo exists in the database.
func (ypi *YhPushInfo) Exists() bool {
	return ypi._exists
}

// Deleted provides information if the YhPushInfo has been deleted from the database.
func (ypi *YhPushInfo) Deleted() bool {
	return ypi._deleted
}

// Insert inserts the YhPushInfo to the database.
func (ypi *YhPushInfo) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if ypi._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO devdb.yh_push_info (` +
		`push_id, mod_date, deleted_flg` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, ypi.PushID, ypi.ModDate, ypi.DeletedFlg)
	_, err = db.Exec(sqlstr, ypi.PushID, ypi.ModDate, ypi.DeletedFlg)
	if err != nil {
		return err
	}

	// set existence
	ypi._exists = true

	return nil
}

// Update updates the YhPushInfo in the database.
func (ypi *YhPushInfo) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ypi._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if ypi._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE devdb.yh_push_info SET ` +
		`mod_date = ?, deleted_flg = ?` +
		` WHERE push_id = ?`

	// run query
	XOLog(sqlstr, ypi.ModDate, ypi.DeletedFlg, ypi.PushID)
	_, err = db.Exec(sqlstr, ypi.ModDate, ypi.DeletedFlg, ypi.PushID)
	return err
}

// Save saves the YhPushInfo to the database.
func (ypi *YhPushInfo) Save(db XODB) error {
	if ypi.Exists() {
		return ypi.Update(db)
	}

	return ypi.Insert(db)
}

// Delete deletes the YhPushInfo from the database.
func (ypi *YhPushInfo) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ypi._exists {
		return nil
	}

	// if deleted, bail
	if ypi._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM devdb.yh_push_info WHERE push_id = ?`

	// run query
	XOLog(sqlstr, ypi.PushID)
	_, err = db.Exec(sqlstr, ypi.PushID)
	if err != nil {
		return err
	}

	// set deleted
	ypi._deleted = true

	return nil
}

// YhPushInfoByPushID retrieves a row from 'devdb.yh_push_info' as a YhPushInfo.
//
// Generated from index 'yh_push_info_push_id_pkey'.
func YhPushInfoByPushID(db XODB, pushID uint) (*YhPushInfo, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`push_id, mod_date, deleted_flg ` +
		`FROM devdb.yh_push_info ` +
		`WHERE push_id = ?`

	// run query
	XOLog(sqlstr, pushID)
	ypi := YhPushInfo{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, pushID).Scan(&ypi.PushID, &ypi.ModDate, &ypi.DeletedFlg)
	if err != nil {
		return nil, err
	}

	return &ypi, nil
}
