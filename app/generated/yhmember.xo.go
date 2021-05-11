// Package generated contains the types for schema 'devdb'.
package generated

import (
	"errors"
)

// YhMember represents a row from 'devdb.yh_member'.
type YhMember struct {
	TeamID     uint   `json:"team_id"`     // team_id
	MemberName string `json:"member_name"` // member_name
	Skill      string `json:"skill"`       // skill

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the YhMember exists in the database.
func (ym *YhMember) Exists() bool {
	return ym._exists
}

// Deleted provides information if the YhMember has been deleted from the database.
func (ym *YhMember) Deleted() bool {
	return ym._deleted
}

// Insert inserts the YhMember to the database.
func (ym *YhMember) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if ym._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO devdb.yh_member (` +
		`team_id, member_name, skill` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, ym.TeamID, ym.MemberName, ym.Skill)
	_, err = db.Exec(sqlstr, ym.TeamID, ym.MemberName, ym.Skill)
	if err != nil {
		return err
	}

	// set existence
	ym._exists = true

	return nil
}

// Update updates the YhMember in the database.
func (ym *YhMember) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ym._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if ym._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query with composite primary key
	const sqlstr = `UPDATE devdb.yh_member SET ` +
		`skill = ?` +
		` WHERE team_id = ? AND member_name = ?`

	// run query
	XOLog(sqlstr, ym.Skill, ym.TeamID, ym.MemberName)
	_, err = db.Exec(sqlstr, ym.Skill, ym.TeamID, ym.MemberName)
	return err
}

// Save saves the YhMember to the database.
func (ym *YhMember) Save(db XODB) error {
	if ym.Exists() {
		return ym.Update(db)
	}

	return ym.Insert(db)
}

// Delete deletes the YhMember from the database.
func (ym *YhMember) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ym._exists {
		return nil
	}

	// if deleted, bail
	if ym._deleted {
		return nil
	}

	// sql query with composite primary key
	const sqlstr = `DELETE FROM devdb.yh_member WHERE team_id = ? AND member_name = ?`

	// run query
	XOLog(sqlstr, ym.TeamID, ym.MemberName)
	_, err = db.Exec(sqlstr, ym.TeamID, ym.MemberName)
	if err != nil {
		return err
	}

	// set deleted
	ym._deleted = true

	return nil
}

// YhMemberByMemberName retrieves a row from 'devdb.yh_member' as a YhMember.
//
// Generated from index 'yh_member_member_name_pkey'.
func YhMemberByMemberName(db XODB, memberName string) (*YhMember, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`team_id, member_name, skill ` +
		`FROM devdb.yh_member ` +
		`WHERE member_name = ?`

	// run query
	XOLog(sqlstr, memberName)
	ym := YhMember{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, memberName).Scan(&ym.TeamID, &ym.MemberName, &ym.Skill)
	if err != nil {
		return nil, err
	}

	return &ym, nil
}
