package sqlite

import (
	"errors"

	"github.com/mattn/go-sqlite3"
)

// isDuplicate reports whether err is a primary key or unique constraint
// violation, i.e. the row already exists.
func isDuplicate(err error) bool {
	var sqliteErr sqlite3.Error

	if !errors.As(err, &sqliteErr) {
		return false
	}

	return sqliteErr.ExtendedCode == sqlite3.ErrConstraintPrimaryKey ||
		sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique
}
