package postgres

import "github.com/lib/pq"

const (
	ErrUniqueViolation pq.ErrorCode = "23505"
)
