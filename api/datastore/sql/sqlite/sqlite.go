package sqlite

import (
	"errors"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/dukhyungkim/fn/api/datastore/sql/dbhelper"
	"github.com/jmoiron/sqlx"
	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
)

type sqliteHelper int

func (sqliteHelper) Supports(scheme string) bool {
	switch scheme {
	case "sqlite3", "sqlite":
		return true
	}
	return false
}

func (sqliteHelper) PreConnect(url *url.URL) (string, error) {
	// make all the dirs so we can make the file..
	dir := filepath.Dir(url.Path)
	err := os.MkdirAll(dir, 0750)
	if err != nil {
		return "", err
	}

	return strings.TrimPrefix(url.String(), url.Scheme+"://"), nil
}

func (sqliteHelper) PostCreate(db *sqlx.DB) (*sqlx.DB, error) {
	db.SetMaxOpenConns(1)
	return db, nil
}

func (sqliteHelper) CheckTableExists(tx *sqlx.Tx, table string) (bool, error) {
	query := tx.Rebind(`SELECT count(*)
		FROM sqlite_master
		WHERE name = ?`)

	row := tx.QueryRow(query, table)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}

	exists := count > 0
	return exists, nil
}

func (sqliteHelper) String() string {
	return "sqlite"
}

func (sqliteHelper) IsDuplicateKeyError(err error) bool {
	var dbErr *sqlite.Error
	if errors.As(err, &dbErr) {
		code := dbErr.Code()
		if code == sqlite3.SQLITE_CONSTRAINT_UNIQUE || code == sqlite3.SQLITE_CONSTRAINT_PRIMARYKEY {
			return true
		}
	}
	return false
}

func init() {
	dbhelper.Register(sqliteHelper(0))
}
