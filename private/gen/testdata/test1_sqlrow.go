// Code generated by "gen_test.go test case 1"; DO NOT EDIT

package testdata

import (
	"github.com/jjeffery/errors"
)

// Get a Document by its primary key. Returns nil if not found.
func (q DocumentQuery) Get(id string) (*Document, error) {
	var row Document
	n, err := q.schema.Select(q.db, &row, "documents", id)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get Document").With(
			"id", id,
		)
	}
	if n == 0 {
		return nil, nil
	}
	return &row, nil
}

// Select a list of Documents from an SQL query.
func (q DocumentQuery) Select(query string, args ...interface{}) ([]*Document, error) {
	var rows []*Document
	_, err := q.schema.Select(q.db, &rows, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot query Documents").With(
			"query", query,
			"args", args,
		)
	}
	return rows, nil
}

// Select a Document from an SQL query. Returns nil if the query returns no rows.
// If the query returns one or more rows the value for the first is returned and any subsequent
// rows are discarded.
func (q DocumentQuery) SelectOne(query string, args ...interface{}) (*Document, error) {
	var row Document
	n, err := q.schema.Select(q.db, &row, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot query one Document").With(
			"query", query,
			"args", args,
		)
	}
	if n == 0 {
		return nil, nil
	}
	return &row, nil
}

// Insert a Document row.
func (q DocumentQuery) Insert(row *Document) error {
	err := q.schema.Insert(q.db, row, "documents")
	if err != nil {
		return errors.Wrap(err, "cannot insert Document").With(
			"ID", row.ID,
		)
	}
	return nil
}

// Update an existing Document row. Returns the number of rows updated,
// which should be zero or one.
func (q DocumentQuery) Update(row *Document) (int, error) {
	n, err := q.schema.Update(q.db, row, "documents")
	if err != nil {
		return 0, errors.Wrap(err, "cannot update Document").With(
			"ID", row.ID,
		)
	}
	return n, nil
}

// Attempt to update a Document row, and if it does not exist then insert it.
func (q DocumentQuery) Upsert(row *Document) error {
	n, err := q.schema.Update(q.db, row, "documents")
	if err != nil {
		return errors.Wrap(err, "cannot update Document for upsert").With(
			"ID", row.ID,
		)
	}
	if n > 0 {
		// update successful, row updated
		return nil
	}
	if err := q.schema.Insert(q.db, row, "documents"); err != nil {
		return errors.Wrap(err, "cannot insert Document for upsert").With(
			"ID", row.ID,
		)
	}
	return nil
}

// Delete a Document row. Returns the number of rows deleted, which should
// be zero or one.
func (q DocumentQuery) Delete(row *Document) (int, error) {
	n, err := q.schema.Delete(q.db, row, "documents")
	if err != nil {
		return 0, errors.Wrap(err, "cannot delete Document").With(
			"ID", row.ID,
		)
	}
	return n, nil
}
