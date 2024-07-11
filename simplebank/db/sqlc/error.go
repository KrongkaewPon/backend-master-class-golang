package db

import (
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var ErrRecordNotFound = pgx.ErrNoRows

const (
	ForeignKeyViolation = "23503"
	UniqueViolation     = "23505"
)

var ErrUniqueViolation = &pgconn.PgError{
	Code: UniqueViolation,
}

func ErrorCode(err error) string {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		// fmt.Println(pgErr.ConstraintName)
		return pgErr.Code
	}
	return ""
}
