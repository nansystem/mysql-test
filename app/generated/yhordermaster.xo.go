// Package generated contains the types for schema 'devdb'.
package generated

import (
	"errors"
	"time"
)

// YhOrderMaster represents a row from 'devdb.yh_order_master'.
type YhOrderMaster struct {
	ID           uint      `json:"id"`             // id
	OrderTime    time.Time `json:"order_time"`     // order_time
	SellerID     uint      `json:"seller_id"`      // seller_id
	ImageID      uint      `json:"image_id"`       // image_id
	ItemID       uint      `json:"item_id"`        // item_id
	IsHiddenPage int8      `json:"is_hidden_page"` // is_hidden_page

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the YhOrderMaster exists in the database.
func (yom *YhOrderMaster) Exists() bool {
	return yom._exists
}

// Deleted provides information if the YhOrderMaster has been deleted from the database.
func (yom *YhOrderMaster) Deleted() bool {
	return yom._deleted
}

// Insert inserts the YhOrderMaster to the database.
func (yom *YhOrderMaster) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if yom._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO devdb.yh_order_master (` +
		`id, order_time, seller_id, image_id, item_id, is_hidden_page` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, yom.ID, yom.OrderTime, yom.SellerID, yom.ImageID, yom.ItemID, yom.IsHiddenPage)
	_, err = db.Exec(sqlstr, yom.ID, yom.OrderTime, yom.SellerID, yom.ImageID, yom.ItemID, yom.IsHiddenPage)
	if err != nil {
		return err
	}

	// set existence
	yom._exists = true

	return nil
}

// Update updates the YhOrderMaster in the database.
func (yom *YhOrderMaster) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !yom._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if yom._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE devdb.yh_order_master SET ` +
		`order_time = ?, seller_id = ?, image_id = ?, item_id = ?, is_hidden_page = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, yom.OrderTime, yom.SellerID, yom.ImageID, yom.ItemID, yom.IsHiddenPage, yom.ID)
	_, err = db.Exec(sqlstr, yom.OrderTime, yom.SellerID, yom.ImageID, yom.ItemID, yom.IsHiddenPage, yom.ID)
	return err
}

// Save saves the YhOrderMaster to the database.
func (yom *YhOrderMaster) Save(db XODB) error {
	if yom.Exists() {
		return yom.Update(db)
	}

	return yom.Insert(db)
}

// Delete deletes the YhOrderMaster from the database.
func (yom *YhOrderMaster) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !yom._exists {
		return nil
	}

	// if deleted, bail
	if yom._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM devdb.yh_order_master WHERE id = ?`

	// run query
	XOLog(sqlstr, yom.ID)
	_, err = db.Exec(sqlstr, yom.ID)
	if err != nil {
		return err
	}

	// set deleted
	yom._deleted = true

	return nil
}

// YhOrderMasterByID retrieves a row from 'devdb.yh_order_master' as a YhOrderMaster.
//
// Generated from index 'yh_order_master_id_pkey'.
func YhOrderMasterByID(db XODB, id uint) (*YhOrderMaster, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, order_time, seller_id, image_id, item_id, is_hidden_page ` +
		`FROM devdb.yh_order_master ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	yom := YhOrderMaster{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&yom.ID, &yom.OrderTime, &yom.SellerID, &yom.ImageID, &yom.ItemID, &yom.IsHiddenPage)
	if err != nil {
		return nil, err
	}

	return &yom, nil
}
