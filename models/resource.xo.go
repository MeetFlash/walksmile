// Package model contains the types for schema 'walksmile'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
	"time"
)

// Resource represents a row from 'walksmile.resources'.
type Resource struct {
	ResourcesID  int       `json:"resources_id"`  // resources_id
	Title        string    `json:"title"`         // title
	Otitle       string    `json:"otitle"`        // otitle
	Image        string    `json:"image"`         // image
	Originsite   string    `json:"originsite"`    // originsite
	Intro        string    `json:"intro"`         // intro
	Kind         int       `json:"kind"`          // kind
	SharedfileID uint      `json:"sharedfile_id"` // sharedfile_id
	Price        int       `json:"price"`         // price
	CreateTime   time.Time `json:"create_time"`   // create_time
	UpdateTime   time.Time `json:"update_time"`   // update_time
	LinkEntity   string    `json:"link_entity"`   // link_entity

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Resource exists in the database.
func (r *Resource) Exists() bool {
	return r._exists
}

// Deleted provides information if the Resource has been deleted from the database.
func (r *Resource) Deleted() bool {
	return r._deleted
}

// Insert inserts the Resource to the database.
func (r *Resource) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if r._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO walksmile.resources (` +
		`title, otitle, image, originsite, intro, kind, sharedfile_id, price, link_entity` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?,  ?` +
		`)`

	// run query
	XOLog(sqlstr, r.Title, r.Otitle, r.Image, r.Originsite, r.Intro, r.Kind, r.SharedfileID, r.Price, r.LinkEntity)
	res, err := db.Exec(sqlstr, r.Title, r.Otitle, r.Image, r.Originsite, r.Intro, r.Kind, r.SharedfileID, r.Price, r.LinkEntity)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	r.ResourcesID = int(id)
	r._exists = true

	return nil
}

// Update updates the Resource in the database.
func (r *Resource) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !r._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if r._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE walksmile.resources SET ` +
		`title = ?, otitle = ?, image = ?, originsite = ?, intro = ?, kind = ?, sharedfile_id = ?, price = ?, create_time = ?, update_time = ?, link_entity = ?` +
		` WHERE resources_id = ?`

	// run query
	XOLog(sqlstr, r.Title, r.Otitle, r.Image, r.Originsite, r.Intro, r.Kind, r.SharedfileID, r.Price, r.CreateTime, r.UpdateTime, r.LinkEntity, r.ResourcesID)
	_, err = db.Exec(sqlstr, r.Title, r.Otitle, r.Image, r.Originsite, r.Intro, r.Kind, r.SharedfileID, r.Price, r.CreateTime, r.UpdateTime, r.LinkEntity, r.ResourcesID)
	return err
}

// Save saves the Resource to the database.
func (r *Resource) Save(db XODB) error {
	if r.Exists() {
		return r.Update(db)
	}

	return r.Insert(db)
}

// Delete deletes the Resource from the database.
func (r *Resource) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !r._exists {
		return nil
	}

	// if deleted, bail
	if r._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM walksmile.resources WHERE resources_id = ?`

	// run query
	XOLog(sqlstr, r.ResourcesID)
	_, err = db.Exec(sqlstr, r.ResourcesID)
	if err != nil {
		return err
	}

	// set deleted
	r._deleted = true

	return nil
}

// ResourceByOriginsite retrieves a row from 'walksmile.resources' as a Resource.
//
// Generated from index 'originsite_UNIQUE'.
func ResourceByOriginsite(db XODB, originsite string) (*Resource, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`resources_id, title, otitle, image, originsite, intro, kind, sharedfile_id, price, create_time, update_time, link_entity ` +
		`FROM walksmile.resources ` +
		`WHERE originsite = ?`

	// run query
	XOLog(sqlstr, originsite)
	r := Resource{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, originsite).Scan(&r.ResourcesID, &r.Title, &r.Otitle, &r.Image, &r.Originsite, &r.Intro, &r.Kind, &r.SharedfileID, &r.Price, &r.CreateTime, &r.UpdateTime, &r.LinkEntity)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// ResourceByResourcesID retrieves a row from 'walksmile.resources' as a Resource.
//
// Generated from index 'resources_resources_id_pkey'.
func ResourceByResourcesID(db XODB, resourcesID int) (*Resource, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`resources_id, title, otitle, image, originsite, intro, kind, sharedfile_id, price, create_time, update_time, link_entity ` +
		`FROM walksmile.resources ` +
		`WHERE resources_id = ?`

	// run query
	XOLog(sqlstr, resourcesID)
	r := Resource{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, resourcesID).Scan(&r.ResourcesID, &r.Title, &r.Otitle, &r.Image, &r.Originsite, &r.Intro, &r.Kind, &r.SharedfileID, &r.Price, &r.CreateTime, &r.UpdateTime, &r.LinkEntity)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// ResourcesBySharedfileID retrieves a row from 'walksmile.resources' as a Resource.
//
// Generated from index 'sharedfile_INDEX'.
func ResourcesBySharedfileID(db XODB, sharedfileID uint) ([]*Resource, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`resources_id, title, otitle, image, originsite, intro, kind, sharedfile_id, price, create_time, update_time, link_entity ` +
		`FROM walksmile.resources ` +
		`WHERE sharedfile_id = ?`

	// run query
	XOLog(sqlstr, sharedfileID)
	q, err := db.Query(sqlstr, sharedfileID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Resource{}
	for q.Next() {
		r := Resource{
			_exists: true,
		}

		// scan
		err = q.Scan(&r.ResourcesID, &r.Title, &r.Otitle, &r.Image, &r.Originsite, &r.Intro, &r.Kind, &r.SharedfileID, &r.Price, &r.CreateTime, &r.UpdateTime, &r.LinkEntity)
		if err != nil {
			return nil, err
		}

		res = append(res, &r)
	}

	return res, nil
}
