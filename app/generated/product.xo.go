// Package generated contains the types for schema 'devdb'.
package generated

import (
	"errors"
	"time"
)

// Product represents a row from 'devdb.products'.
type Product struct {
	ID       uint      `json:"id"`        // id
	ShopID   uint      `json:"shop_id"`   // shop_id
	Name     string    `json:"name"`      // name
	Price    uint      `json:"price"`     // price
	StartsAt time.Time `json:"starts_at"` // starts_at
	EndsAt   time.Time `json:"ends_at"`   // ends_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Product exists in the database.
func (p *Product) Exists() bool {
	return p._exists
}

// Deleted provides information if the Product has been deleted from the database.
func (p *Product) Deleted() bool {
	return p._deleted
}

// Insert inserts the Product to the database.
func (p *Product) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if p._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO devdb.products (` +
		`id, shop_id, name, price, starts_at, ends_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, p.ID, p.ShopID, p.Name, p.Price, p.StartsAt, p.EndsAt)
	_, err = db.Exec(sqlstr, p.ID, p.ShopID, p.Name, p.Price, p.StartsAt, p.EndsAt)
	if err != nil {
		return err
	}

	// set existence
	p._exists = true

	return nil
}

// Update updates the Product in the database.
func (p *Product) Update(db XODB) error {
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
	const sqlstr = `UPDATE devdb.products SET ` +
		`shop_id = ?, name = ?, price = ?, starts_at = ?, ends_at = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, p.ShopID, p.Name, p.Price, p.StartsAt, p.EndsAt, p.ID)
	_, err = db.Exec(sqlstr, p.ShopID, p.Name, p.Price, p.StartsAt, p.EndsAt, p.ID)
	return err
}

// Save saves the Product to the database.
func (p *Product) Save(db XODB) error {
	if p.Exists() {
		return p.Update(db)
	}

	return p.Insert(db)
}

// Delete deletes the Product from the database.
func (p *Product) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM devdb.products WHERE id = ?`

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

// ProductByID retrieves a row from 'devdb.products' as a Product.
//
// Generated from index 'products_id_pkey'.
func ProductByID(db XODB, id uint) (*Product, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, shop_id, name, price, starts_at, ends_at ` +
		`FROM devdb.products ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	p := Product{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&p.ID, &p.ShopID, &p.Name, &p.Price, &p.StartsAt, &p.EndsAt)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
