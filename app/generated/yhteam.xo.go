// Package generated contains the types for schema 'devdb'.
package generated

import (
	"errors"
)

// YhTeam represents a row from 'devdb.yh_team'.
type YhTeam struct {
	TeamID   uint   `json:"team_id"`   // team_id
	TeamName string `json:"team_name"` // team_name

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the YhTeam exists in the database.
func (yt *YhTeam) Exists() bool {
	return yt._exists
}

// Deleted provides information if the YhTeam has been deleted from the database.
func (yt *YhTeam) Deleted() bool {
	return yt._deleted
}

// Insert inserts the YhTeam to the database.
func (yt *YhTeam) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if yt._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO devdb.yh_team (` +
		`team_id, team_name` +
		`) VALUES (` +
		`?, ?` +
		`)`

	// run query
	XOLog(sqlstr, yt.TeamID, yt.TeamName)
	_, err = db.Exec(sqlstr, yt.TeamID, yt.TeamName)
	if err != nil {
		return err
	}

	// set existence
	yt._exists = true

	return nil
}

// Update updates the YhTeam in the database.
func (yt *YhTeam) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !yt._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if yt._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE devdb.yh_team SET ` +
		`team_name = ?` +
		` WHERE team_id = ?`

	// run query
	XOLog(sqlstr, yt.TeamName, yt.TeamID)
	_, err = db.Exec(sqlstr, yt.TeamName, yt.TeamID)
	return err
}

// Save saves the YhTeam to the database.
func (yt *YhTeam) Save(db XODB) error {
	if yt.Exists() {
		return yt.Update(db)
	}

	return yt.Insert(db)
}

// Delete deletes the YhTeam from the database.
func (yt *YhTeam) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !yt._exists {
		return nil
	}

	// if deleted, bail
	if yt._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM devdb.yh_team WHERE team_id = ?`

	// run query
	XOLog(sqlstr, yt.TeamID)
	_, err = db.Exec(sqlstr, yt.TeamID)
	if err != nil {
		return err
	}

	// set deleted
	yt._deleted = true

	return nil
}

// YhTeamByTeamID retrieves a row from 'devdb.yh_team' as a YhTeam.
//
// Generated from index 'yh_team_team_id_pkey'.
func YhTeamByTeamID(db XODB, teamID uint) (*YhTeam, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`team_id, team_name ` +
		`FROM devdb.yh_team ` +
		`WHERE team_id = ?`

	// run query
	XOLog(sqlstr, teamID)
	yt := YhTeam{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, teamID).Scan(&yt.TeamID, &yt.TeamName)
	if err != nil {
		return nil, err
	}

	return &yt, nil
}
