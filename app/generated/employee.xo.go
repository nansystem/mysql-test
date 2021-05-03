// Package generated contains the types for schema 'devdb'.
package generated

import (
	"errors"
	"time"
)

// Employee represents a row from 'devdb.employees'.
type Employee struct {
	EmployeeID  uint      `json:"employee_id"`   // employee_id
	FirstName   string    `json:"first_name"`    // first_name
	LastName    string    `json:"last_name"`     // last_name
	DateOfBirth time.Time `json:"date_of_birth"` // date_of_birth
	PhoneNumber string    `json:"phone_number"`  // phone_number

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Employee exists in the database.
func (e *Employee) Exists() bool {
	return e._exists
}

// Deleted provides information if the Employee has been deleted from the database.
func (e *Employee) Deleted() bool {
	return e._deleted
}

// Insert inserts the Employee to the database.
func (e *Employee) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if e._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO devdb.employees (` +
		`employee_id, first_name, last_name, date_of_birth, phone_number` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, e.EmployeeID, e.FirstName, e.LastName, e.DateOfBirth, e.PhoneNumber)
	_, err = db.Exec(sqlstr, e.EmployeeID, e.FirstName, e.LastName, e.DateOfBirth, e.PhoneNumber)
	if err != nil {
		return err
	}

	// set existence
	e._exists = true

	return nil
}

// Update updates the Employee in the database.
func (e *Employee) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !e._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if e._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE devdb.employees SET ` +
		`first_name = ?, last_name = ?, date_of_birth = ?, phone_number = ?` +
		` WHERE employee_id = ?`

	// run query
	XOLog(sqlstr, e.FirstName, e.LastName, e.DateOfBirth, e.PhoneNumber, e.EmployeeID)
	_, err = db.Exec(sqlstr, e.FirstName, e.LastName, e.DateOfBirth, e.PhoneNumber, e.EmployeeID)
	return err
}

// Save saves the Employee to the database.
func (e *Employee) Save(db XODB) error {
	if e.Exists() {
		return e.Update(db)
	}

	return e.Insert(db)
}

// Delete deletes the Employee from the database.
func (e *Employee) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !e._exists {
		return nil
	}

	// if deleted, bail
	if e._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM devdb.employees WHERE employee_id = ?`

	// run query
	XOLog(sqlstr, e.EmployeeID)
	_, err = db.Exec(sqlstr, e.EmployeeID)
	if err != nil {
		return err
	}

	// set deleted
	e._deleted = true

	return nil
}

// EmployeeByEmployeeID retrieves a row from 'devdb.employees' as a Employee.
//
// Generated from index 'employees_employee_id_pkey'.
func EmployeeByEmployeeID(db XODB, employeeID uint) (*Employee, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`employee_id, first_name, last_name, date_of_birth, phone_number ` +
		`FROM devdb.employees ` +
		`WHERE employee_id = ?`

	// run query
	XOLog(sqlstr, employeeID)
	e := Employee{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, employeeID).Scan(&e.EmployeeID, &e.FirstName, &e.LastName, &e.DateOfBirth, &e.PhoneNumber)
	if err != nil {
		return nil, err
	}

	return &e, nil
}
