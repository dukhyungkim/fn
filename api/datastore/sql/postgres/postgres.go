package postgres

import (
	"errors"
	"net/url"

	"github.com/dukhyungkim/fn/api/datastore/sql/dbhelper"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type postgresHelper int

func (postgresHelper) Supports(scheme string) bool {
	switch scheme {
	case "postgres", "pgx":
		return true
	}
	return false
}

func (postgresHelper) PreConnect(url *url.URL) (string, error) {
	return url.String(), nil
}

func (postgresHelper) PostCreate(db *sqlx.DB) (*sqlx.DB, error) {
	return db, nil
}

func (postgresHelper) CheckTableExists(tx *sqlx.Tx, table string) (bool, error) {
	query := tx.Rebind(`SELECT count(*)
	FROM information_schema.TABLES
	WHERE TABLE_NAME = ?`)

	row := tx.QueryRow(query, table)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}

	exists := count > 0
	return exists, nil
}

func (postgresHelper) String() string {
	return "postgres"
}

func (postgresHelper) IsDuplicateKeyError(err error) bool {
	var dbErr *pq.Error
	if errors.As(err, &dbErr) {
		if dbErr.Code == ErrUniqueViolation {
			return true
		}
	}
	return false
}

func init() {
	dbhelper.Register(postgresHelper(0))
}
