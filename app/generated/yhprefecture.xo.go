// Package generated contains the types for schema 'devdb'.
package generated

import (
	"errors"
)

// YhPrefecture represents a row from 'devdb.yh_prefecture'.
type YhPrefecture struct {
	PrefID         uint   `json:"pref_id"`         // pref_id
	Prefecture     string `json:"prefecture"`      // prefecture
	PrefectureKana string `json:"prefecture_kana"` // prefecture_kana

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the YhPrefecture exists in the database.
func (yp *YhPrefecture) Exists() bool {
	return yp._exists
}

// Deleted provides information if the YhPrefecture has been deleted from the database.
func (yp *YhPrefecture) Deleted() bool {
	return yp._deleted
}

// Insert inserts the YhPrefecture to the database.
func (yp *YhPrefecture) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if yp._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO devdb.yh_prefecture (` +
		`prefecture, prefecture_kana` +
		`) VALUES (` +
		`?, ?` +
		`)`

	// run query
	XOLog(sqlstr, yp.Prefecture, yp.PrefectureKana)
	res, err := db.Exec(sqlstr, yp.Prefecture, yp.PrefectureKana)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	yp.PrefID = uint(id)
	yp._exists = true

	return nil
}

// Update updates the YhPrefecture in the database.
func (yp *YhPrefecture) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !yp._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if yp._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE devdb.yh_prefecture SET ` +
		`prefecture = ?, prefecture_kana = ?` +
		` WHERE pref_id = ?`

	// run query
	XOLog(sqlstr, yp.Prefecture, yp.PrefectureKana, yp.PrefID)
	_, err = db.Exec(sqlstr, yp.Prefecture, yp.PrefectureKana, yp.PrefID)
	return err
}

// Save saves the YhPrefecture to the database.
func (yp *YhPrefecture) Save(db XODB) error {
	if yp.Exists() {
		return yp.Update(db)
	}

	return yp.Insert(db)
}

// Delete deletes the YhPrefecture from the database.
func (yp *YhPrefecture) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !yp._exists {
		return nil
	}

	// if deleted, bail
	if yp._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM devdb.yh_prefecture WHERE pref_id = ?`

	// run query
	XOLog(sqlstr, yp.PrefID)
	_, err = db.Exec(sqlstr, yp.PrefID)
	if err != nil {
		return err
	}

	// set deleted
	yp._deleted = true

	return nil
}

// YhPrefectureByPrefID retrieves a row from 'devdb.yh_prefecture' as a YhPrefecture.
//
// Generated from index 'yh_prefecture_pref_id_pkey'.
func YhPrefectureByPrefID(db XODB, prefID uint) (*YhPrefecture, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`pref_id, prefecture, prefecture_kana ` +
		`FROM devdb.yh_prefecture ` +
		`WHERE pref_id = ?`

	// run query
	XOLog(sqlstr, prefID)
	yp := YhPrefecture{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, prefID).Scan(&yp.PrefID, &yp.Prefecture, &yp.PrefectureKana)
	if err != nil {
		return nil, err
	}

	return &yp, nil
}
